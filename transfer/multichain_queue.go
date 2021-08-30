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

package transfer

import (
	"errors"
	"fmt"
)

// MultichainQueue is map of Queues which are able to
// queue transaction execution funcs in parallel for each chain.
type MultichainQueue struct {
	qs map[int64]Enqueuer
}

type Enqueuer interface {
	TransactionEnqueue(exec TransactionSendFn, resp chan<- QueueResponse) error
}

// ErrMissingQueueForChainId is returned when a given chainID doesnt have an existing queue .
var ErrMissingQueueForChainId = errors.New("no queue for chain")

// NewMultichainQueue returns a new empty multichain queue.
func NewMultichainQueue() *MultichainQueue {
	return &MultichainQueue{
		qs: make(map[int64]Enqueuer),
	}
}

// AddQueue will add a queue with the given chain ID.
func (mq *MultichainQueue) AddQueue(chainID int64, queue Enqueuer) {
	mq.qs[chainID] = queue
}

// TransactionEnqueue will enqueue a given transaction in the queue identified by the given chainId.
//
// Enqueue will fail and instantly return an error if the queue doesnt exist.
func (mq *MultichainQueue) TransactionEnqueue(chainID int64, exec TransactionSendFn, resp chan<- QueueResponse) error {
	c, ok := mq.qs[chainID]
	if !ok {
		return fmt.Errorf("chainID: %v %w", chainID, ErrMissingQueueForChainId)
	}

	return c.TransactionEnqueue(exec, resp)
}
