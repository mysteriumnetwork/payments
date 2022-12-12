/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
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

package client

import (
	"context"
	"sync"

	"github.com/mysteriumnetwork/payments/crypto"
)

// NonceTracker keeps track of nonces atomically.
type NonceTracker struct {
	client    client
	nonces    map[crypto.Address]uint64
	nonceLock sync.Mutex
}

type client interface {
	PendingNonceAt(ctx context.Context, account crypto.Address) (uint64, error)
}

// NewNonceTracker returns a new nonce tracker.
func NewNonceTracker(client client) *NonceTracker {
	return &NonceTracker{
		client: client,
		nonces: make(map[crypto.Address]uint64),
	}
}

// GetNonce returns an atomically increasing nonce for the account.
func (nt *NonceTracker) GetNonce(ctx context.Context, account crypto.Address) (uint64, error) {
	nt.nonceLock.Lock()
	defer nt.nonceLock.Unlock()

	if v, ok := nt.nonces[account]; ok {
		v++
		nt.nonces[account] = v
		return v, nil
	}

	nonce, err := nt.client.PendingNonceAt(ctx, account)
	if err != nil {
		return nonce, err
	}

	nt.nonces[account] = nonce
	return nonce, nil
}

// ForceReloadNonce clears the nonce cache. This will force loading from BC next time.
func (nt *NonceTracker) ForceReloadNonce(account crypto.Address) {
	nt.nonceLock.Lock()
	defer nt.nonceLock.Unlock()
	delete(nt.nonces, account)
}
