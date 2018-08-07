package helpers

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetKeystore(keystoreDir string) *keystore.KeyStore {
	return keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
}

func ListAccounts(ks *keystore.KeyStore) error {
	for i, acc := range ks.Accounts() {
		fmt.Printf("%d: Address: %s\n", i, acc.Address.String())
	}
	return nil
}

func NewAccount(passphrase string, ks *keystore.KeyStore) (err error) {
	_, err = ks.NewAccount(passphrase)
	return
}

func GetUnlockedAcc(address, passphrase string, ks *keystore.KeyStore) (*accounts.Account, error) {
	searchAcc := accounts.Account{Address: common.HexToAddress(address)}
	foundAcc, err := ks.Find(searchAcc)
	if err != nil {
		return nil, err
	}
	err = ks.Unlock(foundAcc, passphrase)
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
