package sample_token

import (
	"testing"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"github.com/stretchr/testify/assert"
)

var (
	deployerPrivateKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	deployerAddress       =  crypto.PubkeyToAddress(deployerPrivateKey.PublicKey)
	deployerTransactor = bind.NewKeyedTransactor(deployerPrivateKey)
)


func TestSampleERC20TokenIsDeployedWithSpecifiedTokenAmount(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		deployerAddress: {Balance: big.NewInt(1000000000)},
	})

	erc20 , err := DeployERC20Token(deployerTransactor , 1000000 , backend)
	backend.Commit() //do we need this?
	assert.NoError(t , err)

	balance, err := erc20.BalanceOf(deployerAddress)
	assert.NoError(t, err)
	assert.Equal(t , big.NewInt(1000000) , balance)

}