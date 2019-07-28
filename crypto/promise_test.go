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
	expectedHash, _ := hex.DecodeString("3c740a97b1ed8bb99ec2d865391367ba0e0b7107b7feed3b2e4993534adbc7ff")
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

func TestValidatePromise(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0x90a17343eab4d298d811c734a9572c1f494f487a")
	assert.True(t, promise.ValidatePromise(expectedSigner))
}

func TestRecoverSigner(t *testing.T) {
	promise := getPromise()
	expectedSigner := common.HexToAddress("0x90a17343eab4d298d811c734a9572c1f494f487a")
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

	channelID, amount, fee, hashlock, signature := getParams()
	r, _ := hex.DecodeString("d33f40690a6cc98e301d7d20a380d3b379372bba7cdac9801fd7752db77664ab")

	promise, err := CreatePromise(channelID, amount, fee, r, ks, account.Address)
	assert.Nil(t, err)
	assert.Equal(t, signature, promise.Signature)
	assert.Equal(t, hashlock, promise.Hashlock)
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
	privateKey := "acad3cafd077d07f9c8d643bca3fa6a435c5569ed3acae3418890a0ee0faebe3"
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	pk, _ := crypto.ToECDSA(privateKeyBytes)
	return pk
}

func getPromise() Promise {
	channelID, amount, fee, hashlock, signature := getParams()

	promise := Promise{
		ChannelID: channelID,
		Amount:    amount,
		Fee:       fee,
		Hashlock:  hashlock,
		Signature: signature,
	}

	return promise
}

// channelID, amount, fee, hashlock, signature
func getParams() (string, uint64, uint64, string, string) {
	return "0x349d0C31a29a33687479490E75Cf93ba8D662482",
		uint64(10),
		uint64(0),
		"e0931113784fb1ac168b8873a2be4f54f45521e9c3d60389911dd9cd966f5ef0",
		"c80721445cabd68ddf725355125020fda81e09f0a49db95315267979724eff3238750072a22e7247f946b64198b42f886d2674a3dbadfd327b41e1158efbe16c1b"
}
