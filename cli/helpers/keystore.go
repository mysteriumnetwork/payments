package helpers

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var KeyStoreDir = flag.String("keystore.directory", "testnet", "specify runtime dir for keystore keys")
var Passphrase = flag.String("keystore.passphrase", "", "Pashprase to unlock specified key from keystore")
var Address = flag.String("ether.address", "", "Ethereum acc to use for deployment")

func GetKeystore() *keystore.KeyStore {
	return keystore.NewKeyStore(*KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
}

func ListAccounts() error {
	ks := GetKeystore()
	for i, acc := range ks.Accounts() {
		fmt.Printf("%d: Address: %s\n", i, acc.Address.String())
	}
	return nil
}

func NewAccount() (err error) {
	ks := GetKeystore()
	_, err = ks.NewAccount(*Passphrase)
	return
}

func GetUnlockedAcc(ks *keystore.KeyStore) (*accounts.Account, error) {
	searchAcc := accounts.Account{Address: common.HexToAddress(*Address)}
	foundAcc, err := ks.Find(searchAcc)
	if err != nil {
		return nil, err
	}
	err = ks.Unlock(foundAcc, *Passphrase)
	if err != nil {
		return nil, err
	}
	return &foundAcc, nil
}

func CreateNewKeystoreTransactor(ks *keystore.KeyStore, acc *accounts.Account) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: acc.Address,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != acc.Address {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := ks.SignHash(*acc, signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)

		},
	}
}
