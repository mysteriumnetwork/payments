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

// SetBeneficiaryRequest represents a request for setting new beneficiary.
type SetBeneficiaryRequest struct {
	ChainID     int64    `json:"chainID"`
	Registry    string   `json:"registry"`
	Identity    string   `json:"channelID"`
	Beneficiary string   `json:"beneficiary"`
	Nonce       *big.Int `json:"nonce"`
	Signature   string   `json:"signature"`
}

func CreateBeneficiaryRequest(chainID int64, identity, registry, beneficiary string, nonce *big.Int, ks hashSigner, signer Address) (*SetBeneficiaryRequest, error) {
	req, err := NewBeneficiaryRequest(chainID, identity, registry, beneficiary, nonce, "")
	if err != nil {
		return nil, err
	}

	signature, err := req.CreateSignature(ks, signer)
	if err != nil {
		return nil, err
	}

	if err := ReformatSignatureVForBC(signature); err != nil {
		return nil, fmt.Errorf("failed to reformat signature: %w", err)
	}

	req.Signature = hex.EncodeToString(signature)
	return req, nil
}

func NewBeneficiaryRequest(chainID int64, identity, registry, beneficiary string, nonce *big.Int, signature string) (*SetBeneficiaryRequest, error) {
	return &SetBeneficiaryRequest{
		ChainID:     chainID,
		Identity:    ensureNoPrefix(strings.ToLower(identity)),
		Registry:    ensureNoPrefix(strings.ToLower(registry)),
		Beneficiary: ensureNoPrefix(strings.ToLower(beneficiary)),
		Nonce:       nonce,
		Signature:   signature,
	}, nil
}

// CreateSignature signs set beneficiary request using keystore
func (r SetBeneficiaryRequest) CreateSignature(ks hashSigner, signer Address) ([]byte, error) {
	message := r.GetMessage()
	hash := crypto.Keccak256(message)
	return ks.SignHash(
		accounts.Account{Address: signer.ToCommon()},
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
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(r.ChainID))
	message = append(message, Pad(b, 32)...)
	message = append(message, common.Hex2Bytes(strings.TrimPrefix(r.Registry, "0x"))...)
	message = append(message, common.HexToAddress(r.Identity).Bytes()...)
	message = append(message, common.HexToAddress(r.Beneficiary).Bytes()...)
	message = append(message, Pad(math.U256(r.Nonce).Bytes(), 32)...)
	return message
}

// RecoverSigner recovers the signer identity from the given request.
func (r SetBeneficiaryRequest) RecoverSigner() (Address, error) {
	signature := r.GetSignatureBytesRaw()

	err := ReformatSignatureVForRecovery(signature)
	if err != nil {
		return Address{}, err
	}

	recoveredAddress, err := RecoverAddress(r.GetMessage(), signature)
	if err != nil {
		return Address{}, err
	}

	return recoveredAddress, nil
}
