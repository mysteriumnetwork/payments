/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package crypto

import (
	"encoding/hex"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func getExchangeMessage() ExchangeMessage {
	return ExchangeMessage{
		Promise:        getPromise("consumer"),
		AgreementID:    big.NewInt(1),
		AgreementTotal: big.NewInt(1401),
		Provider:       "0xf10021ba3b10d023e671668d20daeff821561d09",
		Signature:      "56493421bd2772cca2ba970da27396e103a08027f1ce49de974f789e322b0d7a3f52b9dd745a34bfa2f330ba2d3c442867ebb3753d1f206811ab572ab7d482dc1b",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("c39bda2f6271776edb9d5780210e5f46cfbda7df07d409277cadb452f45cc4ee")
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
	amount := big.NewInt(0).SetUint64(p.Amount)
	fee := big.NewInt(0).SetUint64(p.Fee)
	provider := p.Provider.Hex()

	agreementID := big.NewInt(1)
	agreementTotal := big.NewInt(1401)
	r, _ := hex.DecodeString("5b6b3f31a3acd0e317173d25c8b60503547b741a0e81d6068bb88486967839fa")

	invoice := CreateInvoice(agreementID, agreementTotal, fee, r)
	invoice.Provider = provider

	message, err := CreateExchangeMessage(1, invoice, amount, channelID, "", ks, account.Address)
	assert.Nil(t, err)

	assert.Equal(t, p.PromiseSignature, message.Promise.Signature)
	assert.Equal(t, p.ExchangeMessageSignature, message.Signature)
	assert.Equal(t, p.Hashlock, message.Promise.Hashlock)
}
