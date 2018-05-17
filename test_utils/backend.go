package test_utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"math/big"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/common"
)

// Keep in mind that initial amount is actually in Gwei - you need really HUGE
// numbers here to make more than few simple transactions
func NewSimulatedBackend(genesisAddress common.Address, initialAmmount int64) (* backends.SimulatedBackend) {
	return backends.NewSimulatedBackend(core.GenesisAlloc{
		genesisAddress: {Balance: big.NewInt(initialAmmount)},
	})
}