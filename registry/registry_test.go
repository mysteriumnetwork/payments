package registry

import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/MysteriumNetwork/payments/registry/generated"
	"fmt"
	"github.com/MysteriumNetwork/payments/test_utils"
	"github.com/MysteriumNetwork/payments/mysttoken"
)

func TestRegistryIsDeployable(t *testing.T) {
	backend := test_utils.NewSimulatedBackend(test_utils.Deployer.Address, 1000000)

	registry, err := DeployRegistry(test_utils.Deployer.Transactor , common.Address{} , backend)
	assert.NoError(t ,err)
	backend.Commit()

	registered , err := registry.IsRegistered(test_utils.Deployer.Address)
	assert.NoError(t, err)
	assert.False(t, registered)
}

func TestRegisterIdentityEmitsIdentityRegisteredEvent(t *testing.T) {
	backend := test_utils.NewSimulatedBackend(test_utils.Deployer.Address, 100000000000000)

	mystERC20, err := mysttoken.DeployMystERC20(test_utils.Deployer.Transactor , 1000000, backend)
	assert.NoError(t ,err)
	backend.Commit()

	registry, err := DeployRegistry(test_utils.Deployer.Transactor , mystERC20.Address , backend)
	assert.NoError(t ,err)
	backend.Commit()

	mystIdentity, err := NewMystIdentity()
	assert.NoError(t, err)

	fromBlock := uint64(0)
	//funny thing with simulated backed - we need subscribe to events BEFORE they are generated - limitation of simulator :)
	eventChan := make( chan *generated.IdentityRegistryRegistered)
	subscription, err := registry.Contract.WatchRegistered( &bind.WatchOpts{ Start: &fromBlock} , eventChan, []common.Address{})
	assert.NoError(t , err)

	registered, err := registry.IsRegistered(mystIdentity.Address)
	assert.NoError(t, err)
	assert.False(t, registered)

	proofOfIdentity,err  := CreateProofOfIdentity(mystIdentity)
	assert.NoError(t, err)

	tx , err := registry.RegisterIdentity(proofOfIdentity.RandomNumber , proofOfIdentity.Signature)
	fmt.Printf("Transaction: %+v\n" , tx)
	assert.NoError(t, err)
	backend.Commit()

	registered, err = registry.IsRegistered(mystIdentity.Address)
	assert.NoError(t, err)
	assert.True(t, registered)


	select {
		case event:= <- eventChan :
			assert.Equal(t , mystIdentity.Address, event.Identity)
		case err:= <- subscription.Err() :
			assert.NoError(t , err)
	}
	subscription.Unsubscribe()
}