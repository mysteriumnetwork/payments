package promises

import (
	"github.com/MysteriumNetwork/payments/promises/generated"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/MysteriumNetwork/payments/registry"
)

//go:generate abigen --sol ../contracts/IdentityPromises.sol --exc contract/registry.sol:IdentityRegistry --pkg generated --out generated/IdentityPromises.go

type PromiseClearing struct {
	Address common.Address
	generated.IdentityPromisesSession
}

func DeployPromiseClearer(owner * bind.TransactOpts , erc20Token common.Address , fee int64, backend bind.ContractBackend) (*PromiseClearing, error) {
	address , _ , clearingContract , err := generated.DeployIdentityPromises(owner , backend, erc20Token, big.NewInt(fee))
	if err != nil {
		return nil , err
	}

	return NewPromiseClearer(owner , clearingContract, address), nil
}


func NewPromiseClearer(transactOpts * bind.TransactOpts, contract * generated.IdentityPromises, address common.Address) *PromiseClearing {
	return &PromiseClearing{
		address,
		generated.IdentityPromisesSession{
			Contract: contract,
			CallOpts: bind.CallOpts{},
			TransactOpts: *transactOpts,
		},
	}
}

func (pc *PromiseClearing) RegisterIdentities(identities ...registry.MystIdentity) error {
	for _ , identity := range identities {
		proof, err := registry.CreateProofOfIdentity(&identity)
		if err != nil {
			return err
		}
		sig := proof.Signature
		var pubKeyPart1 [32]byte
		var pubKeyPart2 [32]byte
		copy(pubKeyPart1[:] , proof.Data[0:32])
		copy(pubKeyPart2[:] , proof.Data[32:64])
		_ , err = pc.RegisterIdentity( pubKeyPart1 , pubKeyPart2 , sig.V , sig.R, sig.S)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pc *PromiseClearing) ClearReceivedPromise(promise * ReceivedPromise) error {
	issuerSig , err := registry.DecomposeSignature(promise.IssuerSignature)
	if err != nil {
		return err
	}
	receiverSig , err := registry.DecomposeSignature(promise.ReceiverSignature)
	if err != nil {
		return err
	}
	var packedAddressAndSigns [32]byte
	addressAndSigns := append(promise.Receiver.Bytes() , issuerSig.V, receiverSig.V)
	copy(packedAddressAndSigns[0:22] , addressAndSigns)
	_ , err = pc.ClearPromise(
		packedAddressAndSigns,
		big.NewInt(promise.SeqNo),
		big.NewInt(promise.Amount),
		issuerSig.R,
		issuerSig.S,
		receiverSig.R,
		receiverSig.S,
	)
	return err
}

func (pc *PromiseClearing) BindForEvents(eventChan chan<- *generated.IdentityPromisesPromiseCleared) (event.Subscription , error) {
	start := uint64(0)
	return pc.Contract.WatchPromiseCleared(&bind.WatchOpts{Start:&start} , eventChan, []common.Address{}, []common.Address{})
}

func (pc * PromiseClearing) LastClearedPromise(sender common.Address, receiver common.Address) (uint64 , error) {
	seq , err := pc.ClearedPromises(sender , receiver)
	if err != nil {
		return 0 , err
	}
	return seq.Uint64() , err
}