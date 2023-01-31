// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package tserr
import (
	"github.com/thorstenrie/tserr" // tserr
)

// Variance returns the variance of discrete value array x as float64. It returns zero and an error
// if x is empty or if the calculation of the arithmetic mean fails.
func Variance[T number](x []T) (float64, error) {
	// Return an error if x is empty
	if len(x) == 0 {
		return float64(0), tserr.Empty("x")
	}
	var (
		mean, vari float64
		err        error
	)
	// Retrieve arithmetic mean of x
	if mean, err = ArithmeticMean(x); err != nil {
		// Return zero and an error if arithmetic mean fails
		return float64(0), tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "x", Err: err})
	}
	// Calculate and return the arithmetic mean
	for i := 0; i < len(x); i++ {
		vari += Square(float64(x[i]) - mean)
	}
	return vari / float64(len(x)), nil
}

// VarianceU returns the variance of the uniform distribution in the interval [a,b] as float64. It returns
// zero and an error, if a is higher than b.
func VarianceU[T number](a, b T) (float64, error) {
	// Return zero and an error if a is higher than b
	if a > b {
		return float64(0), tserr.Higher(&tserr.HigherArgs{Var: "b", Actual: int64(b), LowerBound: int64(a)})
	}
	// Calculate and return the variance
	return float64(Square(b-a)) / float64(12), nil
}

// VarianceN returns the Variance of an n-sided die as float64. It returns zero and an
// error if n is zero or negative.
func VarianceN[T integer](n T) (float64, error) {
	// Return zero and an error if n is zero or negative
	if n <= 0 {
		return float64(0), tserr.Higher(&tserr.HigherArgs{Var: "n", Actual: int64(n), LowerBound: 1})
	}
	// Calculate and return the variance
	return (Square(n) - float64(1)) / float64(12), nil
}
