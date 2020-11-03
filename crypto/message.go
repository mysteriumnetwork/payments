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
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

// ExchangeMessage represent a promise exchange message
type ExchangeMessage struct {
	Promise        Promise
	AgreementID    *big.Int
	AgreementTotal *big.Int
	Provider       string
	Signature      string
	HermesID       string
	ChainID        int64
}

type hashSigner interface {
	SignHash(a accounts.Account, hash []byte) ([]byte, error)
}

func CreateExchangeMessageWithPromise(chainID int64, invoice Invoice, promise *Promise, hermesID string, ks hashSigner, signer common.Address) (*ExchangeMessage, error) {
	message := ExchangeMessage{
		Promise:        *promise,
		AgreementID:    new(big.Int).Set(invoice.AgreementID),
		AgreementTotal: new(big.Int).Set(invoice.AgreementTotal),
		Provider:       invoice.Provider,
		HermesID:       hermesID,
		ChainID:        chainID,
	}

	signature, err := message.CreateSignature(ks, signer)
	if err != nil {
		return nil, err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}

	message.Signature = hex.EncodeToString(signature)

	return &message, nil
}

// CreateExchangeMessage creates new exchange message with it's promise
func CreateExchangeMessage(chainID int64, invoice Invoice, promiseAmount *big.Int, channelID, hermesID string, ks hashSigner, signer common.Address) (*ExchangeMessage, error) {
	promise, err := CreatePromise(channelID, chainID, promiseAmount, invoice.TransactorFee, invoice.Hashlock, ks, signer)
	if err != nil {
		return nil, err
	}

	message := ExchangeMessage{
		Promise:        *promise,
		AgreementID:    new(big.Int).Set(invoice.AgreementID),
		AgreementTotal: new(big.Int).Set(invoice.AgreementTotal),
		Provider:       invoice.Provider,
		HermesID:       hermesID,
		ChainID:        invoice.ChainID,
	}

	signature, err := message.CreateSignature(ks, signer)
	if err != nil {
		return nil, err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}

	message.Signature = hex.EncodeToString(signature)

	return &message, nil
}

// GetAgreementTotal returns a big int representation for the agreement total amount
func (m ExchangeMessage) GetAgreementTotal() *big.Int {
	return m.AgreementTotal
}

// GetFee returns the big int representation for the promise settling transaction fee
func (m ExchangeMessage) GetFee() *big.Int {
	return m.Promise.Fee
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
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(m.ChainID))
	message = append(message, Pad(b, 32)...)
	message = append(message, m.Promise.GetHash()...)
	message = append(message, Pad(math.U256(m.AgreementID).Bytes(), 32)...)
	message = append(message, Pad(math.U256(m.AgreementTotal).Bytes(), 32)...)
	message = append(message, common.HexToAddress(m.Provider).Bytes()...)

	// TODO: once all the consumers upgrade, this check needs to go to
	if m.HermesID != "" {
		message = append(message, common.HexToAddress(m.HermesID).Bytes()...)
	}

	return message
}

// GetMessageHash returns a keccak of exchange message params
func (m ExchangeMessage) GetMessageHash() []byte {
	return crypto.Keccak256(m.GetMessage())
}

// CreateSignature signs promise using keystore
func (m ExchangeMessage) CreateSignature(ks hashSigner, signer common.Address) ([]byte, error) {
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

// IsMessageValid validates if given exchange message was signed by expected identity
func (m ExchangeMessage) IsMessageValid(expectedSigner common.Address) bool {
	recoveredSigner, err := m.RecoverConsumerIdentity()
	if err != nil {
		return false
	}

	return recoveredSigner == expectedSigner
}
