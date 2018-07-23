package balances

import (
	"math/big"

	"github.com/MysteriumNetwork/payments/balances/generated"
	"github.com/MysteriumNetwork/payments/identity"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

//go:generate abigen --sol ../contracts/IdentityBalances.sol --pkg generated --out generated/IdentityBalances.go

type IdentityBalances struct {
	Address common.Address
	generated.IdentityBalancesSession
}

func DeployIdentityBalances(transactorOpts *bind.TransactOpts, erc20Address common.Address, backend bind.ContractBackend) (*IdentityBalances, error) {
	identityBalancesAddress, _, identityBalances, err := generated.DeployIdentityBalances(transactorOpts, backend, erc20Address)
	if err != nil {
		return nil, err
	}

	return &IdentityBalances{
		Address: identityBalancesAddress,
		IdentityBalancesSession: generated.IdentityBalancesSession{
			Contract:     identityBalances,
			TransactOpts: *transactorOpts,
			CallOpts:     bind.CallOpts{},
		},
	}, nil
}

func (balances *IdentityBalances) BindForTopUpEvents(sink chan<- *generated.IdentityBalancesToppedUp) (event.Subscription, error) {
	blockStart := uint64(0)
	return balances.Contract.WatchToppedUp(&bind.WatchOpts{Start: &blockStart}, sink, []common.Address{})
}

func (balances *IdentityBalances) BindForWithdrawEvents(sink chan<- *generated.IdentityBalancesWithdrawn) (event.Subscription, error) {
	blockStart := uint64(0)
	return balances.Contract.WatchWithdrawn(&bind.WatchOpts{Start: &blockStart}, sink, []common.Address{})
}

type WithdrawRequest struct {
	Amount    *big.Int
	Signature *identity.DecomposedSignature
}

const withDrawPrefix = "Withdraw request:"

func NewWithdrawRequest(signer identity.Signer, amount int64) (*WithdrawRequest, error) {
	bigAmount := big.NewInt(amount)
	amountBytes := abi.U256(bigAmount)
	signature, err := signer.Sign([]byte(withDrawPrefix), amountBytes)
	if err != nil {
		return nil, err
	}
	decomposed, err := identity.DecomposeSignature(signature)
	if err != nil {
		return nil, err
	}
	return &WithdrawRequest{
		//Amount:    big.NewInt(amount),
		Amount:    bigAmount,
		Signature: decomposed,
	}, nil
}

func (balances *IdentityBalances) Withdraw(request *WithdrawRequest) (err error) {
	signature := request.Signature
	_, err = balances.IdentityBalancesSession.Withdraw(request.Amount, signature.V, signature.R, signature.S)
	return
}
