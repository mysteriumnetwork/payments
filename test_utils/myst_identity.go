package test_utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type MystIdentity struct {
	*privateKeySigner
	*privateKeyProvider
	PublicKey *ecdsa.PublicKey
	Address   common.Address
}

func NewMystIdentity() (*MystIdentity, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &MystIdentity{
		&privateKeySigner{
			key,
		},
		&privateKeyProvider{
			key,
		},
		&key.PublicKey,
		crypto.PubkeyToAddress(key.PublicKey),
	}, nil
}
