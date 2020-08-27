package crypto

import (
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestReferralRequest(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("consumer"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	req, err := CreateReferralTokenRequest(ks, account.Address)
	assert.NoError(t, err)

	err = req.IsValid()
	assert.NoError(t, err)

	// should fail if we change identity
	req.Identity = common.HexToAddress("0x0")
	err = req.IsValid()
	assert.Error(t, err)
	req.Identity = account.Address

	// should fail if we change the signature
	req.Signature = strings.Replace(req.Signature, "a", "b", -1)
	err = req.IsValid()
	assert.Error(t, err)
}
