// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// NearEqual returns whether a and b near equal with the maximum absolute difference maxDiff.
// It returns true if the absolute value of the difference of a and b equals or is lower than maxDiff.
// It returns false if the absolute value of the difference of a and b is higher than maxDiff.
func NearEqual[T Floating](a, b, maxDiff T) bool {
	return Abs(a-b) <= Abs(maxDiff)
}
