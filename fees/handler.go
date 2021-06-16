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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionHandler wraps gas price increasers
// exposing a send method which handles gas prices when sending a transaction.
//
// TransactionHandler expects the gas price increaser to be
// handleded and started by the caller himself.
type TransactionHandler struct {
	inc   GasPriceIncremenetorIface
	logFn func(error)
}

// HandlerOpts are given when sending a new transaction.
type HandlerOpts struct {
	// SenderAddress is the address of the sender who is making the trasaction.
	SenderAddress common.Address
	// GasPriceIncOpts will be used to increase the gas price using an incrementor.
	GasPriceIncOpts TransactionOpts
	// Optional: ForceQueue allows to bypass the check if transaction is able to queue.
	ForceQueue bool
}

// TransactionSendFn wraps a transaction execution which should result
// in a transaction being returned if it was successfull.
//
// It allows to wrap different kinds of contract methods which
// all result in producing a transaction.
type TransactionSendFn func() (*types.Transaction, error)

// ErrQueueFull is returned if the current queue is full and incremeting gas price will be impossible.
// Transactions which receive this error should be retried later.
var ErrQueueFull = errors.New("failed to send a transaction, queue is full")

// ErrNoSigners is returned if there are no signers to use for incrementing this transaction.
var ErrNoSigners = errors.New("failed to send a transaction, no signers for incrementing")

// GasPriceIncremenetorIface abstracts gas price incrementor.
type GasPriceIncremenetorIface interface {
	// InsertInitial inserts a new transaction to the queue.
	InsertInitial(tx *types.Transaction, opts TransactionOpts, senderAddress common.Address) error

	// CanQueue returns true if another transaction can be queue in to the incrementor.
	CanQueue(sender common.Address) (bool, error)

	// CanSign returns if incrementor can sign this transaction when incrementing gas price.
	CanSign(sender common.Address) bool
}

// NewTransactionhandler returns a new transaction handler
func NewTransactionhandler(inc GasPriceIncremenetorIface) *TransactionHandler {
	return &TransactionHandler{
		inc: inc,
	}
}

// AttachLogger allows the caller to attach an optional logger.
// Logger logs non critical errors or warnings that happen during transaction
// handling.
//
// This method is not thread safe and should be called before `SendWithGasPriceHandling`.
func (t *TransactionHandler) AttachLogger(fn func(err error)) {
	t.logFn = fn
}

// SendWithGasPriceHandling given a new watchable transaction with options will send the transaction
// and increase the gas price accordingly.
func (t *TransactionHandler) SendWithGasPriceHandling(txSend TransactionSendFn, opts HandlerOpts) (*types.Transaction, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	if err := t.canQueue(opts); err != nil {
		return nil, err
	}

	tx, err := txSend()
	if err != nil {
		return nil, err
	}

	if err := t.inc.InsertInitial(tx, opts.GasPriceIncOpts, opts.SenderAddress); err != nil {
		t.log(fmt.Errorf("failed to insert initial entry for gas price incremenetor: %w", err))
	}

	return tx, nil
}

func (t *TransactionHandler) canQueue(opts HandlerOpts) error {
	if opts.ForceQueue {
		return nil
	}

	canQ, err := t.inc.CanQueue(opts.SenderAddress)
	if err != nil {
		return err
	}
	if !canQ {
		return ErrQueueFull
	}

	if !t.inc.CanSign(opts.SenderAddress) {
		return ErrNoSigners
	}

	return nil
}

func (opts *HandlerOpts) validate() error {
	if opts.SenderAddress == common.HexToAddress("") {
		return errors.New("sender address must be specified")
	}
	if err := opts.GasPriceIncOpts.validate(); err != nil {
		return err
	}

	return nil
}

func (t *TransactionHandler) log(err error) {
	if t.logFn != nil {
		t.logFn(err)
	}
}
