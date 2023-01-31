// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Sum returns the sum of all elements in array x.
func Sum[T number](x []T) T {
	var sum T
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
