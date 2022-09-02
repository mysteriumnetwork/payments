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

func TestHelpers(t *testing.T) {
	t.Run("is hex number", func(t *testing.T) {
		for _, test := range []struct {
			input  string
			result bool
		}{
			{
				input:  "abcdefg1234",
				result: false,
			},
			{
				input:  "123456",
				result: true,
			},
			{
				input:  "1234567",
				result: false,
			},
			{
				input:  "abcdef",
				result: true,
			},
			{
				input:  "ABCDEF",
				result: true,
			},
			{
				input:  "AbCDeF",
				result: true,
			},
			{
				input:  "12345678abcdE",
				result: false,
			},
		} {
			res := isHex(test.input)
			if test.result != res {
				t.Errorf("isHex(%s) = %t, want %t", test.input, res, test.result)
			}
		}
	})

	t.Run("to bytes", func(t *testing.T) {
		for _, test := range []struct {
			input  string
			result string
		}{
			{
				input:  "0x2345453",
				result: "",
			},
			{
				input:  "0x1be7b0f285d04701f27682f591a60417c47d095a",
				result: strings.ToLower("0000000000000000000000001be7b0f285d04701f27682f591a60417c47d095a"),
			},
			{
				input:  "1be7b0f285d04701f27682f591a60417c47d095a",
				result: strings.ToLower("0000000000000000000000001be7b0f285d04701f27682f591a60417c47d095a"),
			},
			{
				input:  "0x1be7b0f285d04701f27682f591a60417c47d095a1",
				result: "",
			},
		} {
			res, err := toBytes32(test.input)
			if test.result == "" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.result, res)
			}
		}
	})
}
