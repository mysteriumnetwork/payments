package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/MysteriumNetwork/payments/cli/helpers"
	"github.com/MysteriumNetwork/payments/mysttoken/generated"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

var erc20contract = flag.String("erc20.address", "", "Address of ERC20 mintable token")
var whom = flag.String("forAddress", "", "Address for which tokens to mint")
var amount = flag.Int64("amount", 1000000, "Amount of tokens to mint")

func main() {
	flag.Parse()

	err := mintToken()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}
	os.Exit(0)
}

func mintToken() (err error) {
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

	erc20token, err := generated.NewMystTokenTransactor(common.HexToAddress(*erc20contract), client)
	if err != nil {
		return
	}
	tx, err := erc20token.Mint(transactor, common.HexToAddress(*whom), big.NewInt(*amount))
	fmt.Println("Tx id:", tx.Hash())
	return err
}
