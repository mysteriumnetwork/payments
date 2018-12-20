package mysttoken

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen --sol ../contracts/MystToken.sol --pkg mysttoken --out abigen_mysttoken.go

type MystERC20 struct {
	ERC20Session
	Address common.Address
}

func DeployMystERC20(owner *bind.TransactOpts, amount int64, backend bind.ContractBackend) (*MystERC20, error) {
	address, _, mintableToken, err := DeployMystToken(owner, backend)
	if err != nil {
		return nil, err
	}

	mintableTokenSession := MystTokenSession{
		TransactOpts: *owner,
		CallOpts:     bind.CallOpts{},
		Contract:     mintableToken,
	}

	_, err = mintableTokenSession.Mint(owner.From, big.NewInt(amount))
	if err != nil {
		return nil, err
	}

	erc20, err := NewERC20(address, backend)
	if err != nil {
		return nil, err
	}

	return &MystERC20{
		Address: address,
		ERC20Session: ERC20Session{
			TransactOpts: *owner,
			CallOpts:     bind.CallOpts{},
			Contract:     erc20,
		},
	}, nil
}
