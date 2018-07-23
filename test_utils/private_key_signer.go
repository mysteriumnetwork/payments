package test_utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type privateKeySigner struct {
	privateKey *ecdsa.PrivateKey
}

func NewPrivateKeySigner(key *ecdsa.PrivateKey) *privateKeySigner {
	return &privateKeySigner{
		privateKey: key,
	}
}

func (pkh *privateKeySigner) Sign(data ...[]byte) ([]byte, error) {
	return crypto.Sign(crypto.Keccak256(data...), pkh.privateKey)
}
