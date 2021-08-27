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
	"context"
	"errors"
	"sort"
	"time"

	"github.com/mysteriumnetwork/payments/client"
	"github.com/rs/zerolog/log"
)

// DefaultByBlockSize will sort on ticker by block size.
func DefaultByBlockSize(cl SortableClient) {
	clients := cl.CurrentClientOrder()

	type kv struct {
		Address   string
		BlockSize uint64
	}

	blockNumbers := make([]kv, 0)
	for _, address := range clients {
		var blockNumber uint64
		err := cl.CallSpecificClient(address, func(c client.EtherClient) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			bn, err := c.BlockNumber(ctx)
			if err != nil {
				return err
			}

			blockNumber = bn
			return nil
		})
		if err != nil {
			continue
		}
		blockNumbers = append(blockNumbers, kv{
			Address:   address,
			BlockSize: blockNumber,
		})
	}

	if len(blockNumbers) != len(clients) {
		log.Error().Msg("failed to reorder clients, some client block numbers are missing")
		return
	}

	sort.Slice(blockNumbers, func(i, j int) bool {
		return blockNumbers[i].BlockSize > blockNumbers[j].BlockSize
	})

	newOrder := make([]string, len(clients))
	for i, v := range blockNumbers {
		newOrder[i] = v.Address
	}

	if err := cl.ReorderClients(newOrder); err != nil {
		log.Err(err).Msg("failed to re-order the RPC client slice")
	}
}

// DefaultByAvailability will sort by connection problems received from notifications.
// Other errors are ignored.
func DefaultByAvailability(not client.Notification, clients SortableClient) {
	if !errors.Is(not.Error, client.ErrClientNoConnection) && !errors.Is(not.Error, client.ErrClientTooManyRequests) {
		return
	}

	currentOrder := clients.CurrentClientOrder()
	if len(currentOrder) == 1 {
		return
	}

	address := not.Address
	for i, current := range currentOrder {
		if current != address {
			continue
		}

		if len(currentOrder)-1 == i {
			return
		}

		newOrder := append(currentOrder[:i], currentOrder[i+1:]...)
		if err := clients.ReorderClients(append(newOrder, address)); err != nil {
			log.Err(err).Msg("failed to re-order the RPC client slice")
		}

		return
	}
}
