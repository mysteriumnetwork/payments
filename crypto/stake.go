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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

const stakeReturnPrefix = "Stake return request"

// DecreaseProviderStakeRequest represents all the parameters required for decreasing provider stake.
type DecreaseProviderStakeRequest struct {
	ChannelID     [32]byte
	HermesID      common.Address
	Amount        *big.Int
	TransactorFee *big.Int
	Signature     []byte
}

// CreateSignature signs promise using keystore
func (dpsr DecreaseProviderStakeRequest) CreateSignature(ks hashSigner, signer common.Address) ([]byte, error) {
	message := dpsr.GetMessage()
	hash := crypto.Keccak256(message)
	return ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
}

// Sign signs the DecreaseProviderStakeRequest.
func (dpsr *DecreaseProviderStakeRequest) Sign(ks *keystore.KeyStore, signer common.Address) error {
	signature, err := dpsr.CreateSignature(ks, signer)
	if err != nil {
		return err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return fmt.Errorf("failed to reformat signature: %w", err)
	}

	dpsr.Signature = signature

	return nil
}

// GetMessage gets a message representation of the DecreaseProviderStakeRequest.
func (dpsr DecreaseProviderStakeRequest) GetMessage() []byte {
	msg := []byte{}
	msg = append(msg, []byte(stakeReturnPrefix)...)
	msg = append(msg, Pad(dpsr.ChannelID[:], 32)...)
	msg = append(msg, Pad(math.U256(dpsr.Amount).Bytes(), 32)...)
	msg = append(msg, Pad(math.U256(dpsr.TransactorFee).Bytes(), 32)...)
	return msg
}

// IsValid checks if the current stake decrease request is valid or not.
func (dpsr DecreaseProviderStakeRequest) IsValid(expectedSigner common.Address) bool {
	recoveredSigner, err := dpsr.RecoverSigner()
	if err != nil {
		return false
	}

	return recoveredSigner == expectedSigner
}

// RecoverSigner recovers signer address out of request signature.
func (dpsr DecreaseProviderStakeRequest) RecoverSigner() (common.Address, error) {
	sig := make([]byte, 65)
	copy(sig, dpsr.Signature)

	err := ReformatSignatureVForRecovery(sig)
	if err != nil {
		return common.Address{}, err
	}

	return RecoverAddress(dpsr.GetMessage(), sig)
}
