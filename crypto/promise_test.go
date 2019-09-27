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
	expectedHash, _ := hex.DecodeString("8ca6e18427321051653d12a19cbf9e08a13e20b46e840c2a608603c1a67e3750")
	assert.Equal(t, expectedHash, promise.GetHash())

	promise = getPromise("consumer")
	expectedHash, _ = hex.DecodeString("3aa5d7e7d36677d92c9392df57694477e70875f52ee3a735a8ca9c9954df4e22")
	assert.Equal(t, expectedHash, promise.GetHash())
}

func TestGetSignatureHex(t *testing.T) {
	promise := getPromise("provider")
	signature := promise.GetSignatureHexString()

	expectedSignature := "0xd893cb1a6e737ab99518811d18c9310c5478100107c35e60e12f673c5e29ea447da5c26a5f5f3dd6560fa77b4b52df66287eb7b90c9d57cdf91fe5270458701f1c"
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
	messageSig := "a7a201c9ec67d5b627cda20196f80a86e9e03c9dd9e8224a73470605ef40494847dbf5f6d2701c58d9093294fc5cfdbd98e85331c191d49cd1da29d96b0c10f81c"

	if userType == "consumer" {
		channelID, _ = hex.DecodeString("2E9A53c8eECBd8093a079Eb23676eCdeee04D889")
		promiseSig, _ = hex.DecodeString("94ff709867525f27b6a5bab7b6171cf994f21b7b2ed310c7385f4a49cc88ac9001f19ada1b30a7675c521b425b3c56ec3ea351b6571da2850783876aeffa05b41b")
	}

	if userType == "provider" {
		channelID, _ = hex.DecodeString("b32af43608878c7169d6973da7f483c934cd455891b13b1a330ead47c57e8213")
		promiseSig, _ = hex.DecodeString("d893cb1a6e737ab99518811d18c9310c5478100107c35e60e12f673c5e29ea447da5c26a5f5f3dd6560fa77b4b52df66287eb7b90c9d57cdf91fe5270458701f1c")
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
