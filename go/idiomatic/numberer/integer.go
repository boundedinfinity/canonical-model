package numberer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

// GreatestCommonFactor calculates the greatest common factor (GCF), or the largest positive
// integer that evenly divides the input arguments of two non-zero integers.
//
// If either of the inputs are zero, the result is zero.
//
// This is also known as the greatest common divisior/denominator (GCD)
//
// The GCF/GCD is calculated using the Euclidean algorithm.
//
// Reference:
//   - [Greatest common divisor] (Wikipedia)
//   - [EUCLIDEAN ALGORITHM - DISCRETE MATHEMATICS] (YouTube)
//
// [Greatest common divisor]: https://en.wikipedia.org/wiki/Euclidean_algorithm
// [EUCLIDEAN ALGORITHM - DISCRETE MATHEMATICS]: https://www.youtube.com/watch?v=cOwyHTiW4KE
func GreatestCommonFactor[T idiomatic.Integer](a T, b T) T {
	if a == 0 || b == 0 {
		return 0
	}

	return greatestCommonFactor(mather.Abs(a), mather.Abs(b))
}

func greatestCommonFactor[T idiomatic.Integer](a T, b T) T {
	if b == 0 {
		return a
	} else {
		return greatestCommonFactor(b, mather.Mod(a, b))
	}
}

// EnumerateFactors calculates all the factors of the input argument. The return value is a slice
// of all positive integer values that evenly divide into the input argument.
//
// If the input is zero, the result is nil.
func EnumerateFactors[T idiomatic.Integer](n T) []T {
	if n == 0 {
		return nil
	}

	n = mather.Abs(n)
	var factors []T
	start := T(1)

	for i := start; i <= n; i++ {
		mod := n % i

		if mod == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func Magnitude[T idiomatic.Integer](n T) int {
	s := fmt.Sprintf("%v", n)
	return len(s)
}
