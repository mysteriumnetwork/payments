package test_utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthAccount struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	Address    common.Address
	Transactor *bind.TransactOpts
}

func NewAccountFromKey(privateKey *ecdsa.PrivateKey) *EthAccount {

	return &EthAccount{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PublicKey,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		Transactor: bind.NewKeyedTransactor(privateKey),
	}
}

func NewRandomAccount() (*EthAccount, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return NewAccountFromKey(privateKey), nil
}

var (
	deployerPrivateKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	Deployer              = NewAccountFromKey(deployerPrivateKey)
)
