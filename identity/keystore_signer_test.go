package identity

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestKeyStoreSignerProvidesValidSignature(t *testing.T) {

	assert.NoError(t, os.RemoveAll("testoutput"))

	sampleKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	ks := keystore.NewKeyStore("testoutput", keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.ImportECDSA(sampleKey, "")
	assert.NoError(t, err)

	assert.NoError(t, ks.Unlock(acc, ""))
	signer := NewKeystoreSigner(ks, acc.Address)

	signData := []byte("Testing signature")

	extractedSignature, err := signer.Sign(signData)
	assert.NoError(t, err)

	pubKeyBytes, err := crypto.Ecrecover(crypto.Keccak256(signData), extractedSignature)
	assert.NoError(t, err)

	recoveredIdentity := crypto.PubkeyToAddress(*crypto.ToECDSAPub(pubKeyBytes))

	assert.Equal(t, acc.Address, recoveredIdentity)

}
