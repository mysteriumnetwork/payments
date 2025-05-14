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

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/v3/crypto"
)

// MultiChainAddressKeeper keeps track of smart contract addresses.
type MultiChainAddressKeeper struct {
	defaultAddresses map[int64]SmartContractAddresses
}

type apBlockchain interface {
	GetHermes(chainID int64, registryID, hermesID common.Address) (Hermes, error)
	GetChannelImplementationByVersion(chainID int64, registryID common.Address, version *big.Int) (common.Address, error)
}

// SmartContractAddresses represents the mysterium smart contract address collection.
type SmartContractAddresses struct {
	Registry                    common.Address
	Myst                        common.Address
	ActiveHermes                common.Address
	ActiveChannelImplementation common.Address
	KnownHermeses               []common.Address
}

// NewMultiChainAddressKeeper creates a new instance of MultiChainAddressKeeper.
func NewMultiChainAddressKeeper(addresses map[int64]SmartContractAddresses) *MultiChainAddressKeeper {
	return &MultiChainAddressKeeper{
		defaultAddresses: addresses,
	}
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

// GetActiveHermes returns active hermes address for the given chain.
func (mcak *MultiChainAddressKeeper) GetActiveHermes(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.ActiveHermes, nil
}

// GetActiveChannelImplementation returns channel implementation a default address for the given chain.
func (mcak *MultiChainAddressKeeper) GetActiveChannelImplementation(chainID int64) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}
	return v.ActiveChannelImplementation, nil
}

// GetKnownHermeses returns a list of all the known hermes addresses for the given chain.
func (mcak *MultiChainAddressKeeper) GetKnownHermeses(chainID int64) ([]common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return nil, err
	}
	return v.KnownHermeses, nil
}

// GetActiveChannelAddress will calculate a channel address for an identity with active hermes, registry and
// active channel implementation values for a given chain.
func (mcak *MultiChainAddressKeeper) GetActiveChannelAddress(chainID int64, id common.Address) (common.Address, error) {
	v, err := mcak.getAddressesForChain(chainID)
	if err != nil {
		return common.Address{}, err
	}

	addr, err := crypto.GenerateChannelAddress(id.Hex(), v.ActiveHermes.Hex(), v.Registry.Hex(), v.ActiveChannelImplementation.Hex())
	return common.HexToAddress(addr), err
}

// GetArbitraryChannelAddress will will calculate a consumer channel address given all custom parameters.
func (mcak *MultiChainAddressKeeper) GetArbitraryChannelAddress(hermes, registry, channel common.Address, id common.Address) (common.Address, error) {
	addr, err := crypto.GenerateChannelAddress(id.Hex(), hermes.Hex(), registry.Hex(), channel.Hex())
	return common.HexToAddress(addr), err
}

func (mcak *MultiChainAddressKeeper) getAddressesForChain(chainID int64) (SmartContractAddresses, error) {
	if v, ok := mcak.defaultAddresses[chainID]; ok {
		return v, nil
	}
	return SmartContractAddresses{}, ErrUnknownChain
}

// MultiChainAddressKeeper keeps track of smart contract addresses.
type MultiChainAddressProvider struct {
	*MultiChainAddressKeeper
	bc apBlockchain

	// Getting channel implementation is expensive so we want to keep
	// any values that we retrieve, they will not change anyway.
	chCache map[string]common.Address
	m       sync.RWMutex
}

// NewMultiChainAddressProvider creates a new instance of MultiChainAddressKeeper.
func NewMultiChainAddressProvider(k *MultiChainAddressKeeper, bc apBlockchain) *MultiChainAddressProvider {
	return &MultiChainAddressProvider{
		MultiChainAddressKeeper: k,
		bc:                      bc,
		chCache:                 make(map[string]common.Address),
	}
}

// GetChannelImplementationForHermes for hermes will query the blockchain and find a channel implementation address
// that is used by a given hermes. It will then cache the value.
func (mcak *MultiChainAddressProvider) GetChannelImplementationForHermes(chainID int64, hermes common.Address) (common.Address, error) {
	got, ok := mcak.chCacheGet(chainID, hermes)
	if ok {
		return got, nil
	}

	registry, err := mcak.GetRegistryAddress(chainID)
	if err != nil {
		return common.Address{}, err
	}

	bcHermes, err := mcak.bc.GetHermes(chainID, registry, hermes)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not get hermes from BC: %w", err)
	}

	ch, err := mcak.bc.GetChannelImplementationByVersion(chainID, registry, bcHermes.ImplVer)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not get ch implementation from BC: %w", err)
	}

	mcak.chCacheSet(chainID, hermes, ch)
	return ch, nil
}

func (mcak *MultiChainAddressProvider) chCacheSet(chainID int64, hermes, impl common.Address) {
	key := mcak.chCacheKey(chainID, hermes)

	mcak.m.Lock()
	defer mcak.m.Unlock()

	mcak.chCache[key] = impl
}

func (mcak *MultiChainAddressProvider) chCacheGet(chainID int64, hermes common.Address) (common.Address, bool) {
	key := mcak.chCacheKey(chainID, hermes)

	mcak.m.RLock()
	defer mcak.m.RUnlock()

	got, ok := mcak.chCache[key]
	return got, ok
}

func (mcak *MultiChainAddressKeeper) chCacheKey(chainID int64, hermes common.Address) string {
	return fmt.Sprintf("%d|%s", chainID, hermes.Hex())
}

// GetHermesChannelAddress will calculate a channel address for an identity and hermes using the channel
// implementation of that hermes and the registry for a given chain.
func (mcak *MultiChainAddressProvider) GetHermesChannelAddress(chainID int64, id, hermesAddr common.Address) (common.Address, error) {
	activeHermesAddr, err := mcak.GetActiveHermes(chainID)
	if err != nil {
		return common.Address{}, err
	}
	var channelImplementation common.Address
	if hermesAddr == activeHermesAddr {
		return mcak.GetActiveChannelAddress(chainID, id)
	}
	channelImplementation, err = mcak.GetChannelImplementationForHermes(chainID, hermesAddr)
	if err != nil {
		return common.Address{}, err
	}
	registry, err := mcak.GetRegistryAddress(chainID)
	if err != nil {
		return common.Address{}, err
	}

	addr, err := crypto.GenerateChannelAddress(id.Hex(), hermesAddr.Hex(), registry.Hex(), channelImplementation.Hex())
	return common.HexToAddress(addr), err
}
