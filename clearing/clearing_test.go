package clearing

import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"github.com/stretchr/testify/assert"
	"github.com/MysteriumNetwork/payments/clearing/generated"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"time"
	"crypto/ecdsa"
)

var (
	deployerPrivateKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	deployerAddress       =  crypto.PubkeyToAddress(deployerPrivateKey.PublicKey)
	deployerTransactor = bind.NewKeyedTransactor(deployerPrivateKey)
)


func TestClearingContractDeploys(t *testing.T) {
	_ , _ = deployClearing(t)
}

func TestPromiseClearingEmitsClearedEvent(t *testing.T) {
	clearing, backend := deployClearing(t)

	events:=make(chan *generated.ClearingContractPromiseCleared,1000)
	sub , err:= clearing.BindForEvents(events)
	assert.NoError(t, err)

	payer, err := newMystIdentity()
	assert.NoError(t, err)

	receiver, err := newMystIdentity()
	assert.NoError(t, err)

	err = clearing.RegisterIdentities(payer.Address , receiver.Address)
	assert.NoError(t , err)
	backend.Commit()

	receiverSig , err := crypto.Sign(ethHash("abc"), receiver.PrivateKey )
	assert.NoError(t , err)
	payerSig, err := crypto.Sign(ethHash("abc"), payer.PrivateKey )
	assert.NoError(t , err)

	promise := Promise{
		seqNo: 1,
		amount: 100,
		receiverSign: receiverSig,
		payerSign: payerSig,
	}
	err = clearing.ClearMyPromise(promise)
	assert.NoError(t, err)
	backend.Commit()

	select {
	case event:= <- events :
		assert.Equal(t , big.NewInt(1), event.SeqNo)
		assert.Equal(t, big.NewInt(100), event.Amount)
	assert.Equal(t, payer.Address, event.From)
	assert.Equal(t, receiver.Address, event.To)
	case err:= <- sub.Err() :
		assert.NoError(t , err)
	case <- time.After(100 * time.Millisecond):
		assert.Fail(t, "Event from contract expected")
	}

	sub.Unsubscribe()
}

func deployClearing(t * testing.T) (*PromiseClearer , * backends.SimulatedBackend) {
	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		deployerAddress: {Balance: big.NewInt(1000000000)},
	})
	_ , _, clearing, err := generated.DeployClearingContract(deployerTransactor, simulatedBackend)
	simulatedBackend.Commit()
	if err != nil {
		assert.FailNow(t, "Unexpected error: %s", err)
	}
	return NewPromiseClearer( deployerTransactor ,clearing), simulatedBackend
}

type mystIdentity struct {
	PrivateKey  *ecdsa.PrivateKey
	PublicKey * ecdsa.PublicKey
	Address common.Address
}

func newMystIdentity() (*mystIdentity , error) {
	key , err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &mystIdentity{
		key,
		&key.PublicKey,
		crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}

const promisePrefix = "Promise prefix:"
func ethHash(msg string) ([]byte) {
	return crypto.Keccak256( []byte(promisePrefix) , []byte(msg))
}