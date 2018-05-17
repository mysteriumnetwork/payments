package registry

import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/MysteriumNetwork/payments/registry/generated"
	"fmt"
)

var (
	deployerPrivateKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	deployerAddress       =  crypto.PubkeyToAddress(deployerPrivateKey.PublicKey)
	deployerTransactor = bind.NewKeyedTransactor(deployerPrivateKey)
)

func TestRegistryIsDeployable(t *testing.T) {
	registry, _ := deployRegistry(t)

	registered , err := registry.IsRegistered( nil , deployerAddress)
	assert.NoError(t, err)
	assert.False(t, registered)
}

func TestRegisterIdentityEmitsIdentityRegisteredEvent(t *testing.T) {
	registry, backend := deployRegistry(t)

	mystIdentity, err := NewMystIdentity()
	assert.NoError(t, err)

	fromBlock := uint64(0)
	//funny thing with simulated backed - we need subscribe to events BEFORE they are generated - limitation of simulator :)
	eventChan := make( chan *generated.IdentityRegistryRegistered)
	subscription, err := registry.WatchRegistered( &bind.WatchOpts{ Start: &fromBlock} , eventChan, []common.Address{})
	assert.NoError(t , err)

	registered, err := registry.IsRegistered(nil, mystIdentity.Address)
	assert.NoError(t, err)
	assert.False(t, registered)

	proofOfIdentity,err  := CreateProofOfIdentity(mystIdentity)
	assert.NoError(t, err)

	tx , err := registry.RegisterIdentity(deployerTransactor , proofOfIdentity.RandomNumber , proofOfIdentity.Signature)
	fmt.Printf("Transaction: %+v\n" , tx)
	assert.NoError(t, err)
	backend.Commit()

	registered, err = registry.IsRegistered(nil, mystIdentity.Address)
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

func deployRegistry(t * testing.T) (*generated.IdentityRegistry , * backends.SimulatedBackend) {
	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		deployerAddress: {Balance: big.NewInt(1000000000)},
	})
	_ , _, registry, err := generated.DeployIdentityRegistry(deployerTransactor, simulatedBackend)
	simulatedBackend.Commit()
	if err != nil {
		assert.FailNow(t, "Unexpected error: %s", err)
	}
	return registry, simulatedBackend
}