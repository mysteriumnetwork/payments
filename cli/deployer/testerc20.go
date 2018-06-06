package main

import (
	"fmt"
	"github.com/MysteriumNetwork/payments/mysttoken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func DeployTestErc20(transactor *bind.TransactOpts, backend bind.ContractBackend, initialTokenAmount int64) error {
	erc20, err := mysttoken.DeployMystERC20(transactor, initialTokenAmount, backend)
	if err != nil {
		return err
	}
	fmt.Println("Test ERC20 deployed at address: ", erc20.Address.String())
	fmt.Println("Write down this address. You need it to reference smart contract")
	return nil
}
