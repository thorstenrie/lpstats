// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

func Square[T num](x T) float64 {
	y := float64(x)
	return y * y
}
