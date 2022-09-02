package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
)

type HermesImplementationCallerRegistry struct {
	callers map[string]*bindings.HermesImplementationCaller
}

func NewHermesImplementationCallerRegistry() *HermesImplementationCallerRegistry {
	return &HermesImplementationCallerRegistry{callers: make(map[string]*bindings.HermesImplementationCaller)}
}

func (h *HermesImplementationCallerRegistry) Get(hermesID common.Address, etherClient EtherClient) (*bindings.HermesImplementationCaller, error) {
	if caller, exists := h.callers[hermesID.Hex()]; exists {
		return caller, nil
	}

	caller, err := bindings.NewHermesImplementationCaller(hermesID, etherClient)
	if err != nil {
		return caller, err
	}

	h.callers[hermesID.Hex()] = caller

	return caller, nil
}
