package test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/v3/bindings"
	"github.com/stretchr/testify/assert"
)

func TestDeployHermesImplementation(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	hermesImplementationAddress, err := DeployHermesImplementation(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), hermesImplementationAddress.Hex())

	hermesImplementationCaller, err := bindings.NewHermesImplementation(hermesImplementationAddress, client)
	assert.NoError(t, err)

	owner, err := hermesImplementationCaller.Owner(nil)
	assert.NoError(t, err)
	assert.Equal(t, (common.Address{}).Hex(), owner.Hex())
}
