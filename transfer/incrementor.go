/* Mysterium network payment library.
 *
 * Copyright (C) 2021 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package transfer

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// GasPriceIncremenetor exposes a way automatically increment gas fees
// for transactions in a preconfigured manner.
type GasPriceIncremenetor struct {
	storage Storage
	bc      MultichainClient

	cfg     GasIncrementorConfig
	signers safeSigners

	logFn LogFunc
	stop  chan struct{}
	once  sync.Once
}

// GasIncrementorConfig is provided to the incrementor to configure it.
type GasIncrementorConfig struct {
	WorkerCount       int
	PullInterval      time.Duration
	PullEntryCount    int64
	MaxQueuePerSigner int
}

// Storage is given to the Incremeter to be used to
// insert, update or get transactions.
type Storage interface {
	// UpsertIncrementorTransaction is called to upsert a transaction.
	// It either inserts a new entry or updates existing entries.
	UpsertIncrementorTransaction(tx Transaction) error

	// GetIncrementorTransactionsToCheck returns all transaction that need to rechecked.
	//
	// Entries should be filtered by possible signers. If incrementor cannot sign the transaction
	// it should not received it.
	GetIncrementorTransactionsToCheck(maxEntries int64, possibleSigners []string) (tx []Transaction, err error)

	// GasIncrementorSenderQueue returns the length of a queue for a single sender.
	GetIncrementorSenderQueue(sender string) (length int, err error)
}

// MultichainClient handles calls to BC.
type MultichainClient interface {
	TransactionReceipt(chainID int64, hash common.Hash) (*types.Receipt, error)
	SendTransaction(chainID int64, tx *types.Transaction) error
	TransactionByHash(chainID int64, hash common.Hash) (*types.Transaction, bool, error)
}

// LogFunc can be attacheched to Incrementer to enable logging.
type LogFunc func(Transaction, error)

// NewGasPriceIncremenetor returns a new incrementer instance.
func NewGasPriceIncremenetor(cfg GasIncrementorConfig, storage Storage, cl MultichainClient, signers Signers) *GasPriceIncremenetor {
	if cfg.WorkerCount <= 0 {
		cfg.WorkerCount = 1
	}
	if cfg.PullEntryCount <= 0 {
		cfg.PullEntryCount = 1
	}

	return &GasPriceIncremenetor{
		storage: storage,
		bc:      cl,

		cfg: cfg,
		signers: safeSigners{
			signers: signers,
		},

		stop: make(chan struct{}, 0),
	}
}

// Run starts the gas price incrementer.
//
// It will query the given storage for any entries that it needs to check
// for gas increase, trying to check their status.
func (i *GasPriceIncremenetor) Run() {
	for {
		select {
		case <-i.stop:
			return

		case <-time.After(i.cfg.PullInterval):
			txs, err := i.storage.GetIncrementorTransactionsToCheck(i.cfg.PullEntryCount, i.signers.getSigners())
			if err != nil {
				continue
			}

			i.queueWork(txs)
		}
	}
}

func (i *GasPriceIncremenetor) watchOrIncrease(work chan Transaction) {
	for tx := range work {
		// Force skip transactions that are done in case the provider doesn't.
		switch tx.State {
		case TxStateCreated, TxStatePriceIncreased:
		default:
			continue
		}

		now := time.Now().UTC()

		if tx.isExpired() {
			i.transactionExpired(tx)
			continue
		}

		ok, err := i.checkIfSuccess(tx)
		if err != nil {
			i.log(tx, fmt.Errorf("status check failed: %w", err))
			if !errors.Is(err, ethereum.NotFound) {
				continue
			}
		}

		// If transaction was not a success and a certain amount of time
		// has passed since last increase - increase gas price again.
		if !ok && now.After(tx.NextIncreaseAfterUTC) {
			if err := i.increaseGasPrice(tx); err != nil {
				if i.isBlockchainErrorUnhandleable(err) {
					i.log(tx, fmt.Errorf("received unhandleable increase error, marking tx as failed: %w", err))
					i.transactionFailed(tx)
					continue
				}

				i.log(tx, fmt.Errorf("gas price increase failed: %w", err))
			}
		}
	}
}

func (i *GasPriceIncremenetor) queueWork(txs []Transaction) {
	work := make(chan Transaction)

	var wg sync.WaitGroup
	for in := 0; in < i.cfg.WorkerCount; in++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i.watchOrIncrease(work)
		}()
	}

	// Close and wait on func exit
	defer func() {
		close(work)
		wg.Wait()
	}()

	for _, o := range txs {
		select {
		case <-i.stop:
			return
		default:
			work <- o
		}
	}
}

// AttachLogFunc attaches a new log func incremeter thread.
// The given log func is called every time the running incrementer thread
// created by the `Run` method encouters an error.
//
// This method is not thread safe and should be called before Run.
func (i *GasPriceIncremenetor) AttachLogFunc(logFn LogFunc) {
	i.logFn = logFn
}

// Stop stops the execution of GasPriceIncrementer thread created by the Run method.
func (i *GasPriceIncremenetor) Stop() {
	i.once.Do(func() {
		close(i.stop)
	})
}

// InsertInitial uses the given storage to insert an new transaction which
// will later be retreived using `GetTransactionsToCheck` in order to check
// it's state and retry with higher gas price if needed.
func (i *GasPriceIncremenetor) InsertInitial(tx *types.Transaction, opts TransactionOpts, senderAddress common.Address) error {
	if err := opts.validate(); err != nil {
		return fmt.Errorf("invalid opts given: %w", err)
	}
	newTx, err := newTransaction(tx, senderAddress, opts)
	if err != nil {
		return fmt.Errorf("failed to create new transaction: %w", err)
	}

	return i.storage.UpsertIncrementorTransaction(*newTx)
}

// CanSign returns if incrementor is able to sign transactions for the given sender.
func (i *GasPriceIncremenetor) CanSign(sender common.Address) bool {
	_, ok := i.signers.getSignerFunc(sender.Hex())
	return ok
}

// CanQueue returns if incrementor is able to queue a transaction
func (i *GasPriceIncremenetor) CanQueue(sender common.Address) (bool, error) {
	length, err := i.storage.GetIncrementorSenderQueue(sender.Hex())
	if err != nil {
		return false, err
	}

	return length < i.cfg.MaxQueuePerSigner, nil
}

func (i *GasPriceIncremenetor) checkIfSuccess(tx Transaction) (bool, error) {
	status, err := i.getTxStatus(tx)
	if err != nil {
		return false, err
	}

	switch status {
	case StatusSucceeded:
		return true, i.transactionSuccess(tx)
	default:
		return false, i.transactionSetNextCheck(tx)
	}
}

func (i *GasPriceIncremenetor) increaseGasPrice(tx Transaction) error {
	org, err := tx.getLatestTx()
	if err != nil {
		return err
	}

	newGasPrice, err := calculateGasPrice(tx, org.GasPrice())
	if err != nil {
		return err
	}

	newTx, err := i.signAndSend(tx.rebuiledWithNewGasPrice(org, newGasPrice), tx.ChainID, tx.SenderAddressHex)
	if err != nil {
		return err
	}

	return i.transactionPriceIncreased(tx, newTx)
}

func (i *GasPriceIncremenetor) isBlockchainErrorUnhandleable(err error) bool {
	if errors.Is(err, core.ErrNonceTooHigh) || errors.Is(err, core.ErrNonceTooLow) {
		return true
	}

	unwrapped := errors.Unwrap(err)
	if unwrapped == nil {
		return false
	}

	// ethereum sometimes returns errors from some other, non core package, so resort to string checks
	switch unwrapped.Error() {
	case core.ErrNonceTooLow.Error(), core.ErrNonceTooHigh.Error():
		return true
	default:
		return false
	}
}

// BCTxStatus represents the status of tx on blockchain.
type BCTxStatus string

const (
	// StatusPending indicates that the tx is still pending.
	StatusPending BCTxStatus = "Pending"
	// StatusFailed indicates that the tx has failed.
	StatusFailed BCTxStatus = "Failed"
	// StatusSucceeded indicates that the tx has suceeded.
	StatusSucceeded BCTxStatus = "Succeeded"
)

func (i *GasPriceIncremenetor) getTxStatus(tx Transaction) (BCTxStatus, error) {
	org, err := tx.getLatestTx()
	if err != nil {
		return StatusFailed, fmt.Errorf("can't get tx status, malformed internal tx object: %w", err)
	}

	hash := org.Hash()
	_, pending, err := i.bc.TransactionByHash(tx.ChainID, hash)
	if err != nil {
		return StatusFailed, fmt.Errorf("failed to get transaction by hash: %w", err)
	}

	if pending {
		return StatusPending, nil
	}

	receipt, err := i.bc.TransactionReceipt(tx.ChainID, hash)
	if err != nil {
		return StatusFailed, fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	return i.bcTxStatusFromReceipt(tx, receipt), nil
}

func (i *GasPriceIncremenetor) bcTxStatusFromReceipt(tx Transaction, rcp *types.Receipt) BCTxStatus {
	switch rcp.Status {
	case types.ReceiptStatusSuccessful:
		return StatusSucceeded
	case types.ReceiptStatusFailed:
		return StatusFailed
	default:
		hash := ""
		if org, err := tx.getLatestTx(); err == nil {
			hash = org.Hash().Hex()
		}

		i.log(tx, fmt.Errorf("received unknown receipt status %v for tx uniqueID: '%v', lastHash: '%v'", rcp.Status, tx.UniqueID, hash))
		return StatusFailed
	}
}

func (i *GasPriceIncremenetor) signAndSend(tx *types.Transaction, chainID int64, senderAddrHex string) (*types.Transaction, error) {
	signer, ok := i.signers.getSignerFunc(senderAddrHex)
	if !ok {
		return nil, fmt.Errorf("can't retry, no signer for address: %s", senderAddrHex)
	}

	signedTx, err := signer(tx, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to sign a transaction: %w", err)
	}

	if err := i.bc.SendTransaction(chainID, signedTx); err != nil {
		return nil, fmt.Errorf("failed send a transaction: %w", err)
	}

	return signedTx, nil
}

func (i *GasPriceIncremenetor) transactionFailed(tx Transaction) {
	tx.State = TxStateFailed

	err := i.storage.UpsertIncrementorTransaction(tx)
	if err != nil {
		i.log(tx, fmt.Errorf("failed marking transaction as failed: %w", err))
	}
}

func (i *GasPriceIncremenetor) transactionExpired(tx Transaction) {
	tx.State = TxStateExpired

	err := i.storage.UpsertIncrementorTransaction(tx)
	if err != nil {
		i.log(tx, fmt.Errorf("failed marking transaction as expired: %w", err))
	}
}

func (i *GasPriceIncremenetor) transactionSetNextCheck(tx Transaction) error {
	tx.NextCheckAfterUTC = time.Now().UTC().Add(tx.Opts.CheckInterval)
	if err := i.storage.UpsertIncrementorTransaction(tx); err != nil {
		return fmt.Errorf("failed marking transaction succeed: %w", err)
	}
	return nil
}

func (i *GasPriceIncremenetor) transactionSuccess(tx Transaction) error {
	tx.State = TxStateSucceed
	if err := i.storage.UpsertIncrementorTransaction(tx); err != nil {
		return fmt.Errorf("failed marking transaction succeed: %w", err)
	}
	return nil
}

func (i *GasPriceIncremenetor) transactionPriceIncreased(tx Transaction, newTx *types.Transaction) error {
	blob, err := newTx.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal internal transaction object: %w", err)
	}

	tx.State = TxStatePriceIncreased
	tx.NextIncreaseAfterUTC = time.Now().UTC().Add(tx.Opts.IncreaseInterval)
	tx.LatestTx = blob
	if err := i.storage.UpsertIncrementorTransaction(tx); err != nil {
		return fmt.Errorf("failed to update transaction after price increase: %w", err)
	}

	return nil
}

func (i *GasPriceIncremenetor) log(tx Transaction, err error) {
	if i.logFn != nil {
		i.logFn(tx, err)
	}
}

func calculateGasPrice(tx Transaction, currentPrice *big.Int) (*big.Int, error) {
	if currentPrice.Cmp(tx.Opts.MaxPrice) > 0 {
		return nil, fmt.Errorf("transaction with uniqueID '%s' failed, gas price limit of %s reached on chain %d", tx.UniqueID, tx.Opts.MaxPrice.String(), tx.ChainID)
	}

	newGasPrice, _ := new(big.Float).Mul(
		big.NewFloat(tx.Opts.PriceMultiplier),
		new(big.Float).SetInt(currentPrice),
	).Int(nil)

	if newGasPrice.Cmp(tx.Opts.MaxPrice) > 0 {
		if currentPrice.Cmp(tx.Opts.MaxPrice) < 0 {
			return tx.Opts.MaxPrice, nil
		}

		return nil, fmt.Errorf("transaction with uniqueID '%s' failed, gas price limit of %s reached on chain %d", tx.UniqueID, tx.Opts.MaxPrice.String(), tx.ChainID)
	}

	return newGasPrice, nil
}

// SignatureFunc is used to sign transactions when resubmitting them.
type SignatureFunc func(tx *types.Transaction, chainID int64) (*types.Transaction, error)

// Signers is a map that holds all possible signers to sign transactions when resending to the blockchain.
type Signers map[common.Address]SignatureFunc

type safeSigners struct {
	signers map[common.Address]SignatureFunc
	m       sync.Mutex
}

func (s *safeSigners) getSignerFunc(senderAddressHex string) (SignatureFunc, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	// TODO: Remove later. Left for backwards compatability.
	// If only one signer is present and senderAddress is `""` return first signer.
	if len(s.signers) == 1 && (senderAddressHex == "" || senderAddressHex == common.HexToAddress("").Hex()) {
		for _, v := range s.signers {
			return v, true
		}
	}

	addr := common.HexToAddress(senderAddressHex)
	signer, ok := s.signers[addr]
	return signer, ok
}

func (s *safeSigners) getSigners() []string {
	s.m.Lock()
	defer s.m.Unlock()

	signers := make([]string, 0)
	for signer := range s.signers {
		signers = append(signers, signer.Hex())
	}
	return signers
}
