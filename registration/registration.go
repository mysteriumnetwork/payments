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

package registration

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/mysteriumnetwork/payments/v3/crypto"
)

// Request represent a request to register
type Request struct {
	ChainID         int64    `json:"chainID"`
	HermesID        string   `json:"hermesID"`
	Stake           *big.Int `json:"stake"`
	Fee             *big.Int `json:"fee"`
	Beneficiary     string   `json:"beneficiary"`
	Signature       string   `json:"signature"`
	RegistryAddress string   `json:"registryAddress"`
}

// GetStakeAmount returns a big int representation for the stake amount
func (r Request) GetStakeAmount() *big.Int {
	return r.Stake
}

// GetFee returns the big int representation for the fee
func (r Request) GetFee() *big.Int {
	return r.Fee
}

// GetMessage forms the message payload given the registration request
func (r Request) GetMessage() []byte {
	message := []byte{}
	message = append(message, crypto.Pad(math.U256Bytes(big.NewInt(r.ChainID)), 32)...)
	message = append(message, common.HexToAddress(r.RegistryAddress).Bytes()...)
	message = append(message, common.HexToAddress(r.HermesID).Bytes()...)
	message = append(message, crypto.Pad(math.U256(r.Stake).Bytes(), 32)...)
	message = append(message, crypto.Pad(math.U256(r.Fee).Bytes(), 32)...)
	message = append(message, common.HexToAddress(r.Beneficiary).Bytes()...)
	return message
}

// RecoverIdentity recovers the identity from the given request
func (r Request) RecoverIdentity() (common.Address, error) {
	signature := GetSignatureBytesRaw(r.Signature)

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

type OpenConsumerChannelRequest struct {
	ChainID         int64    `json:"chainID"`
	HermesID        string   `json:"hermesID"`
	Signature       string   `json:"signature"`
	TransactorFee   *big.Int `json:"fee"`
	RegistryAddress string   `json:"registryAddress"`
}

func (r *OpenConsumerChannelRequest) GetMessage() []byte {
	message := []byte{}
	message = append(message, crypto.Pad(math.U256Bytes(big.NewInt(r.ChainID)), 32)...)
	message = append(message, common.HexToAddress(r.RegistryAddress).Bytes()...)
	message = append(message, common.HexToAddress(r.HermesID).Bytes()...)
	message = append(message, crypto.Pad(math.U256(r.TransactorFee).Bytes(), 32)...)

	return message
}

// RecoverIdentity recovers the identity from the given request
func (r *OpenConsumerChannelRequest) RecoverIdentity() (common.Address, error) {
	signature := common.Hex2Bytes(strings.TrimPrefix(r.Signature, "0x"))

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

// GetSignatureBytesRaw returns the unadulterated bytes of the signature
func GetSignatureBytesRaw(sig string) []byte {
	return common.Hex2Bytes(strings.TrimPrefix(sig, "0x"))
}
