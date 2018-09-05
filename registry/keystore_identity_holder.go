package registry

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/identity"
)

type keystoreIdentityHolder struct {
	identity.Signer
	keystore *keystore.KeyStore
	account  accounts.Account
}

func (ki *keystoreIdentityHolder) GetPublicKey() (ecdsa.PublicKey, error) {
	hash := crypto.Keccak256([]byte("public key recover"))
	sig, err := ki.keystore.SignHash(ki.account, hash)
	if err != nil {
		return ecdsa.PublicKey{}, err
	}
	pubKeyBytes, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return ecdsa.PublicKey{}, err
	}
	key, err := crypto.UnmarshalPubkey(pubKeyBytes)
	return *key, err
}

func FromKeystore(keystore *keystore.KeyStore, address common.Address) IdentityHolder {
	return &keystoreIdentityHolder{
		Signer:   identity.NewKeystoreSigner(keystore, address),
		keystore: keystore,
		account:  accounts.Account{Address: address},
	}
}
