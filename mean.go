// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

import "errors"

func ArithmeticMean[T num](x []T) (float64, error) {
	if len(x) == 0 {
		return float64(0), errors.New("TODO")
	}
	return float64(Sum(x)) / float64(len(x)), nil
}

// Uniform distribution in interval [a,b]
func ExpectedValueU[T num](a, b T) float64 {
	return (float64(a) + float64(b)) / float64(2)
}
