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
	privateKey := "dbabf174189c2243a45ad6cf42ab6edda35b7488e72dfd340ecf2207cda9ece6"
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	return crypto.ToECDSA(privateKeyBytes)
}

func getPromise() Promise {
	channelID := "0xb014B75485fE0771647AF208C95254Fc25e79fD3"
	amount := uint64(10)
	fee := uint64(0)
	hashlock := "9997bf5efa5c135d1f97b522858ce7b3b88091595409179b3b5a8d703d22f233"
	signature := "d8692923462e081339ffd303ca50cbf43df74837aeabe220e3aff4912c6b333a69d2f48094bf58635a224daf5eb09aeb4b1913bb985ec8beca03cd6c957ad27e1b"

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
	expectedHash, _ := hex.DecodeString("3596062bb6d147a8ccc8d3eef230bbd104750f0f289a47badb26bab6f4539170")
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
	expectedSigner := common.HexToAddress("0x5270e411ece9fc12d5ff1022d1ddb64397ca18e3")
	assert.True(t, promise.ValidatePromise(expectedSigner))
}

func TestRecoverSigner(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0x5270e411ece9fc12d5ff1022d1ddb64397ca18e3")
	recoveredSigner, err := promise.RecoverSigner()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}
