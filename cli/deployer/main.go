package main

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"flag"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"time"
	"github.com/cheggaaa/pb"
	"github.com/ethereum/go-ethereum"
)

var keyStoreDir = flag.String("keystore.directory" , "testnet" , "specify runtime dir for keystore keys")
var address = flag.String("ether.address" , "", "Ethereum acc to use for deployment")
var passphrase = flag.String("keystore.passphrase" , "", "Pashprase to unlock specified key from keystore")
var cmd = flag.String("cmd" , "deploy" , "Command to execute")
var gethUrl = flag.String("geth.url", "", "URL value of started geth to connect")
var contractName = flag.String("contract.name" , "settlement" , "Name of contract to deply (settlement or testerc20(testnet only!))")
var tokenCount = flag.Int64("mysttoken.amount" , 1000000 , "Initial token amount to deploy - can always be minted later")

func main() {
	flag.Parse()
	err := executeCommand(*cmd)
	if err != nil {
		fmt.Printf("error occuried: %v\n" , err)
		os.Exit(-1)
	}
}

func executeCommand(cmd string) error {
	fmt.Println("Executing: " + cmd)
	switch cmd {
	case "deploy" :
		return deployContract()
	case "newAccount" :
		return newAccount()
	case "listAccounts" :
		return listAccounts()
	}
	return errors.New("unknown command: " + cmd)
}
func listAccounts() error {
	ks := getKeystore()
	for i, acc := range ks.Accounts() {
		fmt.Printf("%d: Address: %s\n", i , acc.Address.String())
	}
	return nil
}

func newAccount() (err error) {
	ks := getKeystore()
	_ , err = ks.NewAccount(*passphrase)
	return
}

func deployContract() (err error) {
	ks := getKeystore()
	acc , err := getUnlockedAcc(ks)
	if err != nil {
		return
	}
	fmt.Println("Lookedup acc: " , acc.Address.String())
	transactor := createNewKeystoreTransactor(ks , acc)

	client, err := lookupBackend()
	if err != nil {
		return err
	}
	balance, err := client.BalanceAt(context.Background() , acc.Address, nil)
	if err != nil {
		return
	}
	fmt.Println("Your balance is:" , balance.String() , "wei")

	switch *contractName {
	case "settlement":
		return errors.New("not yet dudes")
	case "testerc20":
		return DeployTestErc20(transactor , client, *tokenCount)
	default:
		return errors.New("unknown contract name" + *contractName)
	}
	return
}

func lookupBackend() (*ethclient.Client , error) {
	ethClient , err := ethclient.Dial(*gethUrl)
	if err != nil {
		return nil, err
	}

	block , err := ethClient.BlockByNumber(context.Background() , nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Latest known block is: " , block.NumberU64())

	progress , err := ethClient.SyncProgress(context.Background())
	if err != nil {
		return nil , err
	}
	if progress != nil {
		fmt.Println("Client is in syncing state - any operations will be delayed until finished")
		go trackGethProgress(ethClient , progress)
	} else {
		fmt.Println("Geth process fully synced")
	}

	return ethClient, nil
}
func trackGethProgress(client *ethclient.Client , lastProgress * ethereum.SyncProgress) {
	bar := pb.New64(int64(lastProgress.HighestBlock)).
	   SetTotal(int64(lastProgress.CurrentBlock)).
	   Start()
		defer bar.Finish()
	for {
		progress, err := client.SyncProgress(context.Background())
		if err != nil {
			fmt.Println("Error querying client progress: " + err.Error())
			return
		}
		if progress == nil {
			bar.Finish()
			fmt.Println("Client in fully synced state. Proceeding...")
			return
		}
		bar.SetCurrent(int64(progress.CurrentBlock))
		bar.SetTotal(int64(progress.HighestBlock))
		time.Sleep(10 * time.Second)
	}
}

func getUnlockedAcc(ks* keystore.KeyStore) (*accounts.Account , error) {
	searchAcc := accounts.Account{ Address:common.HexToAddress(*address)}
	foundAcc , err := ks.Find(searchAcc)
	if err != nil {
		return nil, err
	}
	err = ks.Unlock(foundAcc , *passphrase)
	if err != nil {
		return nil, err
	}
	return &foundAcc , nil
}

func getKeystore() * keystore.KeyStore {
	return keystore.NewKeyStore(*keyStoreDir , keystore.StandardScryptN , keystore.StandardScryptP)
}

func createNewKeystoreTransactor(ks * keystore.KeyStore , acc * accounts.Account) (* bind.TransactOpts) {
	return &bind.TransactOpts{
		From:acc.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != acc.Address {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := ks.SignHash(*acc , signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)

		},
	}
}