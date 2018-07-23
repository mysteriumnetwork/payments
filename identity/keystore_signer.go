package identity

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type KeystoreSigner struct {
	keystore *keystore.KeyStore
	account  accounts.Account
}

func NewKeystoreSigner(ks *keystore.KeyStore, identity common.Address) *KeystoreSigner {
	return &KeystoreSigner{
		keystore: ks,
		account: accounts.Account{
			Address: identity,
		},
	}
}

func (ki *KeystoreSigner) Sign(data ...[]byte) ([]byte, error) {
	hash := crypto.Keccak256(data...)
	return ki.keystore.SignHash(ki.account, hash)
}
