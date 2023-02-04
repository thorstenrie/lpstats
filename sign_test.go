// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package testing
import (
	"testing" // testing
)

// TestSignfp tests the result of Sign for a positive float64. If the result of Sign does not
// equal +1, the test fails.
func TestSignfp(t *testing.T) {
	testf(t, 10 /*a*/, 1 /*w*/, Sign[float64])
}

// TestSignfz tests the result of Sign for the zero float64. If the result of Sign does not
// equal 0, the test fails.
func TestSignfz(t *testing.T) {
	testf(t, 10 /*a*/, 1 /*w*/, Sign[float64])
}

// TestSignfn tests the result of Sign for a negative float64. If the result of Sign does not
// equal -1, the test fails.
func TestSignfn(t *testing.T) {
	testf(t, -10 /*a*/, -1 /*w*/, Sign[float64])
}

// TestSignip tests the result of Sign for a positive integer. If the result of Sign does not
// equal +1, the test fails.
func TestSignip(t *testing.T) {
	testi(t, 10 /*a*/, 1 /*w*/, Sign[int64])
}

// TestSigniz tests the result of Sign for the zero integer. If the result of Sign does not
// equal 0, the test fails.
func TestSigniz(t *testing.T) {
	testi(t, 0 /*a*/, 0 /*w*/, Sign[int64])
}

// TestSignin tests the result of Sign for a negative integer. If the result of Sign does not
// equal -1, the test fails.
func TestSignin(t *testing.T) {
	testi(t, -10 /*a*/, -1 /*w*/, Sign[int64])
}
