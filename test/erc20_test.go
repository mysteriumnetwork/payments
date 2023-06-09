package test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/mysteriumnetwork/payments/units"
	"github.com/stretchr/testify/assert"
)

func TestDeployErc20(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	name := "ERC20"
	symbol := "ERC"
	totalSupply := units.FloatEthToBigIntWei(1000)

	tokenAddress, err := DeployErc20(GetTransactOpts(address, privateKey, chainID), client, name, symbol, totalSupply, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), tokenAddress.Hex())

	erc20Caller, err := bindings.NewErc20Caller(tokenAddress, client)
	assert.NoError(t, err)

	decimals, err := erc20Caller.Decimals(nil)
	assert.NoError(t, err)
	assert.NotZero(t, decimals)

	nameFromContract, err := erc20Caller.Name(nil)
	assert.NoError(t, err)
	assert.Equal(t, name, nameFromContract)

	symbolFromContract, err := erc20Caller.Symbol(nil)
	assert.NoError(t, err)
	assert.Equal(t, symbol, symbolFromContract)

	totalSupplyFromContract, err := erc20Caller.TotalSupply(nil)
	assert.NoError(t, err)
	assert.Equal(t, totalSupply, totalSupplyFromContract)

	balanceFromContract, err := erc20Caller.BalanceOf(nil, address)
	assert.NoError(t, err)
	assert.Equal(t, totalSupply, balanceFromContract)
}
