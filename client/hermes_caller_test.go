package client

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestHermesCaller(t *testing.T) {
	hcr := NewHermesImplementationCallerRegistry()
	address := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec5")
	hc, err := hcr.Get(address, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	caller, ok := hcr.callers[address.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hc, caller)
}
