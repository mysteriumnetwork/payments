package crypto

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

const ExitPrefix = "Exit request:"

type ExitRequest struct {
	ChannelID   common.Address
	Beneficiary common.Address
	ValidUntil  *big.Int
	Signature   []byte
}

func (er ExitRequest) GetMessage() []byte {
	msg := []byte{}
	msg = append(msg, []byte(ExitPrefix)...)
	msg = append(msg, Pad(er.ChannelID[:], 32)...)
	msg = append(msg, Pad(er.Beneficiary[:], 32)...)
	msg = append(msg, Pad(math.U256(er.ValidUntil).Bytes(), 32)...)
	return msg
}

func (er ExitRequest) RecoverSigner() (common.Address, error) {
	sig := make([]byte, 65)
	copy(sig, er.Signature)

	err := ReformatSignatureVForRecovery(sig)
	if err != nil {
		return common.Address{}, err
	}

	return RecoverAddress(er.GetMessage(), sig)
}

func (er ExitRequest) CreateSignature(ks hashSigner, signer common.Address) ([]byte, error) {
	message := er.GetMessage()
	hash := crypto.Keccak256(message)
	return ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
}

func (er ExitRequest) Sign(ks *keystore.KeyStore, signer common.Address) error {
	signature, err := er.CreateSignature(ks, signer)
	if err != nil {
		return err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return fmt.Errorf("failed to reformat signature: %w", err)
	}

	er.Signature = signature

	return nil
}
