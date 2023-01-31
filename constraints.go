// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// number constraints types to all number types
type number interface {
	integer | floating
}

// signed constraints types to signed number types
type signed interface {
	floating | sinteger
}

// integer constraints types to signed and unsigned integer types
type integer interface {
	sinteger | uinteger
}

// sinteger constraints types to signed integer types
type sinteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// uinteger constraints types to unsigned integer types
type uinteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// floating constraints types to floating types
type floating interface {
	~float32 | ~float64
}
