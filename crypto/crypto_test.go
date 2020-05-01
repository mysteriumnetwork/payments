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
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/assert"
)

func TestDeriveCreate2Address(t *testing.T) {
	salt := "000000000000000000000000265B4A774A5CE7A975CA8401A43440EFEE58EB15"
	registry := "0x6bb8345c9d996be4fab652f4a15813303d630b66"
	implementation := "0x99a73d53959a8fcbe6e67631d39de3cffd3ac9a2"
	expectedChannelAddress := "0x777516c36ceef01ff12f954f7b4a0ea4be4abc0a"

	_, err := deriveCreate2Address("", "", "")
	assert.EqualError(t, err, "msgSender and implementation have to be hex addresses")

	channelAddress, err := deriveCreate2Address(salt, registry, implementation)
	assert.Equal(t, expectedChannelAddress, channelAddress)
	assert.Nil(t, err)
}

func TestGenerateChannelAddress(t *testing.T) {
	identity := "0x265B4A774A5CE7A975CA8401A43440EFEE58EB15"
	registry := "0x6bb8345c9d996be4fab652f4a15813303d630b66"
	accountantAddress := "0x676b9a084aC11CEeF680AF6FFbE99b24106F47e7"
	channelImplementation := "0x99a73d53959a8fcbe6e67631d39de3cffd3ac9a2"
	expectedChannelAddress := "0x75bc5ea5f48949032278179132d367f06ab7570e"

	_, err := GenerateChannelAddress("", "", "", "")
	assert.EqualError(t, err, "Given identity, registry and channelImplementation params have to be hex addresses")

	channelAddress, err := GenerateChannelAddress(identity, accountantAddress, registry, channelImplementation)
	assert.Nil(t, err)
	assert.Equal(t, expectedChannelAddress, channelAddress)
}

func TestGenerateProviderChannelID(t *testing.T) {
	providerIdentity := "0x761f2bb3e7AD6385a4c7833c5a26a8Ddfdabf9f3"
	accountantAddress := "0x676b9a084aC11CEeF680AF6FFbE99b24106F47e7"
	expectedChannelID := "0xfebeba54c56610475ae3432199515146096e2d9c5b7bc2b3865c4b1967cf01cc"

	_, err := GenerateProviderChannelID("Identity", "Accountant")
	assert.EqualError(t, err, "Given providerIdentity and accountantAddress params have to be hex addresses")

	channelID, err := GenerateProviderChannelID(providerIdentity, accountantAddress)
	assert.NoError(t, err)
	assert.Equal(t, expectedChannelID, channelID)
}

func TestGenerateProviderChannelIDBytes(t *testing.T) {
	providerIdentity := common.HexToAddress("0x761f2bb3e7AD6385a4c7833c5a26a8Ddfdabf9f3")
	accountantAddress := common.HexToAddress("0x676b9a084aC11CEeF680AF6FFbE99b24106F47e7")
	expectedChannelID := []byte{0xfe, 0xbe, 0xba, 0x54, 0xc5, 0x66, 0x10, 0x47, 0x5a, 0xe3, 0x43, 0x21, 0x99, 0x51, 0x51, 0x46, 0x9, 0x6e, 0x2d, 0x9c, 0x5b, 0x7b, 0xc2, 0xb3, 0x86, 0x5c, 0x4b, 0x19, 0x67, 0xcf, 0x1, 0xcc}

	channelID := GenerateProviderChannelIDBytes(providerIdentity, accountantAddress)
	assert.Equal(t, expectedChannelID, channelID)
}
