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

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// Promise is payment promise object
type Promise struct {
	ChannelID []byte
	ChainID   int64
	Amount    *big.Int
	Fee       *big.Int
	Hashlock  []byte
	R         []byte
	Signature []byte
}

// CreatePromise creates and signs new payment promise
func CreatePromise(channelID string, chainID int64, amount *big.Int, fee *big.Int, hashlock string, ks hashSigner, signer common.Address) (*Promise, error) {
	channelID = ensureNoPrefix(channelID)
	hashlock = ensureNoPrefix(hashlock)

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
		ChainID:   chainID,
	}

	signature, err := promise.CreateSignature(ks, signer)
	if err != nil {
		return nil, err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}

	promise.Signature = signature

	return &promise, nil
}

// NewPromise will create new promise,
// signature can be empty and be created later using `Sign()` method.
func NewPromise(chainID int64, channelID string, amount, fee *big.Int, preimage string, signature string) (*Promise, error) {
	channelID = ensureNoPrefix(channelID)
	preimage = ensureNoPrefix(preimage)

	chID, err := hex.DecodeString(channelID)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding channelID")
	}

	r, err := hex.DecodeString(preimage)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding preimage")
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding signature")
	}

	// hashlock := crypto.Keccak256(r)

	promise := Promise{
		ChannelID: chID,
		Amount:    new(big.Int).Set(amount),
		Fee:       new(big.Int).Set(fee),
		Hashlock:  crypto.Keccak256(r),
		R:         r,
		Signature: sig,
		ChainID:   chainID,
	}

	return &promise, nil
}

// NewRawPromise creates a promise from given params. The promise has no R.
func NewRawPromise(chainID int64, channelID string, amount, fee *big.Int, hashlock string, signature string) (*Promise, error) {
	channelID = ensureNoPrefix(channelID)
	hashlock = ensureNoPrefix(hashlock)

	chID, err := hex.DecodeString(channelID)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding channelID")
	}

	hl, err := hex.DecodeString(hashlock)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding preimage")
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return nil, errors.Wrap(err, "Problem in decoding signature")
	}

	promise := Promise{
		ChannelID: chID,
		Amount:    new(big.Int).Set(amount),
		Fee:       new(big.Int).Set(fee),
		Hashlock:  hl,
		Signature: sig,
		ChainID:   chainID,
	}

	return &promise, nil
}

// Sign signs promise with given keystore and signer
func (p *Promise) Sign(ks *keystore.KeyStore, signer common.Address) error {
	signature, err := p.CreateSignature(ks, signer)
	if err != nil {
		return err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return fmt.Errorf("failed to reformat signature: %w", err)
	}

	p.Signature = signature

	return nil
}

// GetMessage forms the message of payment promise
func (p Promise) GetMessage() []byte {
	message := []byte{}
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(p.ChainID))
	message = append(message, Pad(b, 32)...)
	message = append(message, Pad(p.ChannelID, 32)...)
	message = append(message, Pad(math.U256(p.Amount).Bytes(), 32)...)
	message = append(message, Pad(math.U256(p.Fee).Bytes(), 32)...)
	message = append(message, Pad(p.Hashlock, 32)...)
	return message
}

// GetHash returns a keccak of payment promise message
func (p Promise) GetHash() []byte {
	return crypto.Keccak256(p.GetMessage())
}

// CreateSignature signs promise using keystore
func (p Promise) CreateSignature(ks hashSigner, signer common.Address) ([]byte, error) {
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

	err := ReformatSignatureVForRecovery(sig)
	if err != nil {
		return common.Address{}, err
	}
	return RecoverAddress(p.GetMessage(), sig)
}
