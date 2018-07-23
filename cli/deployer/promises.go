package main

import (
	"fmt"

	"github.com/MysteriumNetwork/payments/mysttoken/generated"
	"github.com/MysteriumNetwork/payments/promises"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func DeployPromises(transactor *bind.TransactOpts, client bind.ContractBackend, erc20tokenAddress common.Address, registrationFee int64) error {
	erc20, err := generated.NewERC20(erc20tokenAddress, client)
	if err != nil {
		return err
	}

	deployerBalance, err := erc20.BalanceOf(&bind.CallOpts{}, transactor.From)
	if err != nil {
		return err
	}
	fmt.Println("Deployer balance of erc20 is: " + deployerBalance.String())

	pc, err := promises.DeployPromiseClearer(transactor, erc20tokenAddress, registrationFee, client)
	if err != nil {
		return err
	}
	fmt.Println("Deployed contract address: " + pc.Address.String())
	return nil
}
