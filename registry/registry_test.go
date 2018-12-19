package registry

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/mysttoken"
	"github.com/mysteriumnetwork/payments/test_utils"
	"github.com/stretchr/testify/assert"
)

var abiMap = test_utils.AbiMap{
	"MystToken": {
		mysttoken.MystTokenABI,
		mysttoken.MystTokenBin,
	},
	"IdentityRegistry": {
		IdentityRegistryABI,
		IdentityRegistryBin,
	},
}

var abiList, _ = test_utils.ParseAbis(abiMap)

func TestRegistryIsDeployable(t *testing.T) {
	backend := test_utils.LoggingBackend(test_utils.NewSimulatedBackend(test_utils.Deployer.Address, 1000000000), abiList)

	registry, err := DeployRegistry(test_utils.Deployer.Transactor, common.Address{}, backend)
	assert.NoError(t, err)
	backend.Commit()

	registered, err := registry.IsRegistered(test_utils.Deployer.Address)
	assert.NoError(t, err)
	assert.False(t, registered)
}

func TestRegisterIdentityEmitsIdentityRegisteredEvent(t *testing.T) {
	backend := test_utils.LoggingBackend(test_utils.NewSimulatedBackend(test_utils.Deployer.Address, 1000000000), abiList)

	mystERC20, err := mysttoken.DeployMystERC20(test_utils.Deployer.Transactor, 1000000, backend)
	assert.NoError(t, err)
	backend.Commit()

	registry, err := DeployRegistry(test_utils.Deployer.Transactor, mystERC20.Address, backend)
	assert.NoError(t, err)
	backend.Commit()

	_, err = mystERC20.Approve(registry.Address, big.NewInt(3000)) // allowance 3000 - fee (1000) = should be left 2000
	assert.NoError(t, err)

	mystIdentity, err := test_utils.NewMystIdentity()
	assert.NoError(t, err)

	fromBlock := uint64(0)
	//funny thing with simulated backed - we need subscribe to events BEFORE they are generated - limitation of simulator :)
	eventChan := make(chan *IdentityRegistryRegistered)
	subscription, err := registry.Contract.WatchRegistered(&bind.WatchOpts{Start: &fromBlock}, eventChan, []common.Address{})
	assert.NoError(t, err)

	registered, err := registry.IsRegistered(mystIdentity.Address)
	assert.NoError(t, err)
	assert.False(t, registered)

	proofOfIdentity, err := CreateRegistrationData(mystIdentity)
	assert.NoError(t, err)

	_, err = registry.RegisterIdentity(proofOfIdentity)
	assert.NoError(t, err)
	backend.Commit()

	registryContractBalance, err := mystERC20.BalanceOf(registry.Address)
	assert.Equal(t, big.NewInt(1000), registryContractBalance) //fee is received for registry contract

	registered, err = registry.IsRegistered(mystIdentity.Address)
	assert.NoError(t, err)
	assert.True(t, registered)

	pubKey, err := registry.LookupPublicKey(mystIdentity.Address)
	assert.NoError(t, err)
	assert.Equal(t, *(mystIdentity.PublicKey), *pubKey)

	feeReceiver, err := test_utils.NewRandomAccount()
	assert.NoError(t, err)

	_, err = registry.TransferCollectedFeeTo(feeReceiver.Address)
	assert.NoError(t, err)
	backend.Commit()

	feeReceiverBalance, err := mystERC20.BalanceOf(feeReceiver.Address)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(1000), feeReceiverBalance)

	select {
	case event := <-eventChan:
		assert.Equal(t, mystIdentity.Address, event.Identity)
	case err := <-subscription.Err():
		assert.NoError(t, err)
	case <-time.After(100 * time.Millisecond):
		assert.Fail(t, "Identity registered event expected")
	}
	subscription.Unsubscribe()
}
