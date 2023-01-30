// Copyright (c) 2023 thorstenrie
// All rights reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package lpstats

// num constraints types to number types
type num interface {
	integral | floating
}

// signed constraints types to signed number types
type signed interface {
	floating | sintegral
}

// unsigned constraints types to unsigned number types
type unsigned interface {
	uintegral
}

// integral constraints types to signed and unsigned integer types
type integral interface {
	sintegral | uintegral
}

// sintegral constraints types to signed integer types
type sintegral interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// uintegral constraints types to unsigned integer types
type uintegral interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// floating constraints types to floating types
type floating interface {
	~float32 | ~float64
}
