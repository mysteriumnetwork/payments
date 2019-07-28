package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func getExchangeMessage() ExchangeMessage {
	return ExchangeMessage{
		Promise:         getPromise(),
		AgreementID:     uint64(1),
		AgreementAmount: uint64(10),
		Provider:        "0x4d55a65fa8516487d7d33aaa3ebbdb08539f547d",
		Signature:       "215aad021cc417f8b072eb574b3d22db41998fa5542b3bee5e05c64d6ae1afdc0560360c66a9cf9517367829dce5c68fd54b751a82de5c2c83e23bc8d94ae1dd1c",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("f7e072ae6f9b72790997abfcf6f3b71e6d667b4ff0696ea2a1f794fb3c93c7f1")
	assert.Equal(t, expectedHash, message.GetMessageHash())
}

func TestRecoverConsumerIdentity(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0x90a17343eab4d298d811c734a9572c1f494f487a")
	recoveredSigner, err := message.RecoverConsumerIdentity()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

func TestValidateExchangeMessage(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0x90a17343eab4d298d811c734a9572c1f494f487a")
	assert.True(t, message.ValidateExchangeMessage(expectedSigner))
}
