package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/MysteriumNetwork/payments/cli/helpers"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

var cmd = flag.String("cmd", "help", "Command to execute")
var contractName = flag.String("contract.name", "settlement", "Name of contract to deply (settlement or testerc20(testnet only!))")
var tokenCount = flag.Int64("mysttoken.amount", 1000000, "Initial token amount to deploy - can always be minted later")
var registrationFee = flag.Int64("payments.registrationFee", 100, "Registration fee for identity. Can be changed later")
var erc20address = flag.String("payments.erc20address", "", "ERC20 token address for payments. In hex (0x...) format")
var contractPath = flag.String("payments.contractPath", "", "Path to bin file of payments contract")
var abiPath = flag.String("payments.abiPath", "", "Path to ABI file of payments contract")

func main() {
	flag.Parse()
	err := executeCommand(*cmd)
	if err != nil {
		fmt.Printf("error occuried: %v\n", err)
		os.Exit(-1)
	}
}

func executeCommand(cmd string) error {
	fmt.Println("Executing: " + cmd)
	switch cmd {
	case "deploy":
		return deployContract()
	case "newAccount":
		return helpers.NewAccount()
	case "listAccounts":
		return helpers.ListAccounts()
	case "clientStatus":
		_, completed, err := helpers.LookupBackend()
		<-completed
		return err
	case "ethBalance":
		return getBalanceOf(*helpers.Address)
	case "help":
		flag.Usage()
		return nil
	}
	return errors.New("unknown command: " + cmd)
}
func getBalanceOf(address string) error {
	backend, syncCompleted, err := helpers.LookupBackend()
	if err != nil {
		return err
	}
	<-syncCompleted
	balance, err := backend.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return err
	}
	fmt.Println("Acc balance:", balance.Div(balance, big.NewInt(1*1000*1000*1000*1000*1000*1000)).String(), "Eth")
	return nil
}

func deployContract() (err error) {
	ks := helpers.GetKeystore()
	acc, err := helpers.GetUnlockedAcc(ks)
	if err != nil {
		return
	}
	fmt.Println("Lookedup acc: ", acc.Address.String())
	transactor := helpers.CreateNewKeystoreTransactor(ks, acc)

	client, _, err := helpers.LookupBackend()
	if err != nil {
		return err
	}
	balance, err := client.BalanceAt(context.Background(), acc.Address, nil)
	if err != nil {
		return
	}
	fmt.Println("Your balance is:", balance.String(), "wei")

	switch *contractName {
	case "payments":
		if *erc20address == "" {
			return errors.New("erc20 token address missing")
		}
		return DeployPromises(transactor, client, common.HexToAddress(*erc20address), *registrationFee)
	case "testerc20":
		return DeployTestErc20(transactor, client, *tokenCount)
	case "payments2":
		addr, _, err := DeploySmartContractFile(transactor, *contractPath, *abiPath, client, common.HexToAddress(*erc20address), big.NewInt(*registrationFee))
		if err != nil {
			return err
		}
		fmt.Println("Address of deployed contract: ", addr.String())
		return nil
	default:
		return errors.New("unknown contract name" + *contractName)
	}
	return
}
