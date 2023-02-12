// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// number constraints types to all number types
type Number interface {
	Integer | Floating
}

// signed constraints types to signed number types
type Signed interface {
	Floating | Sinteger
}

// integer constraints types to signed and unsigned integer types
type Integer interface {
	Sinteger | Uinteger
}

// sinteger constraints types to signed integer types
type Sinteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// uinteger constraints types to unsigned integer types
type Uinteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// floating constraints types to floating types
type Floating interface {
	~float32 | ~float64
}
