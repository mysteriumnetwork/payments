package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {
	promise := getPromise("provider")
	expectedHash, _ := hex.DecodeString("ac65e143143128d390418aa4ac6406266a3768e59430c5f649c9f9dda72a36a0")
	assert.Equal(t, expectedHash, promise.GetHash())

	promise = getPromise("consumer")
	expectedHash, _ = hex.DecodeString("9e5ed67e06c3a0ce87b96b57a3a6e7d4b0bde971a066e51d4d5d196a07713d07")
	assert.Equal(t, expectedHash, promise.GetHash())
}

func TestGetSignatureHex(t *testing.T) {
	promise := getPromise("provider")
	signature := promise.GetSignatureHexString()

	expectedSignature := "0xeb56461cba34ee3447c938b4c267b9a78fd3b5f28e5ff5d3f473e7553ddee6595d2b09e65ba29c5eb77b10e1c888a304a0e14f3ca4fe9970af0602d0bc3d48d51c"
	assert.Equal(t, expectedSignature, signature)
}

func TestCreateConsumerSignature(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("consumer"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	promise := getPromise("consumer")

	signature, err := promise.CreateSignature(ks, account.Address)
	assert.Nil(t, err)

	ReformatSignatureVForBC(signature)
	assert.Equal(t, promise.Signature, signature)
}

func TestCreateProviderSignature(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("provider"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	promise := getPromise("provider")

	signature, err := promise.CreateSignature(ks, account.Address)
	assert.Nil(t, err)

	ReformatSignatureVForBC(signature)
	assert.Equal(t, promise.Signature, signature)
}

func TestConsumerPromiseValidation(t *testing.T) {
	promise := getPromise("consumer")
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	assert.True(t, promise.IsPromiseValid(expectedSigner))
}

func TestProviderPromiseValidation(t *testing.T) {
	promise := getPromise("provider")
	expectedSigner := common.HexToAddress("0x354bd098b4ef8c9e70b7f21be2d455df559705d7")
	assert.True(t, promise.IsPromiseValid(expectedSigner))

	wrongSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	assert.False(t, promise.IsPromiseValid(wrongSigner))
}

func TestRecoverSigner(t *testing.T) {
	promise := getPromise("consumer")
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	recoveredSigner, err := promise.RecoverSigner()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

// This test ensures that promise validation functions will not mutate original signature
func TestValidationAndRecoveryImmutability(t *testing.T) {
	promise := getPromise("consumer")

	originalSignature := make([]byte, 65)
	copy(originalSignature, promise.Signature)

	// Ensure that IsPromiseValid will not mutate original signature
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	promise.IsPromiseValid(expectedSigner)
	assert.Equal(t, originalSignature, promise.Signature)

	// Ensure that RecoverSigner will not mutate original signature
	promise.RecoverSigner()
	assert.Equal(t, originalSignature, promise.Signature)
}

func TestCreatePromise(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("consumer"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	p := getParams("consumer")
	channelID := hex.EncodeToString(p.ChannelID)
	amount := p.Amount
	fee := p.Fee
	hashlock := hex.EncodeToString(p.Hashlock)

	promise, err := CreatePromise(channelID, amount, fee, hashlock, ks, account.Address)
	assert.NoError(t, err)
	assert.Equal(t, p.PromiseSignature, promise.Signature)

	// ChannelID can also be provided with prefix
	promise, err = CreatePromise("0x"+channelID, amount, fee, hashlock, ks, account.Address)
	assert.NoError(t, err)
	assert.Equal(t, p.PromiseSignature, promise.Signature)

	// Should fail when provided not correct channel ID
	_, err = CreatePromise("NotHex", amount, fee, hashlock, ks, account.Address)
	assert.Error(t, err)
	assert.Equal(t, "channelID and hashlock have to be proper hex strings", err.Error())

}

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func tmpKeyStore(t *testing.T, encrypted bool) (string, *keystore.KeyStore) {
	d, err := ioutil.TempDir("", "eth-keystore-test")
	if err != nil {
		t.Fatal(err)
	}
	newKs := keystore.NewPlaintextKeyStore
	if encrypted {
		newKs = func(kd string) *keystore.KeyStore {
			return keystore.NewKeyStore(kd, veryLightScryptN, veryLightScryptP)
		}
	}
	return d, newKs(d)
}

func getPrivKey(userType string) *ecdsa.PrivateKey {
	var privKey []byte

	if userType == "consumer" {
		privKey, _ = hex.DecodeString("1fd26004a85a0ae04e16b444893ba7408dc7215b3962dda782d1afc6d7441d1e")
	}

	if userType == "provider" {
		privKey, _ = hex.DecodeString("45bb96530f3d1972fdcd2005c1987a371d0b6d378b77561c6beeaca27498f46b")
	}

	pk, _ := crypto.ToECDSA(privKey)
	return pk
}

func getPromise(userType string) Promise {
	p := getParams(userType)

	promise := Promise{
		ChannelID: p.ChannelID,
		Amount:    p.Amount,
		Fee:       p.Fee,
		Hashlock:  p.Hashlock,
		Signature: p.PromiseSignature,
	}

	return promise
}

type Params struct {
	ChannelID                []byte
	Amount                   uint64
	Fee                      uint64
	Hashlock                 []byte
	PromiseSignature         []byte
	ExchangeMessageSignature string
	Provider                 common.Address
}

func getParams(userType string) Params {
	var channelID []byte
	var promiseSig []byte

	amount := uint64(1401)
	fee := uint64(0)
	provider := common.HexToAddress("0xf10021ba3b10d023e671668d20daeff821561d09")
	hashlock, _ := hex.DecodeString("4e8444e4bd5721ba00ceb2c6c180c21b2ae43e590172f1b39e51f46312243633")
	messageSig := "d44920d4e0bcb96e836f0731f168e862a9efcd4e72dd093141a3c95205ba6cc115e23fcbfb8625e219be4255e0f9597ef55699d13a207a21881378329b38d0b31c"

	if userType == "consumer" {
		channelID, _ = hex.DecodeString("0000000000000000000000009ec6fd7aa414d26f65a22e333c1bfd02b5208bde")
		promiseSig, _ = hex.DecodeString("848fae23369a9b97430ca5f908b11b54931b2e51e4ba541134f800a8f396eafb1d05594d262912071584af48e755cef038089a78d8520bd658d85313dd523a091c")
	}

	if userType == "provider" {
		channelID, _ = hex.DecodeString("b583afafc3d2c5b9d5cac0f01212bbd2054c385649ca13f034032563be4381ca")
		promiseSig, _ = hex.DecodeString("eb56461cba34ee3447c938b4c267b9a78fd3b5f28e5ff5d3f473e7553ddee6595d2b09e65ba29c5eb77b10e1c888a304a0e14f3ca4fe9970af0602d0bc3d48d51c")
	}

	return Params{
		ChannelID:                channelID,
		Amount:                   amount,
		Fee:                      fee,
		Hashlock:                 hashlock,
		PromiseSignature:         promiseSig,
		ExchangeMessageSignature: messageSig,
		Provider:                 provider,
	}
}
