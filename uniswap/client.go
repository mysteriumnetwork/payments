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

package uniswap

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mysteriumnetwork/payments/uniswap/bindings"
)

// Client allows to do operations on uniswap smart contracts.
type Client struct {
	bc *ethclient.Client
}

// NewClient returns a new instance of uniswap client.
func NewClient(bc *ethclient.Client) *Client {
	return &Client{
		bc: bc,
	}
}

// GetExchangeAmount returns the amount of tokens you'd receive when exchanging the given amount of token0 to token1.
func (c Client) GetExchangeAmount(amount *big.Int, token0, token1 common.Address) (*big.Int, error) {
	addr := GeneratePairAddress(token0, token1)
	caller, err := bindings.NewIUniswapV2PairCaller(addr, c.bc)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	// This is the tricky bit.
	// The reserve call returns the reserves for token0 and token1 in a sorted order.
	// This means we need to check if our token addresses are sorted or not and flip the reserves if they are not sorted.
	stoken0, _ := sortAddressess(token0, token1)
	if stoken0 != token0 {
		// We're not sorted, so the reserves need to be flipped to represent the actual reserves.
		reserves.Reserve0, reserves.Reserve1 = reserves.Reserve1, reserves.Reserve0
	}

	return Quote(amount, reserves.Reserve0, reserves.Reserve1), nil
}
