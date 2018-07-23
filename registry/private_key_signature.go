package registry

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type privateKeyHolder struct {
	privateKey *ecdsa.PrivateKey
}

func (pkh *privateKeyHolder) Sign(data ...[]byte) ([]byte, error) {
	return crypto.Sign(crypto.Keccak256(data...), pkh.privateKey)
}

func (pkh *privateKeyHolder) GetPublicKey() ([]byte, error) {
	pubKeyBytes := crypto.FromECDSAPub(&pkh.privateKey.PublicKey)
	return pubKeyBytes[1:], nil //drop first byte as it's encoded curve type - we are not interested in as so does ethereum EVM
}

func IdentityFromPrivateKey(privateKey *ecdsa.PrivateKey) IdentityHolder {
	return &privateKeyHolder{
		privateKey: privateKey,
	}
}
