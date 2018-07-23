package test_utils

import (
	"crypto/ecdsa"
)

type privateKeyProvider struct {
	privateKey *ecdsa.PrivateKey
}

func (pkh *privateKeyProvider) GetPublicKey() (ecdsa.PublicKey, error) {
	return pkh.privateKey.PublicKey, nil //drop first byte as it's encoded curve type - we are not interested in as so does ethereum EVM
}
