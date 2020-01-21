package client

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

type client interface {
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error)
}

// TxWatcher makes sure that transactions actually get sent to the network.
// It retries for the given amount of times to send the TX, each time increasing the gas price by given percentage.
type TxWatcher struct {
	maxRetries         int
	waitBetweenRetries time.Duration
	percentageIncrease float64

	client        client
	clientTimeout time.Duration
}

// NewTxWatcher returns a new instance of tx watcher
func NewTxWatcher(client client, clientTimeout time.Duration, maxRetries int, waitBetweenRetries time.Duration, percentageIncrease float64) *TxWatcher {
	return &TxWatcher{
		client:             client,
		clientTimeout:      clientTimeout,
		maxRetries:         maxRetries,
		waitBetweenRetries: waitBetweenRetries,
		percentageIncrease: percentageIncrease,
	}
}

// WatchableTransaction represents a transaction that the txwatcher can keep track of
type WatchableTransaction func(gasPrice *big.Int) (*types.Transaction, error)

// EnsureTransactionSuccess keeps track of the transaction on blockchain.
// If it sees that the transaction has not appeared on BC, it will try again and increase the gas price.
func (tw *TxWatcher) EnsureTransactionSuccess(wt WatchableTransaction, initialGasPrice *big.Int) (chan *types.Transaction, func(), chan error) {
	stop := make(chan struct{})
	once := sync.Once{}
	cancel := func() {
		once.Do(func() {
			close(stop)
		})
	}
	errChan := make(chan error)
	txChan := make(chan *types.Transaction)

	go func() {
		defer close(txChan)
		defer close(errChan)
		tx, err := tw.watchTx(wt, initialGasPrice, stop)
		txChan <- tx
		errChan <- err
	}()

	return txChan, cancel, errChan
}

func (tw *TxWatcher) watchTx(wt WatchableTransaction, initialGasPrice *big.Int, stop <-chan struct{}) (*types.Transaction, error) {
	for i := 1; i <= tw.maxRetries; i++ {
		// if sending fails just notify the caller
		gasPrice := tw.calculateGasPrice(initialGasPrice, i)
		tx, err := wt(gasPrice)
		if err != nil {
			return nil, err
		}

		ctx, c := context.WithTimeout(context.Background(), tw.clientTimeout)
		defer c()
		pulledTx, _, err := tw.client.TransactionByHash(ctx, tx.Hash())
		if err != nil {
			log.Warn().Err(err).Msgf("transaction %q not visible on BC, will retry", tx.Hash())
			continue
		}

		if pulledTx != nil {
			return pulledTx, nil
		}

		select {
		case <-stop:
			return nil, errors.New("stop invoked")
		case <-time.After(tw.waitBetweenRetries):
		}
	}
	return nil, errors.New("ran out of retries")
}

func (tw *TxWatcher) calculateGasPrice(initialGasPrice *big.Int, attempt int) *big.Int {
	if attempt <= 1 {
		return initialGasPrice
	}

	timesToMultiply := attempt - 1

	multiplier := big.NewFloat(tw.percentageIncrease * float64(timesToMultiply))
	gasPriceFloat := big.NewFloat(0).SetInt(initialGasPrice)

	intermediary := big.NewFloat(0).Mul(multiplier, gasPriceFloat)
	intermediaryInt, _ := intermediary.Int(nil)

	return big.NewInt(0).Add(initialGasPrice, intermediaryInt)
}
