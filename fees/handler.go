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
// TransactionHandler expects watchers and increasers to be
// handleded and started by the caller hitself
// as it only sends and inserts transactions.
type TransactionHandler struct {
	inc   GasPriceIncremenetorIface
	logFn func(error)
}

// HandlerOpts are given when sending
// a new transaction.
type HandlerOpts struct {
	SenderAddress   common.Address
	GasPriceIncOpts TransactionOpts
}

// TransactionSendFn wraps a transaction execution which should result
// in a transaction being returned if it was successfull.
//
// It allows to wrap different kinds of contract methods which
// all result in producing a transaction.
type TransactionSendFn func() (*types.Transaction, error)

// GasPriceIncremenetorIface abstracts gas pricei incrementor.
type GasPriceIncremenetorIface interface {
	InsertInitial(tx *types.Transaction, opts TransactionOpts, senderAddress common.Address) error
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

	tx, err := txSend()
	if err != nil {
		return nil, err
	}

	if err := t.inc.InsertInitial(tx, opts.GasPriceIncOpts, opts.SenderAddress); err != nil {
		t.log(fmt.Errorf("failed to insert initial entry for gas price incremenetor: %w", err))
	}

	return tx, nil
}

func (opts *HandlerOpts) validate() error {
	if opts.SenderAddress == common.HexToAddress("") {
		return errors.New("sender address must be specified")
	}

	return nil
}

func (t *TransactionHandler) log(err error) {
	if t.logFn != nil {
		t.logFn(err)
	}
}
