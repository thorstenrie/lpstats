// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// Import package testing and tserr
import (
	"errors"
	"testing" // testing
)

func TestSignFP(t *testing.T) {
	var a float64 = 1
	b := Sign(a)
	if b != float64(1) {
		t.Error(errors.New("error"))
	}
}
