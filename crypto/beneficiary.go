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
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// SetBeneficiaryRequest represents a request for setting new beneficiary.
type SetBeneficiaryRequest struct {
	ChannelID   string `json:"channelID"`
	Beneficiary string `json:"beneficiary"`
	Nonce       uint64 `json:"nonce"`
	Signature   string `json:"signature"`
}

func CreateBeneficiaryRequest(channelID, beneficiary string, nonce uint64, ks hashSigner, signer common.Address) (*SetBeneficiaryRequest, error) {
	if hasHexPrefix(channelID) {
		channelID = channelID[2:]
	}

	if hasHexPrefix(beneficiary) {
		beneficiary = beneficiary[2:]
	}

	if !isHex(channelID) || !isHex(beneficiary) {
		return nil, errors.New("channelID and beneficiary have to be proper hex strings")
	}

	request := SetBeneficiaryRequest{
		ChannelID:   strings.ToLower(channelID),
		Beneficiary: strings.ToLower(beneficiary),
		Nonce:       nonce,
	}

	signature, err := request.CreateSignature(ks, signer)
	if err != nil {
		return nil, err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}

	request.Signature = hex.EncodeToString(signature)

	return &request, nil
}

func NewBeneficiaryRequest(channelID, beneficiary string, nonce uint64, signature string) (*SetBeneficiaryRequest, error) {
	if hasHexPrefix(channelID) {
		channelID = channelID[2:]
	}

	if hasHexPrefix(beneficiary) {
		beneficiary = beneficiary[2:]
	}

	if !isHex(channelID) || !isHex(beneficiary) {
		return nil, errors.New("channelID and beneficiary have to be proper hex strings")
	}

	return &SetBeneficiaryRequest{
		ChannelID:   channelID,
		Beneficiary: beneficiary,
		Nonce:       nonce,
		Signature:   signature,
	}, nil
}

// CreateSignature signs set beneficiary request using keystore
func (r SetBeneficiaryRequest) CreateSignature(ks hashSigner, signer common.Address) ([]byte, error) {
	message := r.GetMessage()
	hash := crypto.Keccak256(message)
	return ks.SignHash(
		accounts.Account{Address: signer},
		hash,
	)
}

// GetSignatureBytesRaw returns the unadulterated bytes of the signature.
func (r SetBeneficiaryRequest) GetSignatureBytesRaw() []byte {
	signature := strings.TrimPrefix(r.Signature, "0x")
	signBytes := common.Hex2Bytes(signature)

	return signBytes
}

// GetMessage forms the message payload given the set beneficiary request.
func (r SetBeneficiaryRequest) GetMessage() []byte {
	message := []byte{}
	message = append(message, common.Hex2Bytes(strings.TrimPrefix(r.ChannelID, "0x"))...)
	message = append(message, common.HexToAddress(r.Beneficiary).Bytes()...)
	message = append(message, Pad(abi.U256(big.NewInt(0).SetUint64(r.Nonce)), 32)...)

	return message
}

// RecoverSigner recovers the signer identity from the given request.
func (r SetBeneficiaryRequest) RecoverSigner() (common.Address, error) {
	signature := r.GetSignatureBytesRaw()

	err := ReformatSignatureVForRecovery(signature)
	if err != nil {
		return common.Address{}, err
	}

	recoveredAddress, err := RecoverAddress(r.GetMessage(), signature)
	if err != nil {
		return common.Address{}, err
	}

	return recoveredAddress, nil
}
