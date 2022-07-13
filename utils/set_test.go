/* Mysterium network Hermes server library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSet_AddGet(t *testing.T) {
	s := NewSet[common.Address]()

	a1 := common.HexToAddress("1")
	a2 := common.HexToAddress("2")
	s.Add(a1)
	s.Add(a2)
	assert.Equal(t, 2, s.Size())
}

func TestSet_Iterator(t *testing.T) {
	s := NewSet[common.Address]()
	assert.Equal(t, 0, s.Size())
	a1 := common.HexToAddress("1")
	a2 := common.HexToAddress("2")
	a3 := common.HexToAddress("3")
	s.Add(a1)
	s.Add(a2)
	s.Add(a3)
	i := s.Iterator()
	assert.Equal(t, 3, len(i))
}

func TestSet_Contains(t *testing.T) {
	s := NewSet[common.Address]()
	a1 := common.HexToAddress("1")
	a2 := common.HexToAddress("2")
	s.Add(a1)
	assert.True(t, s.Contains(a1))
	assert.False(t, s.Contains(a2))
}

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet[common.Address]()
	assert.True(t, s.IsEmpty())
}
