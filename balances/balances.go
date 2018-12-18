package balances

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/mysteriumnetwork/payments/identity"
)

//go:generate abigen --sol ../contracts/IdentityBalances.sol --pkg balances --out abigen_balances.go

type IdentityBalance struct {
	Address common.Address
	IdentityBalancesSession
}

func DeployIdentityBalance(transactorOpts *bind.TransactOpts, erc20Address common.Address, backend bind.ContractBackend) (*IdentityBalance, error) {
	identityBalancesAddress, _, identityBalances, err := DeployIdentityBalances(transactorOpts, backend, erc20Address)
	if err != nil {
		return nil, err
	}

	return &IdentityBalance{
		Address: identityBalancesAddress,
		IdentityBalancesSession: IdentityBalancesSession{
			Contract:     identityBalances,
			TransactOpts: *transactorOpts,
			CallOpts:     bind.CallOpts{},
		},
	}, nil
}

func (balances *IdentityBalance) BindForTopUpEvents(sink chan<- *IdentityBalancesToppedUp) (event.Subscription, error) {
	blockStart := uint64(0)
	return balances.Contract.WatchToppedUp(&bind.WatchOpts{Start: &blockStart}, sink, []common.Address{})
}

func (balances *IdentityBalance) BindForWithdrawEvents(sink chan<- *IdentityBalancesWithdrawn) (event.Subscription, error) {
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

func (balances *IdentityBalance) Withdraw(request *WithdrawRequest) (err error) {
	signature := request.Signature
	_, err = balances.IdentityBalancesSession.Withdraw(request.Amount, signature.V, signature.R, signature.S)
	return
}
