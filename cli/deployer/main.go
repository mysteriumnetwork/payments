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
	"math/big"
)

var keyStoreDir = flag.String("keystore.directory" , "testnet" , "specify runtime dir for keystore keys")
var address = flag.String("ether.address" , "", "Ethereum acc to use for deployment")
var passphrase = flag.String("keystore.passphrase" , "", "Pashprase to unlock specified key from keystore")
var cmd = flag.String("cmd" , "help" , "Command to execute")
var gethUrl = flag.String("geth.url", "", "URL value of started geth to connect")
var contractName = flag.String("contract.name" , "settlement" , "Name of contract to deply (settlement or testerc20(testnet only!))")
var tokenCount = flag.Int64("mysttoken.amount" , 1000000 , "Initial token amount to deploy - can always be minted later")
var registrationFee = flag.Int64("payments.registrationFee" , 100 , "Registration fee for identity. Can be changed later")
var erc20address = flag.String("payments.erc20address" , "", "ERC20 token address for payments. In hex (0x...) format")

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
	case "clientStatus" :
		_ , completed , err := lookupBackend()
		<- completed
		return err
	case "ethBalance" :
		return getBalanceOf(*address)
	case "help" :
		flag.Usage()
		return nil
	}
	return errors.New("unknown command: " + cmd)
}
func getBalanceOf(address string) error {
	backend , syncCompleted , err := lookupBackend()
	if err !=  nil {
		return err
	}
	<- syncCompleted
	balance , err := backend.BalanceAt(context.Background() , common.HexToAddress(address) , nil)
	if err != nil {
		return err
	}
	fmt.Println("Acc balance:" , balance.Div(balance , big.NewInt(1 * 1000 * 1000 * 1000 * 1000 * 1000 * 1000)).String() , "Eth")
	return nil
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

	client, _ , err := lookupBackend()
	if err != nil {
		return err
	}
	balance, err := client.BalanceAt(context.Background() , acc.Address, nil)
	if err != nil {
		return
	}
	fmt.Println("Your balance is:" , balance.String() , "wei")

	switch *contractName {
	case "payments":
		if *erc20address == "" {
			return errors.New("erc20 token address missing")
		}
		return DeployPromises(transactor, client , common.HexToAddress(*erc20address), *registrationFee)
	case "testerc20":
		return DeployTestErc20(transactor , client, *tokenCount)
	default:
		return errors.New("unknown contract name" + *contractName)
	}
	return
}

func lookupBackend() (*ethclient.Client , chan bool, error) {
	ethClient , err := ethclient.Dial(*gethUrl)
	if err != nil {
		return nil,nil, err
	}

	block , err := ethClient.BlockByNumber(context.Background() , nil)
	if err != nil {
		return nil,nil, err
	}
	fmt.Println("Latest known block is: " , block.NumberU64())

	progress , err := ethClient.SyncProgress(context.Background())
	if err != nil {
		return nil ,nil, err
	}
	completed := make(chan bool)
	if progress != nil {
		fmt.Println("Client is in syncing state - any operations will be delayed until finished")
		go trackGethProgress(ethClient , progress, completed)
	} else {
		fmt.Println("Geth process fully synced")
		close(completed)
	}

	return ethClient, completed, nil
}
func trackGethProgress(client *ethclient.Client , lastProgress * ethereum.SyncProgress, completed chan<- bool ) {
	bar := pb.New64(int64(lastProgress.HighestBlock)).
	   SetTotal(int64(lastProgress.CurrentBlock)).
	   Start()
		defer bar.Finish()
		defer close(completed)
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