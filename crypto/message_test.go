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
		Provider:        "0x4b7c32c57576894807ff388c903a9592dcf4239b",
		Signature:       "c5da6209c8daa1314737c53dc68bc4505a2fc500e359f02a85d51fc0d46b826651009df0395d5d1101c39d3532ed8571629a58244a43083f9d420d122452f56e1c",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("3a6015813e24267493cc76944e4c55f73cba118165979d0273230601bdb957de")
	assert.Equal(t, expectedHash, message.GetMessageHash())
}

func TestRecoverConsumerIdentity(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0x5270e411ece9fc12d5ff1022d1ddb64397ca18e3")
	recoveredSigner, err := message.RecoverConsumerIdentity()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

func TestValidateExchangeMessage(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0x5270e411ece9fc12d5ff1022d1ddb64397ca18e3")
	assert.True(t, message.ValidateExchangeMessage(expectedSigner))
}
