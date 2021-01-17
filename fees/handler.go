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

	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionHandler wraps gas price increasers
// exposing a send method which handles gas prices when sending a transaction.
//
// TransactionHandler expects watchers and increasers to be
// handleded and started by the caller hitself
// as it only sends and inserts transactions.
type TransactionHandler struct {
	txWatchers map[int64]TxWatcherIface
	inc        GasPriceIncremenetorIface

	logFn func(error)
}

// HandlerOpts are given when sending
// a new transaction.
type HandlerOpts struct {
	ChainID         int64
	InitGasPrice    *big.Int
	GasPriceIncOpts TransactionOpts
}

// TxWatcherIface abstract any transaction watcher.
type TxWatcherIface interface {
	EnsureTransactionSuccess(wt WatchableTransaction, initialGasPrice *big.Int) (*types.Transaction, func(), error)
}

// GasPriceIncremenetorIface abstracts gas pricei incrementor.
type GasPriceIncremenetorIface interface {
	InsertInitial(tx *types.Transaction, opts TransactionOpts, chainID int64) error
}

// ErrOperationStopped is returned if called action is cut short (not finished)
// and the result of the called action is unknown.
var ErrOperationStopped = errors.New("operation cut short, result unknown")

// NewTransactionhandler returns a new transaction handler
func NewTransactionhandler(watchers map[int64]TxWatcherIface, inc GasPriceIncremenetorIface) *TransactionHandler {
	return &TransactionHandler{
		txWatchers: watchers,
		inc:        inc,
	}
}

// AttachLogger allows the caller to attach an optional logger.
// Logger logs non critical errors or warnings that happen during transaction
// handling.
//
// This method is not thread safe and should be called before `SendWithgasPriceHandling`.
func (t *TransactionHandler) AttachLogger(fn func(err error)) {
	t.logFn = fn
}

// SendWithGasPriceHandling given a new watchable transaction with options will send the transaction
// and increase the gas price accordingly.
//
// This method blocks either until transaction is sent and inserted, or stop channel is closed.
func (t *TransactionHandler) SendWithGasPriceHandling(stop chan struct{}, wt WatchableTransaction, opts HandlerOpts) (*types.Transaction, error) {
	cancel := func() {}
	defer cancel()
	type res struct {
		tx  *types.Transaction
		err error
	}

	watcher, ok := t.txWatchers[opts.ChainID]
	if !ok {
		return nil, fmt.Errorf("no tx watcher for chain ID %v", opts.ChainID)
	}

	resChan := make(chan res)
	go func() {
		tx, c, err := watcher.EnsureTransactionSuccess(wt, opts.InitGasPrice)
		cancel = c
		resChan <- res{
			tx:  tx,
			err: err,
		}
	}()

	select {
	case <-stop:
		return nil, ErrOperationStopped
	case r := <-resChan:
		if r.err != nil {
			return r.tx, r.err
		}

		if err := t.inc.InsertInitial(r.tx, opts.GasPriceIncOpts, opts.ChainID); err != nil {
			t.log(fmt.Errorf("failed to insert initial entry for gas price incremenetor: %w", err))
		}

		return r.tx, nil
	}
}

func (t *TransactionHandler) log(err error) {
	if t.logFn != nil {
		t.logFn(err)
	}
}
