package mather

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic"
)

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[N idiomatic.Number](x N) N {
	return N(math.Abs(float64(x)))
}

func Cbrt[T idiomatic.Number](x T) T {
	return T(math.Cbrt(float64(x)))
}

func Ceil[T idiomatic.Float](x T) T {
	return T(math.Ceil(float64(x)))
}

func Copysign[T idiomatic.Number](f, sign T) T {
	return T(math.Copysign(float64(f), float64(sign)))
}

func Dim[T idiomatic.Number](x, y T) T {
	// Calculate the positive difference
	// [fdim() Function in C - GeeksforGeeks](https://www.geeksforgeeks.org/c/fdim-function-in-c/)
	// [fdim(3) - Linux manual page](https://man7.org/linux/man-pages/man3/fdim.3.html)
	return T(math.Dim(float64(x), float64(y)))
}

// func Erf[T idiomatic.Number](x T) T {
// 	// [Error function - Wikipedia](https://en.wikipedia.org/wiki/Error_function)
// 	return T(math.Erf(float64(x)))
// }

// func Erfc[T idiomatic.Number](x T) T {
// 	// [Error function - Wikipedia : Complementary error function](https://en.wikipedia.org/wiki/Error_function#Complementary_error_function)
// 	return T(math.Erfc(float64(x)))
// }

// func Erfcinv[T idiomatic.Number](x T) T {
// 	// [Error function - Wikipedia](https://en.wikipedia.org/wiki/Error_function)
// 	return T(math.Erfcinv(float64(x)))
// }

// func Erfinv[T idiomatic.Number](x T) T {
// 	// [Error function - Wikipedia](https://en.wikipedia.org/wiki/Error_function)
// 	return T(math.Erfinv(float64(x)))
// }

func Exp[T idiomatic.Number](x T) T {
	return T(math.Exp(float64(x)))
}

func Exp2[T idiomatic.Number](x T) T {
	return T(math.Exp2(float64(x)))
}

// func Expm1[T idiomatic.Number](x T) T {
// 	return T(math.Expm1(float64(x)))
// }

// func FMA[T idiomatic.Number](x, y, z T) T {
// 	return T(math.FMA(float64(x), float64(y), float64(z)))
// }

func Floor[T ~float32 | ~float64](x T) T {
	return T(math.Floor(float64(x)))
}

// func Gamma[T idiomatic.Number](x T) T {
// 	return T(math.Gamma(float64(x)))
// }

func Hypot[T idiomatic.Number](p, q T) T {
	// calculate the hypotenuse with sides p and q.
	return T(math.Hypot(float64(p), float64(q)))
}

func Ilogb[F idiomatic.Number, I idiomatic.Integer](x float64) I {
	return I(math.Ilogb(float64(x)))
}

func Inf[F idiomatic.Number, I idiomatic.Integer](x I) F {
	return F(math.Inf(int(x)))
}

func IsInf[F idiomatic.Number, I idiomatic.Integer](f F, sign I) bool {
	return math.IsInf(float64(f), int(sign))
}

func IsNaN[F idiomatic.Number](x F) bool {
	return math.IsNaN(float64(x))
}

// func J0[T idiomatic.Number](x T) T {
// 	return T(math.J0(float64(x)))
// }

// func J1[T idiomatic.Number](x T) T {
// 	return T(math.J1(float64(x)))
// }

// func Jn[F idiomatic.Float, I idiomatic.Integer](n I, x F) float64 {
// 	return math.Jn(int(n), float64(x))
// }

func Ldexp[F idiomatic.Float, I idiomatic.Integer](n F, x I) float64 {
	return math.Ldexp(float64(n), int(x))
}

func Lgamma(x float64) (lgamma float64, sign int) {
	return math.Lgamma(x)
}

func Log[T idiomatic.Number](x T) T {
	return T(math.Log(float64(x)))
}

func Log10[T idiomatic.Number](x T) T {
	return T(math.Log10(float64(x)))
}

func Log1p[T idiomatic.Number](x T) T {
	return T(math.Log1p(float64(x)))
}

func Log2[T idiomatic.Number](x T) T {
	return T(math.Log2(float64(x)))
}

func Logb[T idiomatic.Number](x T) T {
	return T(math.Logb(float64(x)))
}

func Max[T idiomatic.Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

// func MaxAll[T idiomatic.Numbers](xs ...T) T {
// 	return T(0)
// }

func Min[T idiomatic.Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

// func MinAll[T idiomatic.Numbers](xs ...T) T {
// 	return T(0)
// }

func Mod[T idiomatic.Number](x, y T) T {
	return T(math.Mod(float64(x), float64(y)))
}

func Modf[I idiomatic.Number](x I) (float64, float64) {
	return math.Modf(float64(x))
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter[T idiomatic.Number](x, y T) T {
	return T(math.Nextafter(float64(x), float64(y)))
}

func Nextafter32[T idiomatic.Number](x, y T) T {
	return T(math.Nextafter32(float32(x), float32(y)))
}

func Pow[T idiomatic.Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Square[T idiomatic.Number](n T) T {
	return Pow(n, 2)
}

func Qube[T idiomatic.Number](n T) T {
	return Pow(n, 3)
}

func Pow10[T idiomatic.Integer, U idiomatic.Number](n T) U {
	return U(math.Pow10(int(n)))
}

func Remainder[T idiomatic.Number](x, y T) T {
	return T(math.Remainder(float64(x), float64(y)))
}

func Round[T idiomatic.Number](x T) T {
	return T(math.Round(float64(x)))
}

func RoundToEven[T idiomatic.Number](x T) T {
	return T(math.RoundToEven(float64(x)))
}

func Sqrt[T idiomatic.Number](x T) T {
	return T(math.Sqrt(float64(x)))
}

func Trunc[T idiomatic.Number](x T) T {
	return T(math.Trunc(float64(x)))
}

func Y0[T idiomatic.Number](x T) T {
	return T(math.Y0(float64(x)))
}

// Y1 returns the order-one Bessel function of the second kind.
//
// Special cases are:
//
//	Y1(+Inf) = 0
//	Y1(0) = -Inf
//	Y1(x < 0) = NaN
//	Y1(NaN) = NaN
func Y1[T idiomatic.Number](x T) T {
	return T(math.Y1(float64(x)))
}

// Yn returns the order-n Bessel function of the second kind.
//
// Special cases are:
//
//	Yn(n, +Inf) = 0
//	Yn(n ≥ 0, 0) = -Inf
//	Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
//	Yn(n, x < 0) = NaN
//	Yn(n, NaN) = NaN
func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
