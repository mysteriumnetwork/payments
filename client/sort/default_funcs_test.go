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
	"testing"
	"time"

	"github.com/mysteriumnetwork/payments/v3/client"
	"github.com/mysteriumnetwork/payments/v3/client/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_DefaultByBlockSize(t *testing.T) {
	cl := &mocks.EtherClientMock{
		BlockNumberFunc: func(ctx context.Context) (uint64, error) {
			return 2, nil
		},
	}
	cl2 := &mocks.EtherClientMock{
		BlockNumberFunc: func(ctx context.Context) (uint64, error) {
			return 1, nil
		},
	}
	cl3 := &mocks.EtherClientMock{
		BlockNumberFunc: func(ctx context.Context) (uint64, error) {
			return 3, nil
		},
	}
	getter := client.NewDefaultAddressableEthClientGetter("first", cl)
	getter2 := client.NewDefaultAddressableEthClientGetter("second", cl2)
	getter3 := client.NewDefaultAddressableEthClientGetter("third", cl3)

	client, err := client.NewEthMultiClient(time.Second, []client.AddressableEthClientGetter{getter, getter2, getter3})
	assert.NoError(t, err)

	currentOrder := client.CurrentClientOrder()
	assert.Equal(t, []string{"first", "second", "third"}, currentOrder)
	DefaultByBlockSize(client)
	newOrder := client.CurrentClientOrder()
	assert.Equal(t, []string{"third", "first", "second"}, newOrder)
}

func Test_DefaultByAvailability(t *testing.T) {
	cl1 := &mocks.EtherClientMock{}
	cl2 := &mocks.EtherClientMock{}
	getter := client.NewDefaultAddressableEthClientGetter("first", cl1)
	getter2 := client.NewDefaultAddressableEthClientGetter("second", cl2)

	cl, err := client.NewEthMultiClient(time.Second, []client.AddressableEthClientGetter{getter, getter2})
	assert.NoError(t, err)

	currentOrder := cl.CurrentClientOrder()
	assert.Equal(t, []string{"first", "second"}, currentOrder)
	DefaultByAvailability(client.Notification{
		Address: getter.Address(),
		Error:   errors.New("random error"),
	}, cl)

	// Order doesnt change if error is not connection error.
	newOrder := cl.CurrentClientOrder()
	assert.Equal(t, []string{"first", "second"}, newOrder)

	DefaultByAvailability(client.Notification{
		Address: getter.Address(),
		Error:   client.ErrClientNoConnection,
	}, cl)

	// Received connection error for first client, should push to last.
	newOrder = cl.CurrentClientOrder()
	assert.Equal(t, []string{"second", "first"}, newOrder)
}
