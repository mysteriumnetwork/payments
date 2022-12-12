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
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// RecoverAddress recovers the address from message and signature
func RecoverAddress(message []byte, signature []byte) (Address, error) {
	publicKey, err := crypto.Ecrecover(crypto.Keccak256(message), signature)
	if err != nil {
		return Address{}, err
	}
	pubKey, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		return Address{}, err
	}
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	return FromCommonAddress(recoveredAddress), nil
}

// GetProxyCode generates bytecode of minimal proxy contract (EIP 1167)
func GetProxyCode(destinationAddress string) ([]byte, error) {
	return hex.DecodeString("3d602d80600a3d3981f3363d3d373d3d3d363d73" + destinationAddress + "5af43d82803e903d91602b57fd5bf3")
}

// keccak("0xff++msg.sender++salt++keccak(byteCode)")
func deriveCreate2Address(salt string, msgSender string, implementation string) (address string, err error) {
	if !common.IsHexAddress(msgSender) || !common.IsHexAddress(implementation) {
		return "", errors.New("msgSender and implementation have to be hex addresses")
	}

	bytecode, _ := GetProxyCode(ensureNoPrefix(implementation))

	input, _ := hex.DecodeString("ff" + ensureNoPrefix(msgSender) + ensureNoPrefix(salt) + common.Bytes2Hex(crypto.Keccak256(bytecode)))
	return "0x" + common.Bytes2Hex(crypto.Keccak256(input))[24:], nil
}

// GenerateChannelAddress generate channel address from given identity hash
func GenerateChannelAddress(identity, hermes, registry, channelImplementation string) (address string, err error) {
	if !common.IsHexAddress(identity) || !common.IsHexAddress(registry) || !common.IsHexAddress(channelImplementation) {
		return "", errors.New("given identity, registry and channelImplementation params have to be hex addresses")
	}

	saltBytes, err := hex.DecodeString(ensureNoPrefix(identity) + ensureNoPrefix(hermes))
	if err != nil {
		return "", err
	}

	salt := hex.EncodeToString(crypto.Keccak256(saltBytes))
	return deriveCreate2Address(salt, registry, channelImplementation)
}

// GenerateProviderChannelID generated channelID for provider channels from given identity hash
func GenerateProviderChannelID(providerIdentity, hermesAddress string) (string, error) {
	return generateProviderChannelID(providerIdentity, hermesAddress, ProviderChannelTypeDefault)
}

// GenerateProviderChannelIDForPayAndSettle generated channelID for provider channels from given identity hash
func GenerateProviderChannelIDForPayAndSettle(providerIdentity, hermesAddress string) (string, error) {
	return generateProviderChannelID(providerIdentity, hermesAddress, ProviderChannelTypePayAndSettle)
}

type ProviderChannelType string

const (
	ProviderChannelTypeDefault      ProviderChannelType = ""
	ProviderChannelTypePayAndSettle ProviderChannelType = "withdrawal"
)

func generateProviderChannelID(providerIdentity, hermesAddress string, channelType ProviderChannelType) (string, error) {
	if !common.IsHexAddress(providerIdentity) || !common.IsHexAddress(hermesAddress) {
		return "", errors.New("given providerIdentity and hermesAddress params have to be hex addresses")
	}

	input := append(
		common.HexToAddress(providerIdentity).Bytes(),
		common.HexToAddress(hermesAddress).Bytes()...,
	)

	if channelType == ProviderChannelTypePayAndSettle {
		input = append(input, []byte(string(ProviderChannelTypePayAndSettle))...)
	}

	return "0x" + common.Bytes2Hex(crypto.Keccak256(input)), nil
}

// GenerateHermesAddress generate hermes address from given hermes operator address
func GenerateHermesAddress(operator string, registry string, hermesImplementation string) (address string, err error) {
	if !common.IsHexAddress(operator) || !common.IsHexAddress(registry) || !common.IsHexAddress(hermesImplementation) {
		return "", errors.New("given operator, registry and hermesImplementation params have to be hex addresses")
	}

	salt, err := toBytes32(operator)
	if err != nil {
		return "", err
	}

	return deriveCreate2Address(salt, registry, hermesImplementation)
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
