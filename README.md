# lpstats
Go package for simple statistics

[![Go Report Card](https://goreportcard.com/badge/github.com/thorstenrie/lpstats)](https://goreportcard.com/report/github.com/thorstenrie/lpstats)
[![CodeFactor](https://www.codefactor.io/repository/github/thorstenrie/lpstats/badge)](https://www.codefactor.io/repository/github/thorstenrie/lpstats)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorstenrie/lpstats)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorstenrie/lpstats)](https://pkg.go.dev/mod/github.com/thorstenrie/lpstats)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorstenrie/lpstats)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorstenrie/lpstats)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorstenrie/lpstats)
![GitHub last commit](https://img.shields.io/github/last-commit/thorstenrie/lpstats)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorstenrie/lpstats)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorstenrie/lpstats)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorstenrie/lpstats)
![GitHub](https://img.shields.io/github/license/thorstenrie/lpstats)

The package lpstats provides a simple interface for statistics. It provides functions for the calculation of mean and variance.

- **Simple**: Without configuration, just function calls
- **Easy to use**: Most functions take integer and float values as arguments
- **Tested**: Unit tests with high [code coverage](https://gocover.io/github.com/thorstenrie/lpstats)
- **Dependencies**: Only depends on the [Go Standard Library](https://pkg.go.dev/std) and [tserr](https://github.com/thorstenrie/tserr)

## Usage

The package is installed with 

```
go get github.com/thorstenrie/lpstats
```

In the Go app, the package is imported with

```
import "github.com/thorstenrie/lpstats"
```

External functions can be used with arguments of type integer or float. Type constraints are defined in [constraints.go](https://github.com/thorstenrie/lpstats/blob/main/constraints.go).

- `Number`: number types, for example, `func Square[T Number](x T) float64`
- `Signed`: signed number types, for example, `func Sign[T Signed](a T) T`
- `Sinteger`: signed integer types, for example `func EqualS[T Sinteger](x, y T) error`
- `Uinteger`: unsigned integer types, for example, `func VarianceN[T Uinteger](n T) float64`

## Functions

The package provides mathematical and statistical functions

- Absolute value
- Arithmetic mean of a discrete value array
- Equal for slices of signed integers
- Expected value for a uniform distribution
- Max
- Min
- Near equal
- Sign
- Square
- Sum
- Variance of a discrete value array
- Variance of a uniform distribution

## Example

```
package main

import (
	"fmt"

	"github.com/thorstenrie/lpstats"
)

func main() {
	var (
		i []int     = []int{1, 2, 3, 4, 5, 6}
		x []float64 = []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1}
		n uint      = 6
	)

	im, _ := lpstats.ArithmeticMean(i)
	xm, _ := lpstats.ArithmeticMean(x)

	iv, _ := lpstats.Variance(i)
	xv, _ := lpstats.Variance(x)

	fmt.Printf("Mean(i) = %f, Variance(i) = %f\n", im, iv)
	fmt.Printf("Mean(f) = %f, Variance(x) = %f\n", xm, xv)

	fmt.Printf("Variance of a %d-sided die: %.2f\n", n, lpstats.VarianceN(n))
}
```

## Known Limitations

- Most calculations are based on floating point arithmetic. It is not suitable for arbitrary precision fixed-point decimal arithmetic.

## Links

[Godoc](https://pkg.go.dev/github.com/thorstenrie/lpstats)

[Go Report Card](https://goreportcard.com/report/github.com/thorstenrie/lpstats)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorstenrie%2Flpstats)
