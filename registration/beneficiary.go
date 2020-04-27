/*
 * Copyright (C) 2020 The "MysteriumNetwork/payments" Authors.
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

package registration

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/crypto"
)

// SetBeneficiaryRequest represents a request for setting new beneficiary.
type SetBeneficiaryRequest struct {
	AccountantID string `json:"accountantID"`
	ChannelID    string `json:"channelID"`
	Beneficiary  string `json:"beneficiary"`
	Nonce        uint64 `json:"nonce"`
	Signature    string `json:"signature"`
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
	message = append(message, crypto.Pad(abi.U256(big.NewInt(0).SetUint64(r.Nonce)), 32)...)

	return message
}

// RecoverIdentity recovers the identity from the given request.
func (r SetBeneficiaryRequest) RecoverIdentity() (common.Address, error) {
	signature := r.GetSignatureBytesRaw()

	err := crypto.ReformatSignatureVForRecovery(signature)
	if err != nil {
		return common.Address{}, err
	}

	recoveredAddress, err := crypto.RecoverAddress(r.GetMessage(), signature)
	if err != nil {
		return common.Address{}, err
	}

	return recoveredAddress, nil
}
