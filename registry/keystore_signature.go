package registry

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type keystoreIdentity struct {
	keystore *keystore.KeyStore
	account  accounts.Account
}

func (ki *keystoreIdentity) Sign(data ...[]byte) ([]byte, error) {
	hash := crypto.Keccak256(data...)
	return ki.keystore.SignHash(ki.account, hash)
}

func (ki *keystoreIdentity) GetPublicKey() ([]byte, error) {
	hash := crypto.Keccak256([]byte("public key recover"))
	sig, err := ki.keystore.SignHash(ki.account, hash)
	if err != nil {
		return nil, err
	}
	pubKeyBytes, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return nil, err
	}
	return pubKeyBytes[1:], nil
}

func FromKeystoreIdentity(keystore *keystore.KeyStore, identity common.Address) IdentityHolder {
	return &keystoreIdentity{
		keystore: keystore,
		account:  accounts.Account{Address: identity},
	}
}
