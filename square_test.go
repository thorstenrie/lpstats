// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package testing
import (
	"testing" // testing
)

// TestSquarefp tests the result of Square for a positive float64. If the result of Square does not
// equal the square value, the test fails.
func TestSquarefp(t *testing.T) {
	testf(t, 5 /*a*/, 25 /*w*/, Square[float64])
}

// TestSquarefz tests the result of Square for the zero float64. If the result of Square does not
// equal 0, the test fails.
func TestSquarefz(t *testing.T) {
	testf(t, 0 /*a*/, 0 /*w*/, Square[float64])
}

// TestSquarefn tests the result of Square for a negative float64. If the result of Square does not
// equal the square value, the test fails.
func TestSquarefn(t *testing.T) {
	testf(t, -5 /*a*/, 25 /*w*/, Square[float64])
}

// TestSquareip tests the result of Square for a positive integer. If the result of Square does not
// equal the square value, the test fails.
func TestSquareip(t *testing.T) {
	testf(t, 5 /*a*/, 25 /*w*/, Square[int64])
}

// TestSquareiz tests the result of Square for the zero integer. If the result of Square does not
// equal 0, the test fails.
func TestSquareiz(t *testing.T) {
	testf(t, 0 /*a*/, 0 /*w*/, Square[int64])
}

// TestSquarein tests the result of Square for a negative integer. If the result of Square does not
// equal the square value, the test fails.
func TestSquarein(t *testing.T) {
	testf(t, -5 /*a*/, 25 /*w*/, Square[int64])
}
