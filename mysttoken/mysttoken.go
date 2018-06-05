package mysttoken

import (
	"github.com/MysteriumNetwork/payments/mysttoken/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//go:generate abigen --sol ../contracts/MystToken.sol --pkg generated --out generated/mysttoken.go

type MystERC20 struct {
	generated.ERC20Session
	Address common.Address
}

func DeployMystERC20(owner *bind.TransactOpts, amount int64, backend bind.ContractBackend) (*MystERC20, error) {
	address, _, mintableToken, err := generated.DeployMintableToken(owner, backend)
	if err != nil {
		return nil, err
	}

	mintableTokenSession := generated.MintableTokenSession{
		TransactOpts: *owner,
		CallOpts:     bind.CallOpts{},
		Contract:     mintableToken,
	}

	_, err = mintableTokenSession.Mint(owner.From, big.NewInt(amount))
	if err != nil {
		return nil, err
	}

	erc20, err := generated.NewERC20(address, backend)
	if err != nil {
		return nil, err
	}

	return &MystERC20{
		Address: address,
		ERC20Session: generated.ERC20Session{
			TransactOpts: *owner,
			CallOpts:     bind.CallOpts{},
			Contract:     erc20,
		},
	}, nil
}
