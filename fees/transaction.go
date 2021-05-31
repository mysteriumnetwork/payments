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

package fees

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionState marks the current state of a given transaction
type TransactionState string

const (
	// TxStateCreated is given to transactions that have been initially inserted.
	TxStateCreated TransactionState = "created"
	// TxStatePriceIncreased is given to transactions which have received
	// a price increase.
	TxStatePriceIncreased TransactionState = "priceIncreased"
	// TxStateFailed is given to transactions which have
	// failed and should not be retried anymore.
	TxStateFailed TransactionState = "failed"
	// TxStateSucceed is given to transactions which have
	// succeeded and should not be retried.
	TxStateSucceed TransactionState = "succeed"
)

// Transaction objects is used when handling transactions.
type Transaction struct {
	// UniqueID is a combination of the original tx hash and chainID.
	UniqueID string
	Opts     TransactionOpts
	State    TransactionState

	OrignalHashHex   string
	SenderAddressHex string
	ChainID          int64

	LatestTx []byte
}

// TransactionOpts are provided when creating a new transaction.
// They're used when handling a transaction.
type TransactionOpts struct {
	PriceMultiplier  float64
	MaxPrice         *big.Int
	Timeout          time.Duration
	IncreaseInterval time.Duration
	CheckInterval    time.Duration

	// ValidUntil is an optional paremeter which
	// can be given to invalidate a transaction and mark it as failed
	// after the given time.Time.
	ValidUntil *time.Time
}

// TransactionUniqueID returns a unique ID for a transaction.
func TransactionUniqueID(orignalHash string, chainID int64) string {
	return fmt.Sprintf("%s|%d", orignalHash, chainID)
}

func (t *TransactionOpts) validate() error {
	if t.PriceMultiplier <= 1 {
		return errors.New("priceMultiplier must be more than 1")
	}
	if t.MaxPrice == nil || t.MaxPrice.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("max price has to be greater than 0")
	}
	if t.Timeout <= 0 {
		return errors.New("timeout value must be provided")
	}
	if t.IncreaseInterval <= 0 {
		return errors.New("increase interval value must be provided")
	}
	if t.CheckInterval <= 0 {
		return errors.New("check interval value must be provided")
	}
	if t.ValidUntil != nil && t.ValidUntil.Before(time.Now()) {
		return errors.New("given 'ValidUntil' must be in the future")
	}

	return nil
}

func newTransaction(tx *types.Transaction, senderAddress common.Address, opts TransactionOpts) (*Transaction, error) {
	hash := tx.Hash().Hex()

	marshaled, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return &Transaction{
		UniqueID:         TransactionUniqueID(hash, tx.ChainId().Int64()),
		Opts:             opts,
		State:            TxStateCreated,
		OrignalHashHex:   hash,
		SenderAddressHex: senderAddress.Hex(),
		ChainID:          tx.ChainId().Int64(),
		LatestTx:         marshaled,
	}, nil
}

func (t *Transaction) isExpired() bool {
	if t.Opts.ValidUntil == nil {
		return false
	}

	return time.Now().After(*t.Opts.ValidUntil)
}

func (t *Transaction) getLatestTx() (*types.Transaction, error) {
	tx := &types.Transaction{}
	return tx, tx.UnmarshalJSON(t.LatestTx)
}

func (t *Transaction) rebuiledWithNewGasPrice(tx *types.Transaction, newGasPrice *big.Int) *types.Transaction {
	return types.NewTransaction(
		tx.Nonce(),
		*tx.To(),
		tx.Value(),
		tx.Gas(),
		newGasPrice,
		tx.Data(),
	)
}
