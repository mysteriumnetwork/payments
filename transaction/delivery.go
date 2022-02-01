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

package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/mysteriumnetwork/payments/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Delivery struct {
	UniqueID string
	Sender   common.Address

	Nonce    uint64
	ChainID  int64
	GasPrice *big.Int

	Type  DeliverableType
	State DeliveryState

	ShipmentData    []byte
	SentTransaction []byte

	CreatedUTC time.Time
	UpdateUTC  time.Time
}

// DeliverableType is issued for the Courier to determine how to package a transaction.
type DeliverableType string

type DeliveryState string

const (
	// Waiting is state for transaction which were never sent
	DeliveryStateWaiting = "waiting"
	// DeliveryStatePacking is state for transaction that were or are
	// being attempted to be send but no confirmation about them leaving is received.
	DeliveryStatePacking = "packing"
	// DeliveryStateSent is a state for transaction that are for sure sent
	// and we can track them, raising gas price if needed.
	DeliveryStateSent = "sent"
	// DeliveryStateDelivered is state for transaction that have been delivered
	// and we will no longer track it.
	DeliveryStateDelivered = "delivered"
)

var ErrNoTrasactionExists = errors.New("transaction doesn't exist")

type SignFunc func(common.Address, *types.Transaction) (*types.Transaction, error)

// ToWriteRequest will convert a Delivery to a typical write request used by `client` package.
func (t *Delivery) ToWriteRequest(signer SignFunc, gasLimit uint64) client.WriteRequest {
	return client.WriteRequest{
		Identity: t.Sender,
		Signer:   bind.SignerFn(signer),
		GasLimit: gasLimit,
		GasTip:   t.GasPrice,
		Nonce:    new(big.Int).SetUint64(t.Nonce),
	}
}

func (t *Delivery) GetLastTransaction() (*types.Transaction, error) {
	if len(t.SentTransaction) == 0 {
		return nil, ErrNoTrasactionExists
	}

	tx := &types.Transaction{}
	return tx, tx.UnmarshalJSON(t.SentTransaction)
}

type DeliveryRequest struct {
	ChainID int64
	Sender  common.Address

	Type DeliverableType

	// Data must always be a marshable struct or nil
	Data interface{}
}

func (t *DeliveryRequest) toDelivery(nonce uint64) (Delivery, error) {
	now := time.Now().UTC()

	if t.Data == nil {
		t.Data = struct{}{}
	}

	blob, err := json.Marshal(t.Data)
	if err != nil {
		return Delivery{}, err
	}

	return Delivery{
		UniqueID: fmt.Sprintf("%s|%d|%d", t.Sender.Hex(), nonce, t.ChainID),
		Sender:   t.Sender,

		Nonce:    nonce,
		ChainID:  t.ChainID,
		GasPrice: new(big.Int).SetInt64(0),

		Type:  t.Type,
		State: DeliveryStateWaiting,

		ShipmentData:    blob,
		SentTransaction: []byte{},

		CreatedUTC: now,
		UpdateUTC:  now,
	}, nil
}
