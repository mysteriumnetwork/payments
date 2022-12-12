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

	"github.com/stretchr/testify/assert"
)

func getExchangeMessage() ExchangeMessage {
	return ExchangeMessage{
		Promise:        getPromise("consumer"),
		AgreementID:    big.NewInt(1),
		AgreementTotal: big.NewInt(1401),
		Provider:       "0xf10021ba3b10d023e671668d20daeff821561d09",
		Signature:      "3310f94760eeb288598d69aaf6214e123766d3cc988f5d3c6a77d5d5b70dc4b6617eeeb0f49df467527ed4a9cdf6f5c7ae3a1477a84f853b9ef8587607d0be431b",
		ChainID:        1,
		HermesID:       "0xe10021ba3b10d023e671668d20daeff821561d09",
	}
}

func TestGetMessageHash(t *testing.T) {
	message := getExchangeMessage()
	expectedHash, _ := hex.DecodeString("e6b89a054f9f1e820fbc06d727209682adca2d6d24b771c9a82edd1809657e39")
	assert.Equal(t, expectedHash, message.GetMessageHash())
}

func TestRecoverConsumerIdentity(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := HexToAddress("0xCcad590A7a938Cb086e7414e0F0000eD6a56D833")
	recoveredSigner, err := message.RecoverConsumerIdentity()
	assert.Nil(t, err)
	assert.Equal(t, expectedSigner, recoveredSigner)
}

func TestExchangeMessageValidation(t *testing.T) {
	message := getExchangeMessage()
	expectedSigner := HexToAddress("0xCcad590A7a938Cb086e7414e0F0000eD6a56D833")
	assert.True(t, message.IsMessageValid(expectedSigner))

	wrongSigner := HexToAddress("0xf10021ba3b10d023e671668d20daeff821561d09")
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

	invoice, err := CreateInvoice(agreementID, agreementTotal, fee, r, 1)
	assert.NoError(t, err)
	invoice.Provider = provider

	message, err := CreateExchangeMessage(1, invoice, amount, channelID, "", ks, FromCommonAddress(account.Address))
	assert.Nil(t, err)

	assert.Equal(t, p.PromiseSignature, message.Promise.Signature)
	assert.Equal(t, p.ExchangeMessageSignature, message.Signature)
	assert.Equal(t, p.Hashlock, message.Promise.Hashlock)
	assert.Len(t, p.R, 32)
}

func TestCreateExchangeMessageWithPromise(t *testing.T) {
	dir, ks := tmpKeyStore(t, false)
	defer os.RemoveAll(dir)

	account, err := ks.ImportECDSA(getPrivKey("consumer"), "")
	assert.Nil(t, err)
	if err := ks.Unlock(account, ""); err != nil {
		t.Fatal(err)
	}

	p := getParams("consumer")
	fee := big.NewInt(0).SetUint64(p.Fee)
	provider := p.Provider.Hex()

	agreementID := big.NewInt(1)
	agreementTotal := big.NewInt(1401)
	r, _ := hex.DecodeString("5b6b3f31a3acd0e317173d25c8b60503547b741a0e81d6068bb88486967839fa")

	invoice, err := CreateInvoice(agreementID, agreementTotal, fee, r, 1)
	assert.NoError(t, err)
	invoice.Provider = provider

	promise := getPromise("consumer")
	message, err := CreateExchangeMessageWithPromise(1, invoice, &promise, "", ks, FromCommonAddress(account.Address))
	assert.Nil(t, err)

	assert.Equal(t, p.PromiseSignature, message.Promise.Signature)
	assert.Equal(t, p.ExchangeMessageSignature, message.Signature)
	assert.Equal(t, p.Hashlock, message.Promise.Hashlock)
	assert.Len(t, p.R, 32)
	assert.Equal(t, promise, message.Promise)
	assert.Equal(t, agreementTotal, message.GetAgreementTotal())
	assert.Equal(t, fee, message.GetFee())
}
