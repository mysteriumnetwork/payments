package client

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
)

type HermesImplementationRegistry struct {
	callers     map[string]*bindings.HermesImplementationCaller
	filterers   map[string]*bindings.HermesImplementationFilterer
	transactors map[string]*bindings.HermesImplementationTransactor
	lock        sync.RWMutex
}

func newHermesImplementationRegistry() *HermesImplementationRegistry {
	return &HermesImplementationRegistry{
		callers:     make(map[string]*bindings.HermesImplementationCaller),
		filterers:   make(map[string]*bindings.HermesImplementationFilterer),
		transactors: make(map[string]*bindings.HermesImplementationTransactor),
	}
}

func (h *HermesImplementationRegistry) caller(address common.Address, etherClient EtherClient) (*bindings.HermesImplementationCaller, error) {
	h.lock.RLock()
	caller, exists := h.callers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return caller, nil
	}

	caller, err := bindings.NewHermesImplementationCaller(address, etherClient)
	if err != nil {
		return caller, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.callers[address.Hex()] = caller

	return caller, nil
}

func (h *HermesImplementationRegistry) filterer(address common.Address, filterer bind.ContractFilterer) (*bindings.HermesImplementationFilterer, error) {
	h.lock.RLock()
	filter, exists := h.filterers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return filter, nil
	}

	f, err := bindings.NewHermesImplementationFilterer(address, filterer)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.filterers[address.Hex()] = f

	return f, nil
}

func (h *HermesImplementationRegistry) transactor(address common.Address, client EtherClient) (*bindings.HermesImplementationTransactor, error) {
	h.lock.RLock()
	transactor, exists := h.transactors[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return transactor, nil
	}

	f, err := bindings.NewHermesImplementationTransactor(address, client)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.transactors[address.Hex()] = f

	return f, nil
}

type Registry struct {
	callers     map[string]*bindings.RegistryCaller
	filterers   map[string]*bindings.RegistryFilterer
	transactors map[string]*bindings.RegistryTransactor
	lock        sync.RWMutex
}

func newRegistry() *Registry {
	return &Registry{
		callers:     make(map[string]*bindings.RegistryCaller),
		filterers:   make(map[string]*bindings.RegistryFilterer),
		transactors: make(map[string]*bindings.RegistryTransactor),
	}
}

func (h *Registry) caller(address common.Address, etherClient EtherClient) (*bindings.RegistryCaller, error) {
	h.lock.RLock()
	caller, exists := h.callers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return caller, nil
	}

	caller, err := bindings.NewRegistryCaller(address, etherClient)
	if err != nil {
		return caller, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.callers[address.Hex()] = caller

	return caller, nil
}

func (h *Registry) filterer(address common.Address, filterer bind.ContractFilterer) (*bindings.RegistryFilterer, error) {
	h.lock.RLock()
	filter, exists := h.filterers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return filter, nil
	}

	f, err := bindings.NewRegistryFilterer(address, filterer)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.filterers[address.Hex()] = f

	return f, nil
}

func (h *Registry) transactor(address common.Address, client EtherClient) (*bindings.RegistryTransactor, error) {
	h.lock.RLock()
	transactor, exists := h.transactors[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return transactor, nil
	}

	f, err := bindings.NewRegistryTransactor(address, client)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.transactors[address.Hex()] = f

	return f, nil
}

type MystTokenRegistry struct {
	callers     map[string]*bindings.MystTokenCaller
	filterers   map[string]*bindings.MystTokenFilterer
	transactors map[string]*bindings.MystTokenTransactor
	lock        sync.RWMutex
}

func newMystTokenRegistry() *MystTokenRegistry {
	return &MystTokenRegistry{
		callers:     make(map[string]*bindings.MystTokenCaller),
		filterers:   make(map[string]*bindings.MystTokenFilterer),
		transactors: make(map[string]*bindings.MystTokenTransactor),
	}
}

func (h *MystTokenRegistry) caller(address common.Address, etherClient EtherClient) (*bindings.MystTokenCaller, error) {
	h.lock.RLock()
	caller, exists := h.callers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return caller, nil
	}

	caller, err := bindings.NewMystTokenCaller(address, etherClient)
	if err != nil {
		return caller, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.callers[address.Hex()] = caller

	return caller, nil
}

func (h *MystTokenRegistry) filterer(address common.Address, filterer bind.ContractFilterer) (*bindings.MystTokenFilterer, error) {
	h.lock.RLock()
	filter, exists := h.filterers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return filter, nil
	}

	f, err := bindings.NewMystTokenFilterer(address, filterer)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.filterers[address.Hex()] = f

	return f, nil
}

func (h *MystTokenRegistry) transactor(address common.Address, client EtherClient) (*bindings.MystTokenTransactor, error) {
	h.lock.RLock()
	transactor, exists := h.transactors[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return transactor, nil
	}

	f, err := bindings.NewMystTokenTransactor(address, client)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.transactors[address.Hex()] = f

	return f, nil
}

type ChannelImplementationRegistry struct {
	callers     map[string]*bindings.ChannelImplementationCaller
	filterers   map[string]*bindings.ChannelImplementationFilterer
	transactors map[string]*bindings.ChannelImplementationTransactor
	lock        sync.RWMutex
}

func newChannelImplementationRegistry() *ChannelImplementationRegistry {
	return &ChannelImplementationRegistry{
		callers:     make(map[string]*bindings.ChannelImplementationCaller),
		filterers:   make(map[string]*bindings.ChannelImplementationFilterer),
		transactors: make(map[string]*bindings.ChannelImplementationTransactor),
	}
}

func (h *ChannelImplementationRegistry) caller(address common.Address, etherClient EtherClient) (*bindings.ChannelImplementationCaller, error) {
	h.lock.RLock()
	caller, exists := h.callers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return caller, nil
	}

	caller, err := bindings.NewChannelImplementationCaller(address, etherClient)
	if err != nil {
		return caller, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.callers[address.Hex()] = caller

	return caller, nil
}

func (h *ChannelImplementationRegistry) filterer(address common.Address, filterer bind.ContractFilterer) (*bindings.ChannelImplementationFilterer, error) {
	h.lock.RLock()
	filter, exists := h.filterers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return filter, nil
	}

	f, err := bindings.NewChannelImplementationFilterer(address, filterer)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.filterers[address.Hex()] = f

	return f, nil
}

func (h *ChannelImplementationRegistry) transactor(address common.Address, client EtherClient) (*bindings.ChannelImplementationTransactor, error) {
	h.lock.RLock()
	transactor, exists := h.transactors[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return transactor, nil
	}

	f, err := bindings.NewChannelImplementationTransactor(address, client)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.transactors[address.Hex()] = f

	return f, nil
}
