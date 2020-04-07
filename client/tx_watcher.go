/*
 * Copyright (C) 2020 The "MysteriumNetwork/payments" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
func (tw *TxWatcher) EnsureTransactionSuccess(wt WatchableTransaction, initialGasPrice *big.Int) (*types.Transaction, func(), error) {
	stop := make(chan struct{})
	once := sync.Once{}
	cancel := func() {
		once.Do(func() {
			close(stop)
		})
	}
	tx, err := tw.watchTx(wt, initialGasPrice, stop)
	return tx, cancel, err
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
