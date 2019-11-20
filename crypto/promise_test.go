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
	expectedHash, _ := hex.DecodeString("44ebc179879e8649a4734cd12181da1ab10fa6f1374d9fb021c1b541f9f4eff6")
	assert.Equal(t, expectedHash, promise.GetHash())

	promise = getPromise("consumer")
	expectedHash, _ = hex.DecodeString("177adc47f6a0962fc7994d863c0a98a5125f28f336283e0ff58cb5e9ccd18465")
	assert.Equal(t, expectedHash, promise.GetHash())
}

func TestGetSignatureHex(t *testing.T) {
	promise := getPromise("provider")
	signature := promise.GetSignatureHexString()

	expectedSignature := "0x4f9bcffd245718981fdf3997e59e428f0677f3f213af75cc6512f5fe413d053354142446ab0a76fe172328831ee0860e8f544ff478b9827c0bf62e73a0f97d421c"
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

func TestNewPromise(t *testing.T) {
	p := getParams("provider")
	channelID := hex.EncodeToString(p.ChannelID)
	amount := p.Amount
	fee := p.Fee
	preimage := hex.EncodeToString(p.R)
	signature := hex.EncodeToString(p.PromiseSignature)

	promise, err := NewPromise(channelID, amount, fee, preimage, signature)
	assert.NoError(t, err)
	assert.Equal(t, p.Amount, promise.Amount)
	assert.Equal(t, p.ChannelID, promise.ChannelID)
	assert.Equal(t, p.Hashlock, promise.Hashlock)
	assert.Equal(t, p.PromiseSignature, promise.Signature)
}

func TestSign(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("provider"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	p := getParams("provider")
	promise, err := NewPromise(
		hex.EncodeToString(p.ChannelID),
		p.Amount,
		p.Fee,
		hex.EncodeToString(p.R),
		"",
	)
	assert.NoError(t, err)

	promise.Sign(ks, account.Address)
	assert.Equal(t, p.PromiseSignature, promise.Signature)
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
	R                        []byte
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
	preimage, _ := hex.DecodeString("5b6b3f31a3acd0e317173d25c8b60503547b741a0e81d6068bb88486967839fa")
	hashlock, _ := hex.DecodeString("4e8444e4bd5721ba00ceb2c6c180c21b2ae43e590172f1b39e51f46312243633")
	messageSig := "56493421bd2772cca2ba970da27396e103a08027f1ce49de974f789e322b0d7a3f52b9dd745a34bfa2f330ba2d3c442867ebb3753d1f206811ab572ab7d482dc1b"

	if userType == "consumer" {
		channelID, _ = hex.DecodeString("000000000000000000000000d2c94475763fa7e81076ab0bde4dc4b902191498")
		promiseSig, _ = hex.DecodeString("5b2252c91e1de26b084759cc7b8dd0ae4e2feed97013dd4b8777ed098b295fb81a5357f695717e546eecd07b82ac4e2a58bf6395f7477067cd411cb8c133b5451c")
	}

	if userType == "provider" {
		channelID, _ = hex.DecodeString("40684b7886aa813a04bce9efd1c6d45379178e9248bf1cc926a316e4e4924ae8")
		promiseSig, _ = hex.DecodeString("4f9bcffd245718981fdf3997e59e428f0677f3f213af75cc6512f5fe413d053354142446ab0a76fe172328831ee0860e8f544ff478b9827c0bf62e73a0f97d421c")
	}

	return Params{
		ChannelID:                channelID,
		Amount:                   amount,
		Fee:                      fee,
		R:                        preimage,
		Hashlock:                 hashlock,
		PromiseSignature:         promiseSig,
		ExchangeMessageSignature: messageSig,
		Provider:                 provider,
	}
}
