package test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/bindings"
	"github.com/stretchr/testify/assert"
)

func TestDeployErc721(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	name := "ERC721"
	symbol := "ERC"

	tokenAddress, err := DeployErc721(GetTransactOpts(address, privateKey, chainID), client, name, symbol, 10*time.Second)
	assert.NoError(t, err)
	assert.NotEqual(t, (common.Address{}).Hex(), tokenAddress.Hex())

	erc721, err := bindings.NewErc721(tokenAddress, client)
	assert.NoError(t, err)

	nameFromContract, err := erc721.Name(nil)
	assert.NoError(t, err)
	assert.Equal(t, name, nameFromContract)

	symbolFromContract, err := erc721.Symbol(nil)
	assert.NoError(t, err)
	assert.Equal(t, symbol, symbolFromContract)

	tokenId := big.NewInt(10)
	_, err = erc721.Mint(GetTransactOpts(address, privateKey, chainID), address, tokenId)
	assert.NoError(t, err)

	balanceFromContract, err := erc721.BalanceOf(nil, address)
	assert.NoError(t, err)
	assert.Equal(t, 0, balanceFromContract.Cmp(big.NewInt(1)))

	owner, err := erc721.OwnerOf(nil, tokenId)
	assert.NoError(t, err)
	assert.Equal(t, address.Hex(), owner.Hex())
}
