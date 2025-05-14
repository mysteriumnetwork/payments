package test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/v3/bindings"
	"github.com/mysteriumnetwork/payments/v3/units"
	"github.com/stretchr/testify/assert"
)

func TestDeployOldToken(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	tokenAddress, err := DeployOldToken(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), tokenAddress.Hex())

	oldTokenCaller, err := bindings.NewOldMystToken(tokenAddress, client)
	assert.NoError(t, err)

	decimals, err := oldTokenCaller.Decimals(nil)
	assert.NoError(t, err)
	assert.NotZero(t, decimals)
}

func TestDeployTokenV2(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	tokenAddress, err := DeployOldToken(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), tokenAddress.Hex())

	tokenv2Address, err := DeployTokenV2(GetTransactOpts(address, privateKey, chainID), client, tokenAddress, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), tokenv2Address.Hex())

	tokenCaller, err := bindings.NewMystToken(tokenv2Address, client)
	assert.NoError(t, err)

	symbol, err := tokenCaller.Symbol(nil)
	assert.NoError(t, err)
	assert.Equal(t, "MYST", symbol)
}

func TestDeployTokens(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	tokenAddresses, err := DeployTokenV2WithDependencies(GetTransactOpts(address, privateKey, chainID), client, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), tokenAddresses.TokenAddress.Hex())
	assert.NotEqual(t, (common.Address{}).Hex(), tokenAddresses.TokenV2Address.Hex())

	oldTokenCaller, err := bindings.NewOldMystToken(tokenAddresses.TokenAddress, client)
	assert.NoError(t, err)

	decimals, err := oldTokenCaller.Decimals(nil)
	assert.NoError(t, err)
	assert.NotZero(t, decimals)

	tokenCaller, err := bindings.NewMystToken(tokenAddresses.TokenV2Address, client)
	assert.NoError(t, err)

	symbol, err := tokenCaller.Symbol(nil)
	assert.NoError(t, err)
	assert.Equal(t, "MYST", symbol)
}

func TestMintTokens(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	topts := GetTransactOpts(address, privateKey, chainID)
	tokenAddresses, err := DeployTokenV2WithDependencies(topts, client, 10*time.Second)
	assert.NoError(t, err)

	key, err := crypto.GenerateKey()
	assert.NoError(t, err)
	recipient := crypto.PubkeyToAddress(key.PublicKey)

	tokenCaller, err := bindings.NewMystToken(tokenAddresses.TokenV2Address, client)
	assert.NoError(t, err)

	balance, err := tokenCaller.BalanceOf(nil, recipient)
	assert.NoError(t, err)
	assert.Equal(t, units.FloatEthToBigIntWei(0).String(), balance.String())

	err = MintTokens(topts, tokenAddresses.TokenV2Address, client, 10*time.Second, recipient, units.FloatEthToBigIntWei(10))
	assert.NoError(t, err)

	balance, err = tokenCaller.BalanceOf(nil, recipient)
	assert.NoError(t, err)
	assert.Equal(t, units.FloatEthToBigIntWei(10).String(), balance.String())
}
