// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

import (
	"testing"

	"github.com/thorstenrie/tserr"
)

// TestNEf tests NearEqual to return false for a difference of a and b greater than the
// provided maximum difference. It fails if NearEqual returns true.
func TestNEf(t *testing.T) {
	if NearEqual(0.01, 0.03, -0.01) {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "NearEqual", Actual: "true", Want: "false"}))
	}
}

// TestNEe tests NearEqual to return true for a difference of a and b equal to the
// provided maximum difference. It fails if NearEqual returns false.
func TestNEe(t *testing.T) {
	if !NearEqual(0.01, 0.02, -0.01) {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "NearEqual", Actual: "false", Want: "true"}))
	}
}

// TestNEt tests NearEqual to return true for a difference of a and b lower than the
// provided maximum difference. It fails if NearEqual returns false.
func TestNEt(t *testing.T) {
	if !NearEqual(0.01, 0.02, -0.02) {
		t.Error(tserr.Return(&tserr.ReturnArgs{Op: "NearEqual", Actual: "false", Want: "true"}))
	}
}
