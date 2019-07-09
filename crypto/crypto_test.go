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
	channelImplementation := "0x99a73d53959a8fcbe6e67631d39de3cffd3ac9a2"
	expectedChannelAddress := "0x777516c36ceef01ff12f954f7b4a0ea4be4abc0a"

	_, err := GenerateChannelAddress("", "", "")
	assert.EqualError(t, err, "Given identity, registry and channelImplementation params have to be hex addresses")

	channelAddress, err := GenerateChannelAddress(identity, registry, channelImplementation)
	assert.Equal(t, expectedChannelAddress, channelAddress)
	assert.Nil(t, err)
}
