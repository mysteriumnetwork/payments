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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/crypto"
)

// Request represent a request to register
type Request struct {
	AccountantID    string
	Stake           uint64
	Fee             uint64
	Beneficiary     string
	Signature       string
	RegistryAddress string
}

// GetLoanAmount retursn a big int representation for the loan amount
func (r Request) GetLoanAmount() *big.Int {
	return big.NewInt(0).SetUint64(r.Stake)
}

// GetFee returns the big int representation for the fee
func (r Request) GetFee() *big.Int {
	return big.NewInt(0).SetUint64(r.Fee)
}

// GetSignatureBytesRaw returns the unadulterated bytes of the signature
func (r Request) GetSignatureBytesRaw() []byte {
	signature := strings.TrimPrefix(r.Signature, "0x")
	signBytes := common.Hex2Bytes(signature)
	return signBytes
}

// GetMessage forms the message payload given the registration request
func (r Request) GetMessage() []byte {
	message := []byte{}
	message = append(message, common.HexToAddress(r.RegistryAddress).Bytes()...)
	message = append(message, common.HexToAddress(r.AccountantID).Bytes()...)
	message = append(message, crypto.Pad(abi.U256(big.NewInt(0).SetUint64(r.Stake)), 32)...)
	message = append(message, crypto.Pad(abi.U256(big.NewInt(0).SetUint64(r.Fee)), 32)...)
	message = append(message, common.HexToAddress(r.Beneficiary).Bytes()...)
	return message
}

// RecoverIdentity recovers the identity from the given request
func (r Request) RecoverIdentity() (common.Address, error) {
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
