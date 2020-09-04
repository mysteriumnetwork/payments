package crypto

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ReferralTokenRequest struct {
	Identity  common.Address
	Signature string
}

// CreateSignature signs promise using keystore
func CreateReferralTokenRequest(ks hashSigner, signer common.Address) (ReferralTokenRequest, error) {
	message := signer.Bytes()
	hash := crypto.Keccak256(message)
	signature, err := ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
	if err != nil {
		return ReferralTokenRequest{}, err
	}

	return ReferralTokenRequest{
		Signature: hex.EncodeToString(signature),
		Identity:  signer,
	}, nil
}

func (rtr ReferralTokenRequest) IsValid() error {
	b, err := hex.DecodeString(rtr.Signature)
	if err != nil {
		return err
	}

	k, err := crypto.Ecrecover(crypto.Keccak256(rtr.Identity.Bytes()), b)
	if err != nil {
		return err
	}
	pubKey, err := crypto.UnmarshalPubkey(k)
	if err != nil {
		return err
	}

	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	if !bytes.Equal(rtr.Identity.Bytes(), recoveredAddress.Bytes()) {
		return errors.New("identities do not match")
	}

	return nil
}
