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
	"strings"
	"testing"

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
	hermesAddress := "0x676b9a084aC11CEeF680AF6FFbE99b24106F47e7"
	channelImplementation := "0x99a73d53959a8fcbe6e67631d39de3cffd3ac9a2"
	expectedChannelAddress := "0x75bc5ea5f48949032278179132d367f06ab7570e"

	_, err := GenerateChannelAddress("", "", "", "")
	assert.EqualError(t, err, "given identity, registry and channelImplementation params have to be hex addresses")

	channelAddress, err := GenerateChannelAddress(identity, hermesAddress, registry, channelImplementation)
	assert.Nil(t, err)
	assert.Equal(t, expectedChannelAddress, channelAddress)
}

func TestGenerateProviderChannelID(t *testing.T) {
	providerIdentity := "0x761f2bb3e7AD6385a4c7833c5a26a8Ddfdabf9f3"
	hermesAddress := "0x676b9a084aC11CEeF680AF6FFbE99b24106F47e7"
	expectedChannelID := "0xfebeba54c56610475ae3432199515146096e2d9c5b7bc2b3865c4b1967cf01cc"

	_, err := GenerateProviderChannelID("Identity", "Hermes")
	assert.EqualError(t, err, "given providerIdentity and hermesAddress params have to be hex addresses")

	channelID, err := GenerateProviderChannelID(providerIdentity, hermesAddress)
	assert.NoError(t, err)
	assert.Equal(t, expectedChannelID, channelID)
}

func TestGenerateHermesAddress(t *testing.T) {
	operator := "0x76259c949bee90c3ef6f1d04a3cb50ed0de7763c"
	registry := "0x1ba2df26371e83d87afee2f27a42f5a7fe9e5219"
	hermesImplementation := "0xac69e0c98a688e35698630eb0c741eb2a2fc5ef1"
	expectedHermesAddress := strings.ToLower("0xcAeF9A6E9C2d9C0Ee3333529922c280580365b51")

	_, err := GenerateHermesAddress("", "", "")
	assert.EqualError(t, err, "given operator, registry and hermesImplementation params have to be hex addresses")

	hermesAddress, err := GenerateHermesAddress(operator, registry, hermesImplementation)
	assert.Nil(t, err)
	assert.Equal(t, expectedHermesAddress, hermesAddress)
}
