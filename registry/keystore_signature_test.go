package registry

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestKeystoreSignerProvidesValidPublicKey(t *testing.T) {
	assert.NoError(t, os.RemoveAll("testoutput"))

	sampleKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	keyHolder := IdentityFromPrivateKey(sampleKey)

	ks := keystore.NewKeyStore("testoutput", keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.ImportECDSA(sampleKey, "")
	assert.NoError(t, err)

	assert.NoError(t, ks.Unlock(acc, ""))
	identityHolder := FromKeystoreIdentity(ks, acc.Address)
	extractedPubKey, err := identityHolder.GetPublicKey()
	assert.NoError(t, err)

	expectedPubKey, err := keyHolder.GetPublicKey()
	assert.NoError(t, err)

	assert.Equal(t, expectedPubKey, extractedPubKey)
}

func TestKeyStoreSignerProvidesValidSignature(t *testing.T) {

	assert.NoError(t, os.RemoveAll("testoutput"))

	sampleKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	keyHolder := IdentityFromPrivateKey(sampleKey)

	ks := keystore.NewKeyStore("testoutput", keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.ImportECDSA(sampleKey, "")
	assert.NoError(t, err)

	assert.NoError(t, ks.Unlock(acc, ""))
	identityHolder := FromKeystoreIdentity(ks, acc.Address)

	signData := []byte("Testing signature")

	extractedSignature, err := identityHolder.Sign(signData)
	assert.NoError(t, err)

	expectedSignature, err := keyHolder.Sign(signData)
	assert.NoError(t, err)

	assert.Equal(t, expectedSignature, extractedSignature)

}
