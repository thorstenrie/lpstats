// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package testing
import (
	"testing" // testing

	"github.com/thorstenrie/tserr"
)

// TestExpectedValueUi tests the returned value of ExpectedValueU for integers as arguments.
// It fails if ExpectedValueU does not return the wanted value.
func TestExpectedValueUi(t *testing.T) {
	testf2(t, 10, 20, 15, ExpectedValueU[int])
}

// TestExpectedValueUf tests the returned value of ExpectedValueU for float64 as arguments.
// It fails if ExpectedValueU does not return the wanted value.
func TestExpectedValueUf(t *testing.T) {
	testf2(t, 10.1, 20.1, 15.1, ExpectedValueU[float64])
}

// TestArithmeticMeanE tests if ArithmeticMean returns an error for an empty slice
func TestArithmeticMeanE(t *testing.T) {
	if e := testfa(t, nil, 0, ArithmeticMean[int]); e == nil {
		t.Error(tserr.NilFailed("ArithmeticMean"))
	}
}

// TestArithmeticMeani tests the returned value of ArithmeticMean for a slice of integers.
// It fails if ArithmeticMean does not return the wanted value.
func TestArithmeticMeani(t *testing.T) {
	if e := testfa(t, []int{2500, 2700, 2400, 2300, 2550, 2650, 2750, 2450, 2600, 2400}, 2530, ArithmeticMean[int]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "slice of integers", Err: e}))
	}
}

// TestArithmeticMeanf tests the returned value of ArithmeticMean for a slice of float64.
// It fails if ArithmeticMean does not return the wanted value.
func TestArithmeticMeanf(t *testing.T) {
	if e := testfa(t, []float64{2500, 2700, 2400, 2300, 2550, 2650, 2750, 2450, 2600, 2400}, 2530, ArithmeticMean[float64]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "ArithmeticMean", Fn: "slice of integers", Err: e}))
	}
}
