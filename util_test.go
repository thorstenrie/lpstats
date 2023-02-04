// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

import (
	"math"
	"testing"

	"github.com/thorstenrie/tserr"
)

// testi tests function f to return wanted integer w for argument a. If the
// result of f for a does not equal w, the test fails.
func testi[T number](t *testing.T, a T, w int64, f func(T) int64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a in r
	r := f(a)
	// If the result r does not equal the wanted result w, the test fails
	if r != w {
		t.Error(tserr.Equal(&tserr.EqualArgs{Var: "function of a", Actual: r, Want: w}))
	}
}

// testf tests function f to return wanted float64 w for argument a. If the
// result of f for a does not equal w with the precision of two decimal places, the test fails.
func testf[T number](t *testing.T, a T, w float64, f func(T) float64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a in r
	r := f(a)
	// If the result r does not equal the wanted result w with the precision of two decimal places, the test fails
	if !testEqualp2(r, w) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "function of a", Actual: r, Want: w}))
	}
}

// testf2 tests function f to return wanted float64 w for arguments a and b. If the
// result of f for a and b does not equal w with the precision of two decimal places, the test fails.
func testf2[T number](t *testing.T, a, b T, w float64, f func(T, T) float64) {
	if t == nil {
		panic("nil pointer")
	}
	// Retrieve f of a in b
	r := f(a, b)
	// If the result r does not equal the wanted result w with the precision of two decimal places, the test fails
	if !testEqualp2(r, w) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "function of a", Actual: r, Want: w}))
	}
}

// testfa tests function f to return wanted float64 w for slice a. If the
// result of f for slice a does not equal w with the precision of two decimal places, the test fails.
// It returns the error of f, if any
func testfa[T number](t *testing.T, a []T, w float64, f func([]T) (float64, error)) error {
	if t == nil {
		panic("nil pointer")
	}
	r, e := f(a)
	// If the result b does not equal the wanted result w with the precision of two decimal places, the test fails
	if !testEqualp2(r, w) {
		t.Error(tserr.Equalf(&tserr.EqualfArgs{Var: "function of a", Actual: r, Want: w}))
	}
	return e
}

// testEqualp2 returns whether a and b equal with the precision of two decimal places.
// It returns true, if a and b are equal with the precision of two decimals places.
// It returns false otherwise.
func testEqualp2(a float64, b float64) bool {
	var p float64 = 100
	return math.Round(a*p) == math.Round(b*p)
}
