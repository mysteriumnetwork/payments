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

	ks := keystore.NewKeyStore("testoutput", keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.ImportECDSA(sampleKey, "")
	assert.NoError(t, err)

	assert.NoError(t, ks.Unlock(acc, ""))
	publicKeyProvider := FromKeystore(ks, acc.Address)
	extractedPubKey, err := publicKeyProvider.GetPublicKey()
	assert.NoError(t, err)

	assert.Equal(t, sampleKey.PublicKey, extractedPubKey)
}
