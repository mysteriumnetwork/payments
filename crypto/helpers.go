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
	"errors"
	"strings"
)

func hasHexPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}

	return true
}

func isHexAddress(s string) bool {
	if hasHexPrefix(s) {
		s = s[2:]
	}

	return len(s) == 40 && isHex(s)
}

func toBytes32(address string) (string, error) {
	if !isHexAddress(address) {
		return "", errors.New("Given string is not a hex address")
	}

	if hasHexPrefix(address) {
		address = address[2:]
	}

	res := strings.ToLower("000000000000000000000000" + address)

	return res, nil
}

func ensureNoPrefix(address string) string {
	return strings.TrimPrefix(address, "0x")
}

// Pad pads the given byte array to given size by prefixing with zeros
func Pad(b []byte, size int) []byte {
	if len(b) >= size {
		return b
	}
	tmp := make([]byte, size)
	copy(tmp[size-len(b):], b)
	return tmp
}
