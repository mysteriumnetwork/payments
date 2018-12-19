package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/cli/helpers"
	"github.com/mysteriumnetwork/payments/mysttoken"
)

var erc20contract = flag.String("erc20.address", "", "Address of ERC20 mintable token")
var whom = flag.String("forAddress", "", "Address for which tokens to mint")
var amount = flag.Int64("amount", 1000000, "Amount of tokens to mint")
var gethUrl = flag.String("geth.url", "", "URL value of started geth to connect")
var keystoreDir = flag.String("keystore.directory", "testnet", "specify runtime dir for keystore keys")
var passphrase = flag.String("keystore.passphrase", "", "Pashprase to unlock specified key from keystore")
var etherAddress = flag.String("ether.address", "", "Ethereum acc to use for deployment")

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
	ks := helpers.GetKeystore(*keystoreDir)
	acc, err := helpers.GetUnlockedAcc(*etherAddress, *passphrase, ks)
	if err != nil {
		return
	}
	fmt.Println("Lookedup acc: ", acc.Address.String())
	transactor := helpers.CreateNewKeystoreTransactor(ks, acc)

	client, _, err := helpers.LookupBackend(*gethUrl)
	if err != nil {
		return err
	}
	balance, err := client.BalanceAt(context.Background(), acc.Address, nil)
	if err != nil {
		return
	}
	fmt.Println("Your balance is:", balance.String(), "wei")

	erc20token, err := mysttoken.NewMystTokenTransactor(common.HexToAddress(*erc20contract), client)
	if err != nil {
		return
	}
	tx, err := erc20token.Mint(transactor, common.HexToAddress(*whom), big.NewInt(*amount))
	fmt.Println("Tx id:", tx.Hash())
	return err
}
