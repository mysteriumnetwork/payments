package crypto

import (
	"encoding/hex"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func getExchangeMessage() ExchangeMessage {
	return ExchangeMessage{
		Promise:        getPromise(),
		AgreementID:    uint64(1),
		AgreementTotal: uint64(10),
		Provider:       "0x1a9fab9aba871ed0f5bff28f9f9e52d374376611",
		Signature:      "1b6412d22cd2322408fa1f55506c99a6b2901e2f1e9605685c1cc6a4dfb32f0e6c9aceac46bf1447f22e3b7a596a90d464a0fc9d56d9a6ba83e93521793493df1c",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("7a14ba6d4315bf50f2e6570574fd734edca97fc96bd446c8347efa46709b04d6")
	assert.Equal(t, expectedHash, message.GetMessageHash())
}

func TestRecoverConsumerIdentity(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0x0d71535454f4fc153e545c3fc7cfc412ad7782c8")
	recoveredSigner, err := message.RecoverConsumerIdentity()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

func TestValidateExchangeMessage(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0x0d71535454f4fc153e545c3fc7cfc412ad7782c8")
	assert.True(t, message.ValidateExchangeMessage(expectedSigner))
}

func TestCreateExchangeMessage(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey(), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	channelID, amount, fee, hashlock, promiseSignature, messageSignature, provider := getParams()

	agreementID := uint64(1)
	agreementTotal := uint64(10)
	r, _ := hex.DecodeString("23320e3545e5005c8af0cf8130dd156023178c58da0347090d4bb02b6e0938a9")

	invoice := CreateInvoice(agreementID, agreementTotal, fee, r)
	invoice.Provider = provider

	message, err := CreateExchangeMessage(invoice, amount, channelID, ks, account.Address)
	assert.Nil(t, err)

	assert.Equal(t, promiseSignature, message.Promise.Signature)
	assert.Equal(t, messageSignature, message.Signature)
	assert.Equal(t, hashlock, message.Promise.Hashlock)
}
