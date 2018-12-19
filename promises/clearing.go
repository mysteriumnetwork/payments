package promises

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/mysteriumnetwork/payments/contracts/abigen"
	"github.com/mysteriumnetwork/payments/identity"
	"github.com/mysteriumnetwork/payments/registry"
)

type PromiseClearing struct {
	Address common.Address
	abigen.IdentityPromisesSession
}

func DeployPromiseClearer(owner *bind.TransactOpts, erc20Token common.Address, fee int64, backend bind.ContractBackend) (*PromiseClearing, error) {
	address, _, clearingContract, err := abigen.DeployIdentityPromises(owner, backend, erc20Token, big.NewInt(fee))
	if err != nil {
		return nil, err
	}

	return NewPromiseClearer(owner, clearingContract, address), nil
}

func NewPromiseClearer(transactOpts *bind.TransactOpts, contract *abigen.IdentityPromises, address common.Address) *PromiseClearing {
	return &PromiseClearing{
		address,
		abigen.IdentityPromisesSession{
			Contract:     contract,
			CallOpts:     bind.CallOpts{},
			TransactOpts: *transactOpts,
		},
	}
}

func (pc *PromiseClearing) RegisterIdentities(identities ...registry.IdentityHolder) error {
	for _, identity := range identities {
		registrationData, err := registry.CreateRegistrationData(identity)
		if err != nil {
			return err
		}
		sig := registrationData.Signature
		var pubKeyPart1 [32]byte
		var pubKeyPart2 [32]byte
		copy(pubKeyPart1[:], registrationData.PublicKey.Part1)
		copy(pubKeyPart2[:], registrationData.PublicKey.Part2)
		_, err = pc.RegisterIdentity(pubKeyPart1, pubKeyPart2, sig.V, sig.R, sig.S)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pc *PromiseClearing) ClearReceivedPromise(promise *ReceivedPromise) error {
	issuerSig, err := identity.DecomposeSignature(promise.IssuerSignature)
	if err != nil {
		return err
	}
	receiverSig, err := identity.DecomposeSignature(promise.ReceiverSignature)
	if err != nil {
		return err
	}
	var packedAddressAndSigns [32]byte
	addressAndSigns := append([]byte{issuerSig.V, receiverSig.V}, promise.Receiver.Bytes()...)
	copy(packedAddressAndSigns[10:32], addressAndSigns)

	var extraDataHash [32]byte
	copy(extraDataHash[:], promise.ConsumerHash())

	_, err = pc.ClearPromise(
		packedAddressAndSigns,
		extraDataHash,
		big.NewInt(promise.SeqNo),
		big.NewInt(promise.Amount),
		issuerSig.R,
		issuerSig.S,
		receiverSig.R,
		receiverSig.S,
	)
	return err
}

func (pc *PromiseClearing) BindForEvents(eventChan chan<- *abigen.IdentityPromisesPromiseCleared) (event.Subscription, error) {
	start := uint64(0)
	return pc.Contract.WatchPromiseCleared(&bind.WatchOpts{Start: &start}, eventChan, []common.Address{}, []common.Address{})
}

func (pc *PromiseClearing) LastClearedPromise(sender common.Address, receiver common.Address) (uint64, error) {
	seq, err := pc.ClearedPromises(sender, receiver)
	if err != nil {
		return 0, err
	}
	return seq.Uint64(), err
}
