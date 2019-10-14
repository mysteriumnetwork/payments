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
		Signature:      "d44920d4e0bcb96e836f0731f168e862a9efcd4e72dd093141a3c95205ba6cc115e23fcbfb8625e219be4255e0f9597ef55699d13a207a21881378329b38d0b31c",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("fcfeaa61ec8cc9328725b7829c28389457eb0dc65b4daf8ee1e974f213ce6da1")
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
