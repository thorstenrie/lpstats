// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

import "errors"

func Variance[T num](x []T) (float64, error) {
	if len(x) == 0 {
		return float64(0), errors.New("TODO")
	}
	var (
		mean, vari float64
		err        error
	)
	if mean, err = ArithmeticMean(x); err != nil {
		return float64(0), errors.New("TODO")
	}
	for i := 0; i < len(x); i++ {
		vari += Square(float64(x[i]) - mean)
	}
	return vari / float64(len(x)), nil
}

// Uniform distribution [a,b]
func VarianceU[T num](a, b T) (float64, error) {
	if a > b {
		return float64(0), errors.New("TODO")
	}
	return float64(Square(b-a)) / float64(12), nil
}

// Variance of n-sided die
func VarianceN[T integral](n T) (float64, error) {
	if n <= 0 {
		return float64(0), errors.New("TODO")
	}
	return (Square(n) - float64(1)) / float64(12), nil
}
