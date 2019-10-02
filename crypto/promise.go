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

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// Promise is payment promise object
type Promise struct {
	ChannelID []byte
	Amount    uint64
	Fee       uint64
	Hashlock  []byte
	R         []byte
	Signature []byte
}

// CreatePromise creates new payment promise
func CreatePromise(channelID string, amount uint64, fee uint64, hashlock string, ks *keystore.KeyStore, signer common.Address) (*Promise, error) {
	if hasHexPrefix(channelID) {
		channelID = channelID[2:]
	}

	if !isHex(channelID) || !isHex(hashlock) {
		return nil, errors.New("channelID and hashlock have to be proper hex strings")
	}

	chID, err := hex.DecodeString(channelID)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding channelID")
	}

	hl, err := hex.DecodeString(hashlock)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding hashlock")
	}

	promise := Promise{
		ChannelID: chID,
		Amount:    amount,
		Fee:       fee,
		Hashlock:  hl,
	}

	signature, err := promise.CreateSignature(ks, signer)
	if err != nil {
		return nil, err
	}

	ReformatSignatureVForBC(signature)
	promise.Signature = signature

	return &promise, nil
}

// GetMessage forms the message of payment promise
func (p Promise) GetMessage() []byte {
	message := []byte{}
	message = append(message, p.ChannelID...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(p.Amount)), 32)...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(p.Fee)), 32)...)
	message = append(message, Pad(p.Hashlock, 32)...)
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

// GetSignatureHexString returns signature in hex sting format
func (p Promise) GetSignatureHexString() string {
	return "0x" + hex.EncodeToString(p.Signature)
}

// IsPromiseValid validates if given promise params are properly signed
func (p Promise) IsPromiseValid(expectedSigner common.Address) bool {
	sig := make([]byte, 65)
	copy(sig, p.Signature)

	err := ReformatSignatureVForRecovery(sig)
	if err != nil {
		return false
	}

	recoveredSigner, err := RecoverAddress(p.GetMessage(), sig)
	if err != nil {
		return false
	}

	return recoveredSigner == expectedSigner
}

// RecoverSigner recovers signer address out of promise signature
func (p Promise) RecoverSigner() (common.Address, error) {
	sig := make([]byte, 65)
	copy(sig, p.Signature)

	ReformatSignatureVForRecovery(sig)
	return RecoverAddress(p.GetMessage(), sig)
}
