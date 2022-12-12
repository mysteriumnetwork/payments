package client

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/mysteriumnetwork/payments/crypto"
)

type hermesImplementationRegistry struct {
	callers     map[string]*bindings.HermesImplementationCaller
	filterers   map[string]*bindings.HermesImplementationFilterer
	transactors map[string]*bindings.HermesImplementationTransactor
	lock        sync.RWMutex
}

func newHermesImplementationRegistry() *hermesImplementationRegistry {
	return &hermesImplementationRegistry{
		callers:     make(map[string]*bindings.HermesImplementationCaller),
		filterers:   make(map[string]*bindings.HermesImplementationFilterer),
		transactors: make(map[string]*bindings.HermesImplementationTransactor),
	}
}

func (h *hermesImplementationRegistry) caller(address crypto.Address, etherClient EtherClient) (*bindings.HermesImplementationCaller, error) {
	h.lock.RLock()
	caller, exists := h.callers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return caller, nil
	}

	caller, err := bindings.NewHermesImplementationCaller(address.ToCommon(), etherClient)
	if err != nil {
		return caller, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.callers[address.Hex()] = caller

	return caller, nil
}

func (h *hermesImplementationRegistry) filterer(address crypto.Address, filterer bind.ContractFilterer) (*bindings.HermesImplementationFilterer, error) {
	h.lock.RLock()
	filter, exists := h.filterers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return filter, nil
	}

	f, err := bindings.NewHermesImplementationFilterer(address.ToCommon(), filterer)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.filterers[address.Hex()] = f

	return f, nil
}

func (h *hermesImplementationRegistry) transactor(address crypto.Address, client EtherClient) (*bindings.HermesImplementationTransactor, error) {
	h.lock.RLock()
	transactor, exists := h.transactors[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return transactor, nil
	}

	f, err := bindings.NewHermesImplementationTransactor(address.ToCommon(), client)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.transactors[address.Hex()] = f

	return f, nil
}

type registry struct {
	callers     map[string]*bindings.RegistryCaller
	filterers   map[string]*bindings.RegistryFilterer
	transactors map[string]*bindings.RegistryTransactor
	lock        sync.RWMutex
}

func newRegistry() *registry {
	return &registry{
		callers:     make(map[string]*bindings.RegistryCaller),
		filterers:   make(map[string]*bindings.RegistryFilterer),
		transactors: make(map[string]*bindings.RegistryTransactor),
	}
}

func (h *registry) caller(address crypto.Address, etherClient EtherClient) (*bindings.RegistryCaller, error) {
	h.lock.RLock()
	caller, exists := h.callers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return caller, nil
	}

	caller, err := bindings.NewRegistryCaller(address.ToCommon(), etherClient)
	if err != nil {
		return caller, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.callers[address.Hex()] = caller

	return caller, nil
}

func (h *registry) filterer(address crypto.Address, filterer bind.ContractFilterer) (*bindings.RegistryFilterer, error) {
	h.lock.RLock()
	filter, exists := h.filterers[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return filter, nil
	}

	f, err := bindings.NewRegistryFilterer(address.ToCommon(), filterer)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.filterers[address.Hex()] = f

	return f, nil
}

func (h *registry) transactor(address crypto.Address, client EtherClient) (*bindings.RegistryTransactor, error) {
	h.lock.RLock()
	transactor, exists := h.transactors[address.Hex()]
	h.lock.RUnlock()
	if exists {
		return transactor, nil
	}

	f, err := bindings.NewRegistryTransactor(address.ToCommon(), client)
	if err != nil {
		return f, err
	}

	h.lock.Lock()
	defer h.lock.Unlock()
	h.transactors[address.Hex()] = f

	return f, nil
}
