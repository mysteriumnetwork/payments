/*
 * Copyright (C) 2021 The "MysteriumNetwork/payments" Authors.
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

package sort

import (
	"sync"
	"time"

	"github.com/mysteriumnetwork/payments/client"
)

// MultiClientSorter is able to sort multiclient according to provided funcs.
// It can sort on multi triggered actions:
// * Notifications from multiclient warning of errors
// * Ticker which is ticked every X time.Duration
type MultiClientSorter struct {
	sort SortableClient

	// notifications channel is used to listen for events and execute actions.
	notifications <-chan client.Notification
	// checkEvery is optional duration that can be given to create a ticker
	// and execute on ticker actions.
	checkEvery *time.Duration

	onn []OnNotificationAction
	ont []OnTickerAction

	stop chan struct{}
	once sync.Once
}

// SortableClient clients must implement the interface to be interacted with.
type SortableClient interface {
	ReorderClients(addresses []string) error
	CurrentClientOrder() []string
	CallSpecificClient(address string, call func(c client.EtherClient) error) error
}

// ClientCall gives a EtherClient to call.
type ClientCall func(c client.EtherClient) error

// OnNotificationAction is called when a notification is received
type OnNotificationAction func(notification client.Notification, clients SortableClient)

// OnTickerAction is called when a ticker tick is received.
type OnTickerAction func(clients SortableClient)

// NewMultiClientSorter bootstraps the default sorter.
func NewMultiClientSorter(cl SortableClient, notifications <-chan client.Notification, tickEvery time.Duration) *MultiClientSorter {
	sorter := newSorter(cl, notifications)
	sorter.checkEvery = &tickEvery
	return sorter
}

// NewMultiClientSorterNoTicker bootraps the sorter without ticker, only listening for notifications.
func NewMultiClientSorterNoTicker(cl SortableClient, notifications <-chan client.Notification) *MultiClientSorter {
	return newSorter(cl, notifications)
}

func newSorter(cl SortableClient, notifications <-chan client.Notification) *MultiClientSorter {
	// Allow for nil notitification channel incase that was never created.
	if notifications == nil {
		notifications = make(chan client.Notification)
	}

	return &MultiClientSorter{
		sort:          cl,
		notifications: notifications,

		onn: make([]OnNotificationAction, 0),
		ont: make([]OnTickerAction, 0),

		stop: make(chan struct{}),
	}
}

// Run starts the sorter. It can start two different sorters:
// with ticker or without one depending on if ticker timer was set.
func (m *MultiClientSorter) Run() {
	if m.checkEvery != nil {
		m.runWithTicker()
	} else {
		m.runNoTicker()
	}
}

func (m *MultiClientSorter) runNoTicker() {
	for {
		select {
		case <-m.stop:
			return
		case notify := <-m.notifications:
			for _, fn := range m.onn {
				fn(notify, m.sort)
			}
		}
	}
}

func (m *MultiClientSorter) runWithTicker() {
	ticker := time.NewTicker(*m.checkEvery)
	defer ticker.Stop()

	for {
		select {
		case <-m.stop:
			return
		case notify := <-m.notifications:
			for _, fn := range m.onn {
				fn(notify, m.sort)
			}
		case <-ticker.C:
			for _, fn := range m.ont {
				fn(m.sort)
			}
		}
	}
}

// AddOnNotificationAction adds a new action on notification trigger.
// This method is not thread safe and should not be called after `Run`.
func (m *MultiClientSorter) AddOnNotificationAction(action OnNotificationAction) {
	m.onn = append(m.onn, action)
}

// AddOnTickerAction adds a new action on ticker trigger.
// This method is not thread safe and should not be called after `Run`.
func (m *MultiClientSorter) AddOnTickerAction(action OnTickerAction) {
	m.ont = append(m.ont, action)
}

// Stop stops the multisorter.
func (m *MultiClientSorter) Stop() {
	m.once.Do(func() {
		close(m.stop)
	})
}
