// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Max returns the maximum of a and b. If a is higher than b, it returns a.
// Otherwise, it returns b.
func Max[T Number](a, b T) T {
	// Return a, if a is higher than b
	if a > b {
		return a
	}
	// Otherwise, return b
	return b
}

// Min returns the minimum of a and b. If a is lower than b, it returns a.
// Otherwise, it returns b.
func Min[T Number](a, b T) T {
	// Returns a, if a is lower than b
	if a < b {
		return a
	}
	// Otherwise, return b
	return b
}
