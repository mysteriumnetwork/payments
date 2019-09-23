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
	promise := getPromise()
	expectedHash, _ := hex.DecodeString("8595953ff94b9c0d60c98448971b86dfb9bd83ec0d8a701d72363380fadf2f16")
	assert.Equal(t, expectedHash, promise.GetHash())
}

func TestCreateSignature(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey(), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	promise := getPromise()
	expectedSignature, _ := hex.DecodeString(promise.Signature)

	signature, err := promise.CreateSignature(ks, account.Address)
	assert.Nil(t, err)

	ReformatSignatureVForBC(signature)
	assert.Equal(t, expectedSignature, signature)
}

func TestPromiseValidation(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0x0d71535454f4fc153e545c3fc7cfc412ad7782c8")
	assert.True(t, promise.IsPromiseValid(expectedSigner))
}

func TestRecoverSigner(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0x0d71535454f4fc153e545c3fc7cfc412ad7782c8")
	recoveredSigner, err := promise.RecoverSigner()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

func TestCreatePromise(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey(), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	channelID, amount, fee, hashlock, signature, _, _ := getParams()

	promise, err := CreatePromise(channelID, amount, fee, hashlock, ks, account.Address)
	assert.Nil(t, err)
	assert.Equal(t, signature, promise.Signature)
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

func getPrivKey() *ecdsa.PrivateKey {
	privateKey := "d112a9bc923053f9dc2e91c49d109458eba7f5620f827ae07193cc58018e27e9"
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	pk, _ := crypto.ToECDSA(privateKeyBytes)
	return pk
}

func getPromise() Promise {
	channelID, amount, fee, hashlock, signature, _, _ := getParams()

	promise := Promise{
		ChannelID: channelID,
		Amount:    amount,
		Fee:       fee,
		Hashlock:  hashlock,
		Signature: signature,
	}

	return promise
}

// channelID, amount, fee, hashlock, promise signature, exchange message signature, provider
func getParams() (string, uint64, uint64, string, string, string, string) {
	return "0xe64eD307C0a90751923094337C2423BeF874598b",
		uint64(10),
		uint64(0),
		"6606c733283a2350e549d4148543959b21d1c1bdb2b3040621b8b444b43348b2",
		"48eeaf5d3373d946435263a33a414a3a35099fb69f69338a955ad326eb7a4d3b61764a69b285d40651ed5c9e5d4365bba0f08caacd03742c63d51d2da4372dc61c",
		"1b6412d22cd2322408fa1f55506c99a6b2901e2f1e9605685c1cc6a4dfb32f0e6c9aceac46bf1447f22e3b7a596a90d464a0fc9d56d9a6ba83e93521793493df1c",
		"0x1a9fab9aba871ed0f5bff28f9f9e52d374376611"
}
