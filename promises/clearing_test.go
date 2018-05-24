package promises

import (
	"testing"
	"math/big"
	"github.com/stretchr/testify/assert"
	"github.com/MysteriumNetwork/payments/promises/generated"
	"github.com/ethereum/go-ethereum/crypto"
	"time"
	"github.com/MysteriumNetwork/payments/registry"
	"github.com/MysteriumNetwork/payments/test_utils"
	"github.com/MysteriumNetwork/payments/mysttoken"
)

func TestPromiseClearingEmitsClearedEvent(t *testing.T) {
	backend := test_utils.NewSimulatedBackend(test_utils.Deployer.Address , 10000000000)

	mystErc20 , err := mysttoken.DeployMystERC20(test_utils.Deployer.Transactor , 1000000, backend)
	assert.NoError(t , err)

	clearing, err := DeployPromiseClearer(test_utils.Deployer.Transactor , mystErc20.Address , 1000 , backend)
	assert.NoError(t ,err)
	backend.Commit()

	_ , err = mystErc20.Approve( clearing.Address , big.NewInt(2000))
	assert.NoError(t , err)

	events:=make(chan *generated.IdentityPromisesPromiseCleared,1)
	sub , err:= clearing.BindForEvents(events)
	assert.NoError(t, err)

	payer, err := registry.NewMystIdentity()
	assert.NoError(t, err)

	receiver, err := registry.NewMystIdentity()
	assert.NoError(t, err)

	err = clearing.RegisterIdentities(*payer, *receiver)
	assert.NoError(t , err)
	backend.Commit()

	promise := Promise{
		Receiver: receiver.Address,
		SeqNo:    1,
		Amount:   100,
	}

	issuedPromise, err := SignByPayer(&promise , payer)
	assert.NoError(t, err)
	receivedPromise , err := SignByReceiver(issuedPromise, receiver)
	assert.NoError(t, err)

	err = clearing.ClearReceivedPromise(receivedPromise)
	assert.NoError(t, err)
	backend.Commit()

	select {
	case event:= <- events :
		assert.Equal(t , big.NewInt(1), event.SeqNo)
		assert.Equal(t, big.NewInt(100), event.Amount)
		assert.Equal(t, payer.Address , event.From)
		assert.Equal(t, receiver.Address , event.To)
	case err:= <- sub.Err() :
		assert.NoError(t , err)
	case <- time.After(100 * time.Millisecond):
		assert.Fail(t, "Event from contract expected")
	}

	sub.Unsubscribe()
}

const promisePrefix = "Promise prefix:"
func ethHash(msg string) ([]byte) {
	return crypto.Keccak256( []byte(promisePrefix) , []byte(msg))
}