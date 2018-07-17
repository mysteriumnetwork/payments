package registry

import (
	"fmt"
)

const registerPrefix = "Register prefix:"

type Signer interface {
	Sign(data ...[]byte) ([]byte, error)
}

type PublicKeyHolder interface {
	GetPublicKey() ([]byte, error)
}

type IdentityHolder interface {
	Signer
	PublicKeyHolder
}

type ProofOfIdentity struct {
	Data      []byte
	Signature *DecomposedSignature
}

func (proof *ProofOfIdentity) String() string {
	return fmt.Sprintf("Proof: %+v", *proof)
}

func CreateProofOfIdentity(identityHolder IdentityHolder) (*ProofOfIdentity, error) {
	pubKeyBytes, err := identityHolder.GetPublicKey()
	if err != nil {
		return nil, err
	}

	signature, err := identityHolder.Sign([]byte(registerPrefix), pubKeyBytes)
	if err != nil {
		return nil, err
	}

	decSig, err := DecomposeSignature(signature)
	if err != nil {
		return nil, err
	}

	return &ProofOfIdentity{
		pubKeyBytes,
		decSig,
	}, nil
}
