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
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Promise is payment promise object
type Promise struct {
	ChannelID string
	Amount    uint64
	Fee       uint64
	Hashlock  string
	R         string
	Signature string
}

// CreatePromise creates new promise
func CreatePromise(channelID string, amount uint64, fee uint64, ks *keystore.KeyStore, signer common.Address) (Promise, error) {
	// TODO validate channelID, it have top be proper address, or request here address already

	r := make([]byte, 64)
	rand.Read(r)

	promise := Promise{
		ChannelID: channelID,
		Amount:    amount,
		Fee:       fee,
		Hashlock:  hex.EncodeToString(crypto.Keccak256(r)),
		R:         hex.EncodeToString(r),
	}

	signature, _ := promise.CreateSignature(ks, signer)
	promise.Signature = hex.EncodeToString(signature)

	return promise, nil
}

// GetMessage forms the message of payment promise
func (p Promise) GetMessage() []byte {
	hashlock, _ := hex.DecodeString(p.Hashlock)

	message := []byte{}
	message = append(message, common.HexToAddress(p.ChannelID).Bytes()...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(p.Amount)), 32)...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(p.Fee)), 32)...)
	message = append(message, hashlock...)
	return message
}

// GetHash returns a keccak of payment promise message
func (p Promise) GetHash() []byte {
	return crypto.Keccak256(p.GetMessage())
}

// CreateSignature signs promise using keystore
func (p Promise) CreateSignature(ks *keystore.KeyStore, signer common.Address) ([]byte, error) {
	message := p.GetMessage()
	hash := crypto.Keccak256(message)
	return ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
}

// CreateSignature2 signs promise using private key
func (p Promise) CreateSignature2(pk *ecdsa.PrivateKey) ([]byte, error) {
	message := p.GetMessage()
	hash := crypto.Keccak256Hash(message)
	signature, err := crypto.Sign(hash.Bytes(), pk)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// GetSignatureBytesRaw returns the unadulterated bytes of the signature
func (p Promise) GetSignatureBytesRaw() []byte {
	signature := strings.TrimPrefix(p.Signature, "0x")
	signBytes := common.Hex2Bytes(signature)
	return signBytes
}

// ValidatePromise validates if given promise params are properly signed
func (p Promise) ValidatePromise(expectedSigner common.Address) bool {
	signature := p.GetSignatureBytesRaw()
	err := ReformatSignatureVForRecovery(signature)
	if err != nil {
		return false
	}

	recoveredSigner, err := RecoverAddress(p.GetMessage(), signature)
	if err != nil {
		return false
	}

	return recoveredSigner == expectedSigner
}

// RecoverSigner recovers signer address out of promise signature
func (p Promise) RecoverSigner() (common.Address, error) {
	signature := p.GetSignatureBytesRaw()
	ReformatSignatureVForRecovery(signature)
	return RecoverAddress(p.GetMessage(), signature)
}
