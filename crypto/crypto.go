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
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// RecoverAddress recovers the address from message and signature
func RecoverAddress(message []byte, signature []byte) (common.Address, error) {
	publicKey, err := crypto.Ecrecover(crypto.Keccak256(message), signature)
	if err != nil {
		return common.Address{}, err
	}
	pubKey, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		return common.Address{}, err
	}
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddress, nil
}

// GetProxyCode generates bytecode of minimal proxy contract (EIP 1167)
func GetProxyCode(destinationAddress string) ([]byte, error) {
	return hex.DecodeString("3d602d80600a3d3981f3363d3d373d3d3d363d73" + destinationAddress + "5af43d82803e903d91602b57fd5bf3")
}

// keccak("0xff++msg.sender++salt++keccak(byteCode)")
func deriveCreate2Address(salt string, msgSender string, implementation string) (address string, err error) {
	if !isHexAddress(msgSender) || !isHexAddress(implementation) {
		return "", errors.New("msgSender and implementation have to be hex addresses")
	}

	bytecode, _ := GetProxyCode(ensureNoPrefix(implementation))

	input, _ := hex.DecodeString("ff" + ensureNoPrefix(msgSender) + salt + common.Bytes2Hex(crypto.Keccak256(bytecode)))
	return "0x" + common.Bytes2Hex(crypto.Keccak256(input))[24:], nil
}

// GenerateChannelAddress generate channel address from given identity hash
func GenerateChannelAddress(identity string, registry string, channelImplementation string) (address string, err error) {
	if !isHexAddress(identity) || !isHexAddress(registry) || !isHexAddress(channelImplementation) {
		return "", errors.New("Given identity, registry and channelImplementation params have to be hex addresses")
	}

	salt, err := toBytes32(identity)
	if err != nil {
		return "", err
	}

	return deriveCreate2Address(salt, registry, channelImplementation)
}

// GenerateProviderChannelID generated channelID for provider channels from given identity hash
func GenerateProviderChannelID(providerIdentity, accountantAddress string) (string, error) {
	if !isHexAddress(providerIdentity) || !isHexAddress(accountantAddress) {
		return "", errors.New("Given providerIdentity and accountantAddress params have to be hex addresses")
	}

	channelID := crypto.Keccak256(append(
		common.HexToAddress(providerIdentity).Bytes(),
		common.HexToAddress(accountantAddress).Bytes()...,
	))

	return "0x" + common.Bytes2Hex(channelID), nil
}

// GenerateProviderChannelIDBytes received provider and accountnat Address
func GenerateProviderChannelIDBytes(providerIdentity, accountantAddress common.Address) []byte {
	return crypto.Keccak256(append(
		providerIdentity.Bytes(),
		accountantAddress.Bytes()...,
	))
}

// GenerateAccountantAddress generate accountant address from given accountant operator address
func GenerateAccountantAddress(operator string, registry string, accountantImplementation string) (address string, err error) {
	if !isHexAddress(operator) || !isHexAddress(registry) || !isHexAddress(accountantImplementation) {
		return "", errors.New("Given operator, registry and accountantImplementation params have to be hex addresses")
	}

	salt, err := toBytes32(operator)
	if err != nil {
		return "", err
	}

	return deriveCreate2Address(salt, registry, accountantImplementation)
}

// ReformatSignatureVForBC takes in the signature and modifies its last byte to correspond to the format required for SC
func ReformatSignatureVForBC(signature []byte) error {
	if len(signature) != 65 {
		return errors.New("the signature must be 65 bytes long")
	}

	var v = 27 + (uint64(signature[len(signature)-1]) % 2)
	buffer := make([]byte, 1)
	_ = binary.PutUvarint(buffer, v)
	signature[64] = buffer[0]

	return nil
}

// ReformatSignatureVForRecovery takes in  the signature and modifies its last byte to normalize V to either 0 or 1
func ReformatSignatureVForRecovery(signature []byte) error {
	if len(signature) != 65 {
		return errors.New("the signature must be 65 bytes long")
	}

	v := uint64(signature[64]) % 27
	buffer := make([]byte, 1)
	_ = binary.PutUvarint(buffer, v)
	signature[64] = buffer[0]

	return nil
}
