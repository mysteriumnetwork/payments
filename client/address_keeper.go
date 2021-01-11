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

package client

import "github.com/ethereum/go-ethereum/common"

// MultiChainAddressKeeper keeps track of smart contract addresses.
type MultiChainAddressKeeper struct {
	addresses map[int64]SmartContractAddresses
}

// SmartContractAddresses represents the mysterium smart contract address collection.
type SmartContractAddresses struct {
	Registry              common.Address
	Myst                  common.Address
	HermesImplementation  common.Address
	ChannelImplementation common.Address
	Hermes                common.Address
}

// NewMultiChainAddressKeeper creates a new instance of MultiChainAddressKeeper.
func NewMultiChainAddressKeeper(addresses map[int64]SmartContractAddresses) *MultiChainAddressKeeper {
	return &MultiChainAddressKeeper{
		addresses: addresses,
	}
}

func (mcak *MultiChainAddressKeeper) getAddressesForChain(chainID int64) (SmartContractAddresses, error) {
	if v, ok := mcak.addresses[chainID]; ok {
		return v, nil
	}
	return SmartContractAddresses{}, ErrUnknownChain
}

// GetAddressesForChain returns the SmartContractAddresses for the given chain.
func (mcak *MultiChainAddressKeeper) GetAddressesForChain(chainID int64) (SmartContractAddresses, error) {
	return mcak.getAddressesForChain(chainID)
}

// GetRegistryAddress returns the registry sc address for the given chain.
func (mcak *MultiChainAddressKeeper) GetRegistryAddress(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.Registry, nil
}

// GetMystAddress returns the myst sc address for the given chain.
func (mcak *MultiChainAddressKeeper) GetMystAddress(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.Myst, nil
}

// GetHermesImplementation returns the hermes implementation sc address for the given chain.
func (mcak *MultiChainAddressKeeper) GetHermesImplementation(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.HermesImplementation, nil
}

// GetChannelImplementation returns the channel implementation sc address for the given chain.
func (mcak *MultiChainAddressKeeper) GetChannelImplementation(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.ChannelImplementation, nil
}

// GetActiveHermes returns active hermes address for the given chain.
func (mcak *MultiChainAddressKeeper) GetActiveHermes(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.Hermes, nil
}
