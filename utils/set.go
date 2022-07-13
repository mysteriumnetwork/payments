/* Mysterium network Hermes server library.
 *
 * Copyright (C) 2022 BlockDev AG
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

// Set is a collection that Contains no duplicate elements
type Set[t any] struct {
	m map[any]struct{}
}

// newset create new instance of Set
func NewSet[t any]() *Set[t] {
	return &Set[t]{m: make(map[any]struct{})}
}

// Add the specific element to Set if it is not already present
func (s *Set[t]) Add(e any) {
	s.m[e] = struct{}{}
}

// Size returns number of elements in Set
func (s *Set[t]) Size() int {
	return len(s.m)
}

// Iterator returns an Iterator over the elements in this Set.
func (s *Set[t]) Iterator() []t {
	r := []t{}
	for e := range s.m {
		r = append(r, e.(t))
	}

	return r
}

// Contains returns true if this Set Contains the specified element.
func (s *Set[t]) Contains(e t) bool {
	_, exists := s.m[e]

	return exists
}

// IsEmpty returns true is Set Contains no elements
func (s *Set[t]) IsEmpty() bool {
	return len(s.m) == 0
}
