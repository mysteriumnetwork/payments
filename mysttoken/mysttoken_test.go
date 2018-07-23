package mysttoken

import (
	"math/big"
	"testing"

	"github.com/MysteriumNetwork/payments/mysttoken/generated"
	"github.com/MysteriumNetwork/payments/test_utils"
	"github.com/stretchr/testify/assert"
)

var abiMap = test_utils.AbiMap{
	"MystToken": {
		generated.MystTokenABI,
		generated.MystTokenBin,
	},
}

func TestSampleERC20TokenIsDeployedWithSpecifiedTokenAmount(t *testing.T) {
	abiList, err := test_utils.ParseAbis(abiMap)
	assert.NoError(t, err)

	backend := test_utils.LoggingBackend(test_utils.NewSimulatedBackend(test_utils.Deployer.Address, 100000000), abiList)

	erc20, err := DeployMystERC20(test_utils.Deployer.Transactor, 1000000, backend)
	backend.Commit() //do we need this?
	assert.NoError(t, err)

	balance, err := erc20.BalanceOf(test_utils.Deployer.Address)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(1000000), balance)

}
