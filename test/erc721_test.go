package test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/v3/bindings"
	"github.com/stretchr/testify/assert"
)

func TestErc721(t *testing.T) {
	address, privateKey, err := GetKeyPair(privateKey0)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	assert.NoError(t, err)

	tokenAddress := common.Address{}

	t.Run("deploy", func(t *testing.T) {
		name := "ERC721"
		symbol := "ERC"

		tokenAddress, err = DeployErc721(GetTransactOpts(address, privateKey, chainID), client, name, symbol, 10*time.Second)
		assert.NoError(t, err)
		assert.NotEqual(t, (common.Address{}).Hex(), tokenAddress.Hex())

		erc721Caller, err := bindings.NewErc721Caller(tokenAddress, client)
		assert.NoError(t, err)

		nameFromContract, err := erc721Caller.Name(nil)
		assert.NoError(t, err)
		assert.Equal(t, name, nameFromContract)

		symbolFromContract, err := erc721Caller.Symbol(nil)
		assert.NoError(t, err)
		assert.Equal(t, symbol, symbolFromContract)

	})

	tokenId := big.NewInt(10)
	t.Run("mint", func(t *testing.T) {
		erc721Caller, err := bindings.NewErc721Caller(tokenAddress, client)
		assert.NoError(t, err)

		balanceFromContract, err := erc721Caller.BalanceOf(nil, address)
		assert.NoError(t, err)
		assert.Equal(t, 0, balanceFromContract.Cmp(big.NewInt(0)))

		err = MintErc721(GetTransactOpts(address, privateKey, chainID), client, tokenAddress, address, tokenId, 5*time.Second)
		assert.NoError(t, err)

		balanceFromContract, err = erc721Caller.BalanceOf(nil, address)
		assert.NoError(t, err)
		assert.Equal(t, 0, balanceFromContract.Cmp(big.NewInt(1)))

		owner, err := erc721Caller.OwnerOf(nil, tokenId)
		assert.NoError(t, err)
		assert.Equal(t, address.Hex(), owner.Hex())
	})

	t.Run("transfer", func(t *testing.T) {
		to := common.HexToAddress("0x123452222222")
		err := TransferErc721(GetTransactOpts(address, privateKey, chainID), client, tokenAddress, to, tokenId, 10*time.Second)
		assert.NoError(t, err)

		erc721Caller, err := bindings.NewErc721Caller(tokenAddress, client)
		assert.NoError(t, err)

		balanceFromContract, err := erc721Caller.BalanceOf(nil, address)
		assert.NoError(t, err)
		assert.Equal(t, 0, balanceFromContract.Cmp(big.NewInt(0)))

		balanceFromContract, err = erc721Caller.BalanceOf(nil, to)
		assert.NoError(t, err)
		assert.Equal(t, 0, balanceFromContract.Cmp(big.NewInt(1)))

		owner, err := erc721Caller.OwnerOf(nil, tokenId)
		assert.NoError(t, err)
		assert.Equal(t, to.Hex(), owner.Hex())
	})
}
