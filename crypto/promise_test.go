package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func getPrivKey() (*ecdsa.PrivateKey, error) {
	privateKey := "4ce370439d91a8ebdad9d5210ce2e44e423ca4c29ca558024d9535874b4c9b0c"
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	return crypto.ToECDSA(privateKeyBytes)
}

func getPromise() Promise {
	channelID := "0x9ED0b24549200F4a887126C44bC0a643c26F8253"
	amount := uint64(10)
	fee := uint64(0)
	hashlock := "46b6ad53754e9f0cb132c24933fb82c62c43285991778582350ff598af9ad515"
	signature := "61e563bca6e18e669b582599291bcc2352a86b3a76df525fb9b0bfaa883d3ee1737ad93feebcfa711f79fdf775244f5a94d151c91223d3a2ca16fff069292d451c"

	promise := Promise{
		ChannelID: channelID,
		Amount:    amount,
		Fee:       fee,
		Hashlock:  hashlock,
		Signature: signature,
	}

	return promise
}

func TestGetHash(t *testing.T) {
	promise := getPromise()
	expectedHash, _ := hex.DecodeString("fb19a309c101782b340eaeb5293b2517d30ba077c257de7c333a19658b4b91bd")
	assert.Equal(t, expectedHash, promise.GetHash())
}

func TestCreateSignature(t *testing.T) {
	promise := getPromise()
	expectedSignature, _ := hex.DecodeString(promise.Signature)

	pk, err := getPrivKey()
	assert.Nil(t, err)

	signature, err := promise.CreateSignature2(pk)
	assert.Nil(t, err)

	ReformatSignatureVForBC(signature)
	assert.Equal(t, expectedSignature, signature)
}

func TestValidatePromise(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0xbd908c8eda0ebc327c3a032d98ce541bb80a7d79")
	assert.True(t, promise.ValidatePromise(expectedSigner))
}

func TestRecoverSigner(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0xbd908c8eda0ebc327c3a032d98ce541bb80a7d79")
	recoveredSigner, err := promise.RecoverSigner()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}
