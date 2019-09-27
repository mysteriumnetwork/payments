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
		Promise:        getPromise("consumer"),
		AgreementID:    uint64(1),
		AgreementTotal: uint64(1401),
		Provider:       "0xf10021ba3b10d023e671668d20daeff821561d09",
		Signature:      "a7a201c9ec67d5b627cda20196f80a86e9e03c9dd9e8224a73470605ef40494847dbf5f6d2701c58d9093294fc5cfdbd98e85331c191d49cd1da29d96b0c10f81c",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("f0fcbf8b48bd5b10696113c7f51cf5b06f16e89cfb7d3ebae58dfd9612234919")
	assert.Equal(t, expectedHash, message.GetMessageHash())
}

func TestRecoverConsumerIdentity(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	recoveredSigner, err := message.RecoverConsumerIdentity()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

func TestExchangeMessageValidation(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := common.HexToAddress("0xf53acdd584ccb85ee4ec1590007ad3c16fdff057")
	assert.True(t, message.IsMessageValid(expectedSigner))

	wrongSigner := common.HexToAddress("0xf10021ba3b10d023e671668d20daeff821561d09")
	assert.False(t, message.IsMessageValid(wrongSigner))
}

func TestCreateExchangeMessage(t *testing.T) {
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
	provider := p.Provider.Hex()

	agreementID := uint64(1)
	agreementTotal := uint64(1401)
	r, _ := hex.DecodeString("5b6b3f31a3acd0e317173d25c8b60503547b741a0e81d6068bb88486967839fa")

	invoice := CreateInvoice(agreementID, agreementTotal, fee, r)
	invoice.Provider = provider

	message, err := CreateExchangeMessage(invoice, amount, channelID, ks, account.Address)
	assert.Nil(t, err)

	assert.Equal(t, p.PromiseSignature, message.Promise.Signature)
	assert.Equal(t, p.ExchangeMessageSignature, message.Signature)
	assert.Equal(t, p.Hashlock, message.Promise.Hashlock)
}
