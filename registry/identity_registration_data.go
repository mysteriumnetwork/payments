package registry

import (
	"fmt"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mysteriumnetwork/payments/identity"
)

const registerPrefix = "Register prefix:"

type PublicKeyProvider interface {
	GetPublicKey() (ecdsa.PublicKey, error)
}

type IdentityHolder interface {
	identity.Signer
	PublicKeyProvider
}

type PublicKeyParts struct {
	Part1 []byte
	Part2 []byte
}

type RegistrationData struct {
	PublicKey PublicKeyParts
	Signature *identity.DecomposedSignature
}

func (proof *RegistrationData) String() string {
	return fmt.Sprintf("Proof: %+v", *proof)
}

func CreateRegistrationData(identityHolder IdentityHolder) (*RegistrationData, error) {
	pubKey, err := identityHolder.GetPublicKey()
	if err != nil {
		return nil, err
	}

	splitedKey := SplitPublicKey(&pubKey)

	signature, err := identityHolder.Sign([]byte(registerPrefix), splitedKey.Part1, splitedKey.Part2)
	if err != nil {
		return nil, err
	}

	decSig, err := identity.DecomposeSignature(signature)
	if err != nil {
		return nil, err
	}

	return &RegistrationData{
		splitedKey,
		decSig,
	}, nil
}

func SplitPublicKey(key *ecdsa.PublicKey) PublicKeyParts {
	pubKeyBytes := crypto.FromECDSAPub(key)
	return PublicKeyParts{
		Part1: pubKeyBytes[1:33],
		Part2: pubKeyBytes[33:65],
	}
}
