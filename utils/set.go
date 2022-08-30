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

// Set is a collection that Contains no duplicate elements.
type Set[T comparable] struct {
	m map[T]struct{}
}

// newset create new instance of Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

// Add the specific element to Set if it is not already present.
func (s *Set[T]) Add(val T) {
	s.m[val] = struct{}{}
}

// Size returns number of elements in Set.
func (s *Set[T]) Size() int {
	return len(s.m)
}

// Iterator returns an Iterator over the elements in this Set.
func (s *Set[T]) Iterator() []T {
	r := []T{}
	for e := range s.m {
		r = append(r, e)
	}
	return r
}

// Contains returns true if this Set Contains the specified element.
func (s *Set[T]) Contains(e T) bool {
	_, exists := s.m[e]

	return exists
}

// IsEmpty returns true is Set Contains no elements.
func (s *Set[T]) IsEmpty() bool {
	return len(s.m) == 0
}

// SliceToSet converts a slice to a Set.
func SliceToSet[T comparable](val []T) Set[T] {
	result := *NewSet[T]()
	for _, v := range val {
		result.Add(v)
	}
	return result
}

// SliceToSetWithTransformation converts a slice to a Set applying a transformation.
func SliceToSetWithTransformation[T comparable, K comparable](val []T, transformFunc func(T) K) Set[K] {
	result := *NewSet[K]()
	for _, v := range val {
		result.Add(transformFunc(v))
	}
	return result
}
