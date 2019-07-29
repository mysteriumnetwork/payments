/*
 * Copyright (C) 2019 The "MysteriumNetwork/payments" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package crypto

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ExchangeMessage represent a promise exchange message
type ExchangeMessage struct {
	Promise        Promise
	AgreementID    uint64
	AgreementTotal uint64
	Provider       string
	Signature      string
}

// CreateExchangeMessage creates new exchange message with it's promise
func CreateExchangeMessage(invoice Invoice, promiseAmount uint64, channelID string, ks *keystore.KeyStore, signer common.Address) (ExchangeMessage, error) {
	promise, err := CreatePromise(channelID, promiseAmount, invoice.Fee, invoice.Hashlock, ks, signer)
	if err != nil {
		return ExchangeMessage{}, err
	}

	message := ExchangeMessage{
		Promise:        promise,
		AgreementID:    invoice.AgreementID,
		AgreementTotal: invoice.AgreementTotal,
		Provider:       invoice.Provider,
	}

	signature, err := message.CreateSignature(ks, signer)
	if err != nil {
		return ExchangeMessage{}, err
	}

	ReformatSignatureVForBC(signature)
	message.Signature = hex.EncodeToString(signature)

	return message, nil
}

// GetAgreementTotal returns a big int representation for the agreement total amount
func (m ExchangeMessage) GetAgreementTotal() *big.Int {
	return big.NewInt(0).SetUint64(m.AgreementTotal)
}

// GetFee returns the big int representation for the promise settling transaction fee
func (m ExchangeMessage) GetFee() *big.Int {
	return big.NewInt(0).SetUint64(m.Promise.Fee)
}

// GetSignatureBytesRaw returns the unadulterated bytes of the signature
func (m ExchangeMessage) GetSignatureBytesRaw() []byte {
	signature := strings.TrimPrefix(m.Signature, "0x")
	signBytes := common.Hex2Bytes(signature)
	return signBytes
}

// GetMessage forms the message of promise exchange request
func (m ExchangeMessage) GetMessage() []byte {
	message := []byte{}
	message = append(message, m.Promise.GetHash()...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(m.AgreementID)), 32)...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(m.AgreementTotal)), 32)...)
	message = append(message, common.HexToAddress(m.Provider).Bytes()...)
	return message
}

// GetMessageHash returns a keccak of exchange message params
func (m ExchangeMessage) GetMessageHash() []byte {
	return crypto.Keccak256(m.GetMessage())
}

// CreateSignature signs promise using keystore
func (m ExchangeMessage) CreateSignature(ks *keystore.KeyStore, signer common.Address) ([]byte, error) {
	return ks.SignHash(
		accounts.Account{Address: signer},
		m.GetMessageHash(),
	)
}

// RecoverConsumerIdentity recovers the identity from the given request
func (m ExchangeMessage) RecoverConsumerIdentity() (common.Address, error) {
	signature := m.GetSignatureBytesRaw()

	err := ReformatSignatureVForRecovery(signature)
	if err != nil {
		return common.Address{}, err
	}

	return RecoverAddress(m.GetMessage(), signature)
}

// ValidateExchangeMessage validates if given exchange message was signed by expected identity
func (m ExchangeMessage) ValidateExchangeMessage(expectedSigner common.Address) bool {
	recoveredSigner, err := m.RecoverConsumerIdentity()
	if err != nil {
		return false
	}

	return recoveredSigner == expectedSigner
}
