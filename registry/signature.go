package registry

import (
	"errors"
	"fmt"
)

type DecomposedSignature struct {
	R [32]byte
	S [32]byte
	V uint8
}

func DecomposeSignature(signature []byte) (*DecomposedSignature, error) {
	if len(signature) != 65 {
		return nil , errors.New("65 bytes length signature expected")
	}
	sign := &DecomposedSignature{}
	copy(sign.R[:],signature[0:32])
	copy(sign.S[:],signature[32:64])
	sign.V=signature[64] + 27
	return sign, nil
}

func (sig * DecomposedSignature)String() string {
	return fmt.Sprintf("Signature: %+v", *sig)
}