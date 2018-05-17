package mysttoken

import (
	"testing"
	"math/big"
	"github.com/stretchr/testify/assert"
	"github.com/MysteriumNetwork/payments/test_utils"
)


func TestSampleERC20TokenIsDeployedWithSpecifiedTokenAmount(t *testing.T) {
	backend := test_utils.NewSimulatedBackend(test_utils.Deployer.Address , 100000000)

	erc20 , err := DeployMystERC20(test_utils.Deployer.Transactor , 1000000 , backend)
	backend.Commit() //do we need this?
	assert.NoError(t , err)

	balance, err := erc20.BalanceOf(test_utils.Deployer.Address)
	assert.NoError(t, err)
	assert.Equal(t , big.NewInt(1000000) , balance)

}