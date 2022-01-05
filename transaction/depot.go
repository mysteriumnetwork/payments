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
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Depot is a transaction delivery depot. Using a given `DeliveryCourier`
// it will try to deliver transactions to the blockchain.
//
// Upon successfull dispatch it will track transactions making
// sure that they reach their destination. If they're stuck
// it will reissue a transaction with more gas.
type Depot struct {
	handler DeliveryCourier
	storage DepotStorage

	gasStation   *GasTracker
	nonceTracker *DepotNonceTracker

	config DepotConfig
	logFn  func(error)

	once sync.Once
	stop chan struct{}
}

type DepotConfig struct {
	Workers         []DepotWorker
	MaxNonDelivered uint
}

// DepotWorker is a worker that will spawn upon starting `Run`.
// Each worker is reponsible for its own transactions only.
type DepotWorker struct {
	Address         common.Address
	ChainID         int64
	ProcessInterval time.Duration
	ProcessCount    uint
}

// DeliveryCourier is responsible for executing transaction on the blockchain.
// It should be able to send and resend the same transaction multiple times.
// If it cannot handle a transaction, it will remain in the depot.
type DeliveryCourier interface {
	// DeliverTransaction should determine the transaction type and try to delivery it.
	DeliverTransaction(tx Delivery) (*types.Transaction, error)
	// CanDeliver should validate the given type and see if the courier will be able to delivery it.
	// If returned false, transaction will not be queued.
	CanDeliver(DeliverableType) bool
}

// DepotStorage is a persistent storage used to keep track of work to do.
type DepotStorage interface {
	// GetOrderedDeliveryRequests should return a list of delivery requests ordered by nonce, from lowest to highest.
	GetOrderedDeliveryRequests(count uint, chainID int64, sender common.Address) ([]Delivery, error)
	// GetLastQueuedDelivery is used to track nonce
	GetLastQueuedDelivery(chainID int64, sender common.Address) (*Delivery, error)
	// GetLastDelivered is used to check when the last transaction was delivered.
	// It will be used to determine the next time we need to increase gas price.
	GetLastDelivered(chainID int64, sender common.Address) (*Delivery, error)

	// GetNonDeliveredCount is used to get a total number of non delivered transactions.
	// If that number is higher or equal to the configured max number, no new
	// transaction will be set in the queue.
	GetNonDeliveredCount(chainID int64, sender common.Address) (uint, error)

	// UpsertDeliveryRequest is used to update a delivery. If no delivery with a given
	// unique ID exists, it should create it.
	UpsertDeliveryRequest(tx Delivery) error
}

var ErrImpossibleToDeliver = errors.New("impossible to deliver")

// NewDepot will returns a new depot.
func NewDepot(handler DeliveryCourier, storage DepotStorage, nonce *DepotNonceTracker, gasStation *GasTracker, cfg DepotConfig) *Depot {
	return &Depot{
		handler:      handler,
		storage:      storage,
		nonceTracker: nonce,
		gasStation:   gasStation,

		config: cfg,
		stop:   make(chan struct{}),
	}
}

// Run will spawn a goroutine for each loaded `DepotWorker`.
func (d *Depot) Run() {
	for _, s := range d.config.Workers {
		go d.watchDeliveries(s)
	}
}

// EnqueueDelivery will submit a new transaction to the delivery queue.
// It will return a unique tracking number which can be used to see the status of a transaction.
func (d *Depot) EnqueueDelivery(req DeliveryRequest) (string, error) {
	if !d.workerExists(req) {
		return "", fmt.Errorf("failed to enqueue for sender %q: no worker found", req.Sender.Hex())
	}

	if !d.handler.CanDeliver(req.Type) {
		return "", fmt.Errorf("transaction will not set in queue, not possible to delivery type %q", req.Type)
	}

	count, err := d.storage.GetNonDeliveredCount(req.ChainID, req.Sender)
	if err != nil {
		return "", fmt.Errorf("could not get non delivered count: %w", err)
	}

	if d.config.MaxNonDelivered <= count {
		return "", fmt.Errorf("cannot queue a new entry, max count of %d reached", d.config.MaxNonDelivered)
	}

	unqID := ""
	setFn := func(nonce uint64) error {
		td, err := req.toDelivery(nonce)
		if err != nil {
			return fmt.Errorf("failed to marshal queue entry: %w", err)
		}
		unqID = td.UniqueID

		if err := d.storage.UpsertDeliveryRequest(td); err != nil {
			return fmt.Errorf("failed to insert a delivery request: %w", err)
		}

		return nil
	}

	err = d.nonceTracker.SetNextNonce(req.ChainID, req.Sender, setFn)
	if err != nil {
		return "", fmt.Errorf("failed to issue a nonce for a transaction delivery: %w", err)
	}

	return unqID, nil
}

// Stop will stop the Deposit goroutines.
func (d *Depot) Stop() {
	d.once.Do(func() {
		close(d.stop)
	})
}

// AttachLogger allows the caller to attach an optional logger.
// Logger logs non critical errors that happen during transaction
// handling and will be eventually handled by the Depot.
//
// This method is not thread safe and should be called before `Run`.
func (d *Depot) AttachLogger(fn func(err error)) {
	d.logFn = fn
}

func (d *Depot) workerExists(req DeliveryRequest) bool {
	for _, s := range d.config.Workers {
		if s.Address.Hex() == req.Sender.Hex() && req.ChainID == s.ChainID {
			return true
		}
	}

	return false
}

func (d *Depot) watchDeliveries(s DepotWorker) {
	for {
		select {
		case <-d.stop:
			return
		case <-time.After(s.ProcessInterval):
			tds, err := d.storage.GetOrderedDeliveryRequests(s.ProcessCount, s.ChainID, s.Address)
			if err != nil {
				d.logFn(err)
				break
			}

			for _, td := range tds {
				select {
				case <-d.stop:
					return
				default:
					d.handleDeliveryRequest(td)
				}
			}
		}
	}
}

func (d *Depot) handleDeliveryRequest(td Delivery) {
	switch td.State {
	case DeliveryStateWaiting:
		if err := d.handleWaiting(td); err != nil {
			d.logFn(err)
		}
	case DeliveryStatePacking, DeliveryStateSent:
		if err := d.handleTrackingRequired(td); err != nil {
			d.logFn(err)
		}
	}
}

func (d *Depot) handleTrackingRequired(td Delivery) error {
	currentNonce, err := d.nonceTracker.GetConfirmedNonce(td.ChainID, td.Sender)
	if err != nil {
		return err
	}

	// If our nonce increased that means someone is aware about this transaction
	// and we can continue.
	if currentNonce > td.Nonce {
		if td.State == DeliveryStateWaiting {
			return fmt.Errorf("refusing to confirm transaction as it was never sent from our side: %q", td.UniqueID)
		}

		if _, err := d.markDeliveryAsDelivered(td); err != nil {
			return fmt.Errorf("failed to mark delivery as sent: %w", err)
		}

		return nil
	}

	// Only try to resubmit the earliest transaction sent.
	// Other transactions might have enough gas and this
	// might be the only blocking transaction.
	if currentNonce != td.Nonce {
		return nil
	}

	// Don't spam. Wait 1 minute in between increment checks.
	if td.UpdateUTC.Add(time.Minute).After(time.Now().UTC()) {
		return nil
	}

	ld, err := d.storage.GetLastDelivered(td.ChainID, td.Sender)
	if err != nil {
		return fmt.Errorf("failed to get last delivered: %w", err)
	}

	lastTime := td.UpdateUTC
	if ld != nil {
		lastTime = ld.UpdateUTC
	}

	ok, err := d.gasStation.CanReceiveMoreGas(td.ChainID, lastTime)
	if err != nil {
		return fmt.Errorf("tried to redeliver %q, but got err: %w", td.UniqueID, err)
	}

	// We can only resend a transaction only if we give more gas
	// because if not, it will simply get rejected by the node
	// as you cannot replace a transaction without giving more gas.
	if ok {
		return d.sendOutTransaction(td)
	}

	return nil
}

func (d *Depot) handleWaiting(td Delivery) error {
	if _, err := d.markDeliveryAsPacking(td); err != nil {
		return fmt.Errorf("failed to mark packing for sender %q reason: %w", td.Sender.Hex(), err)
	}

	return d.sendOutTransaction(td)
}

func (d *Depot) sendOutTransaction(td Delivery) error {
	var err error
	td, err = d.deliveryRecalculateAndSaveGasPrice(td)
	if err != nil {
		return err
	}

	tx, err := d.handler.DeliverTransaction(td)
	if err != nil {
		return fmt.Errorf("attempted to delivery a transaction %q for account %q but failed: %w", td.UniqueID, td.Sender.Hex(), err)
	}

	if _, err := d.markDeliveryAsSent(td, tx); err != nil {
		return fmt.Errorf("failed to mark delivery as sent: %w", err)
	}

	return nil
}

func (d *Depot) log(err error) {
	if d.logFn != nil {
		d.logFn(err)
	}
}

func (d *Depot) deliveryRecalculateAndSaveGasPrice(td Delivery) (Delivery, error) {
	gas, err := d.gasStation.RecalculateDeliveryGas(td.ChainID, td.GasPrice)
	if err != nil {
		return td, fmt.Errorf("delivery %q failed to calculate gas price: %w", td.UniqueID, err)
	}

	td.GasPrice = gas
	td.UpdateUTC = time.Now().UTC()
	if err := d.storage.UpsertDeliveryRequest(td); err != nil {
		return td, fmt.Errorf("failed to mark delivery as packing: %w", err)
	}

	return td, nil
}

func (d *Depot) markDeliveryAsPacking(td Delivery) (Delivery, error) {
	td.State = DeliveryStatePacking
	td.UpdateUTC = time.Now().UTC()
	if err := d.storage.UpsertDeliveryRequest(td); err != nil {
		return td, fmt.Errorf("failed to mark delivery as packing: %w", err)
	}

	return td, nil
}

func (d *Depot) markDeliveryAsDelivered(td Delivery) (Delivery, error) {
	td.State = DeliveryStateDelivered
	td.UpdateUTC = time.Now().UTC()
	if err := d.storage.UpsertDeliveryRequest(td); err != nil {
		return td, fmt.Errorf("failed to mark delivery as sent: %w", err)
	}

	return td, nil
}

func (d *Depot) markDeliveryAsSent(td Delivery, tx *types.Transaction) (Delivery, error) {
	marshaled, err := tx.MarshalJSON()
	if err != nil {
		return td, err
	}

	td.State = DeliveryStateSent
	td.SentTransaction = marshaled
	td.UpdateUTC = time.Now().UTC()
	if err := d.storage.UpsertDeliveryRequest(td); err != nil {
		return td, fmt.Errorf("failed to mark delivery as sent: %w", err)
	}

	return td, nil
}
