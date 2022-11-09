package client

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestHermesImplementation(t *testing.T) {
	hcr := newHermesImplementationRegistry()

	callerAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec1")
	hc, err := hcr.caller(callerAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	caller, ok := hcr.callers[callerAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hc, caller)

	transactorAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec2")
	ht, err := hcr.transactor(transactorAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	transactor, ok := hcr.transactors[transactorAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, ht, transactor)

	filtererAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec3")
	hf, err := hcr.filterer(filtererAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	filterer, ok := hcr.filterers[filtererAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hf, filterer)
}

func TestMystTokenImplementation(t *testing.T) {
	hcr := newMystTokenRegistry()

	callerAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec1")
	hc, err := hcr.caller(callerAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	caller, ok := hcr.callers[callerAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hc, caller)

	transactorAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec2")
	ht, err := hcr.transactor(transactorAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	transactor, ok := hcr.transactors[transactorAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, ht, transactor)

	filtererAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec3")
	hf, err := hcr.filterer(filtererAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	filterer, ok := hcr.filterers[filtererAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hf, filterer)
}

func TestChannellImplementation(t *testing.T) {
	hcr := newChannelImplementationRegistry()

	callerAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec1")
	hc, err := hcr.caller(callerAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	caller, ok := hcr.callers[callerAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hc, caller)

	transactorAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec2")
	ht, err := hcr.transactor(transactorAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	transactor, ok := hcr.transactors[transactorAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, ht, transactor)

	filtererAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec3")
	hf, err := hcr.filterer(filtererAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	filterer, ok := hcr.filterers[filtererAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hf, filterer)
}

func TestRegistry(t *testing.T) {
	hcr := newRegistry()

	callerAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec1")
	hc, err := hcr.caller(callerAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	caller, ok := hcr.callers[callerAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hc, caller)

	transactorAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec2")
	ht, err := hcr.transactor(transactorAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	transactor, ok := hcr.transactors[transactorAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, ht, transactor)

	filtererAddress := common.HexToAddress("0xEA672fdDe714fd979de3EdF0F56AA9716B898ec3")
	hf, err := hcr.filterer(filtererAddress, nil)
	assert.NoError(t, err)
	assert.Len(t, hcr.callers, 1)

	filterer, ok := hcr.filterers[filtererAddress.Hex()]
	assert.True(t, ok)
	assert.Equal(t, hf, filterer)
}
