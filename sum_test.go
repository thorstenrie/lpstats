// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

import (
	"testing"

	"github.com/thorstenrie/tserr"
)

// TestSumE tests if Sum returns an error for an empty slice
func TestSumE(t *testing.T) {
	if e := testfa(t, nil, 0, Sum[int]); e == nil {
		t.Error(tserr.NilFailed("Sum"))
	}
}

// TestSumi tests the returned value of Sum for a slice of integers.
// It fails if Sum does not return the wanted value.
func TestSumi(t *testing.T) {
	if e := testfa(t, []int{1, 2, 3, 4, 5, 6}, 21, Sum[int]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Sum", Fn: "slice of integers", Err: e}))
	}
}

// TestSumi tests the returned value of Sum for a slice of float64.
// It fails if Sum does not return the wanted value.
func TestSumf(t *testing.T) {
	if e := testfa(t, []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1}, 21.6, Sum[float64]); e != nil {
		t.Error(tserr.Op(&tserr.OpArgs{Op: "Sum", Fn: "slice of integers", Err: e}))
	}
}
