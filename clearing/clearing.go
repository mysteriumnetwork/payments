package clearing

import (
	"github.com/MysteriumNetwork/payments/clearing/generated"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate ../scripts/abigen.sh --sol ../contracts/clearingContract.sol --exc contract/registry.sol:IdentityRegistry --pkg generated --out generated/clearing.go

type PromiseClearer struct {
	generated.ClearingContractSession
}

type Promise struct {
	seqNo int
	amount int
	receiverSign []byte
	payerSign []byte
}

func NewPromiseClearer(transactOpts * bind.TransactOpts, contract * generated.ClearingContract) *PromiseClearer {
	return &PromiseClearer{
		generated.ClearingContractSession{
			Contract: contract,
			CallOpts: bind.CallOpts{},
			TransactOpts: *transactOpts,
		},
	}
}

func (pc * PromiseClearer) ClearMyPromise(promise Promise) error {
	seqNo :=big.NewInt(int64(promise.seqNo))
	amount :=big.NewInt(int64(promise.amount))
	_ , err := pc.ClearPromise( seqNo, amount, promise.receiverSign, promise.payerSign)
	return err
}

func (pc * PromiseClearer) BindForEvents(eventChan chan<- *generated.ClearingContractPromiseCleared) (event.Subscription , error) {
	start := uint64(0)
	return pc.Contract.WatchPromiseCleared(&bind.WatchOpts{Start:&start} , eventChan, []common.Address{}, []common.Address{})
}