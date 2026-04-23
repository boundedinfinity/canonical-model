package trigonometry

import (
	"math"

	"github.com/boundedinfinity/go-commoner/idiomatic"
)

// https://byjus.com/degree-and-radian-measure-formula/
func DegreeToRadian[T idiomatic.Number](x T) T {
	return T(float64(x) * math.Pi / 180)
}

// https://byjus.com/degree-and-radian-measure-formula/
func RadianToDegree[T idiomatic.Number](x T) T {
	return T(float64(x) * 180 / math.Pi)
}

// https://byjus.com/degree-and-radian-measure-formula/
func ArcLengthRadiusToRadian[T idiomatic.Number](arcLength, radius T) T {
	return arcLength / radius
}

// ArcCosine returns the arccosine, in radians, of x.
//
// Special case is:
//
//	Acos(x) = NaN if x < -1 or x > 1
func ArcCosine[N idiomatic.Number](x N) N {
	return N(math.Acos(float64(x)))
}

// InverseHyperbolicCosine returns the inverse hyperbolic cosine of x.
//
// Special cases are:
//
//	InverseHyperbolicCosine(+Inf) = +Inf
//	InverseHyperbolicCosine(x) = NaN if x < 1
//	InverseHyperbolicCosine(NaN) = NaN
func InverseHyperbolicCosine[N idiomatic.Number](x N) N {
	return N(math.Acosh(float64(x)))
}

// ArcSine returns the arcsine, in radians, of x.
//
// Special cases are:
//
//	ArcSine(±0) = ±0
//	ArcSine(x) = NaN if x < -1 or x > 1
func ArcSine[N idiomatic.Number](x N) N {
	return N(math.Asin(float64(x)))
}

// InverseHyperbolicSine returns the inverse hyperbolic sine of x.
//
// Special cases are:
//
//	InverseHyperbolicSine(±0) = ±0
//	InverseHyperbolicSine(±Inf) = ±Inf
//	InverseHyperbolicSine(NaN) = NaN
func InverseHyperbolicSine[N idiomatic.Number](x N) N {
	return N(math.Asinh(float64(x)))
}

// ArcTangent returns the arctangent, in radians, of x.
//
// Special cases are:
//
//	ArcTangent(±0) = ±0
//	ArcTangent(±Inf) = ±Pi/2
func ArcTangent[N idiomatic.Number](x N) N {
	return N(math.Atan(float64(x)))
}

// ArcTangent2 returns the arc tangent of y/x, using
// the signs of the two to determine the quadrant
// of the return value.
//
// Special cases are (in order):
//
//	ArcTangent2(y, NaN) = NaN
//	ArcTangent2(NaN, x) = NaN
//	ArcTangent2(+0, x>=0) = +0
//	ArcTangent2(-0, x>=0) = -0
//	ArcTangent2(+0, x<=-0) = +Pi
//	ArcTangent2(-0, x<=-0) = -Pi
//	ArcTangent2(y>0, 0) = +Pi/2
//	ArcTangent2(y<0, 0) = -Pi/2
//	ArcTangent2(+Inf, +Inf) = +Pi/4
//	ArcTangent2(-Inf, +Inf) = -Pi/4
//	ArcTangent2(+Inf, -Inf) = 3Pi/4
//	ArcTangent2(-Inf, -Inf) = -3Pi/4
//	ArcTangent2(y, +Inf) = 0
//	ArcTangent2(y>0, -Inf) = +Pi
//	ArcTangent2(y<0, -Inf) = -Pi
//	ArcTangent2(+Inf, x) = +Pi/2
//	ArcTangent2(-Inf, x) = -Pi/2
func ArcTangent2[N idiomatic.Number](y, x N) N {
	return N(math.Atan2(float64(x), float64(y)))
}

// InverseHyperbolicTangent returns the inverse hyperbolic tangent of x.
//
// Special cases are:
//
//	InverseHyperbolicTangent(1) = +Inf
//	InverseHyperbolicTangent(±0) = ±0
//	InverseHyperbolicTangent(-1) = -Inf
//	InverseHyperbolicTangent(x) = NaN if x < -1 or x > 1
//	InverseHyperbolicTangent(NaN) = NaN
func InverseHyperbolicTangent[N idiomatic.Number](x N) N {
	return N(math.Atanh(float64(x)))
}

// Cosine returns the cosine of the radian argument x.
//
// Special cases are:
//
//	Cosine(±Inf) = NaN
//	Cosine(NaN) = NaN
func Cosine[N idiomatic.Number](x N) N {
	return N(math.Cos(float64(x)))
}

// HyperbolicCosine returns the hyperbolic cosine of x.
//
// Special cases are:
//
//	HyperbolicCosine(±0) = 1
//	HyperbolicCosine(±Inf) = +Inf
//	HyperbolicCosine(NaN) = NaN
func HyperbolicCosine[N idiomatic.Number](x N) N {
	return N(math.Cosh(float64(x)))
}

// Sine returns the sine of the radian argument x.
//
// Special cases are:
//
//	Sine(±0) = ±0
//	Sine(±Inf) = NaN
//	Sine(NaN) = NaN
func Sine[N idiomatic.Number](x N) N {
	return N(math.Sin(float64(x)))
}

// Sincos returns Sin(x), Cos(x).
//
// Special cases are:
//
//	Sincos(±0) = ±0, 1
//	Sincos(±Inf) = NaN, NaN
//	Sincos(NaN) = NaN, NaN
func Sincos[N idiomatic.Number](x N) (N, N) {
	sin, cos := math.Sincos(float64(x))
	return N(sin), N(cos)
}

// HyperbolicSine returns the hyperbolic sine of x.
//
// Special cases are:
//
//	HyperbolicSine(±0) = ±0
//	HyperbolicSine(±Inf) = ±Inf
//	HyperbolicSine(NaN) = NaN
func HyperbolicSine[N idiomatic.Number](x N) N {
	return N(math.Sinh(float64(x)))
}

// Tangent returns the tangent of the radian argument x.
//
// Special cases are:
//
//	Tangent(±0) = ±0
//	Tangent(±Inf) = NaN
//	Tangent(NaN) = NaN
func Tangent[N idiomatic.Number](x N) N {
	return N(math.Tan(float64(x)))
}

// HyperbolicTan returns the hyperbolic tangent of x.
//
// Special cases are:
//
//	HyperbolicTan(±0) = ±0
//	HyperbolicTan(±Inf) = ±1
//	HyperbolicTan(NaN) = NaN
func HyperbolicTan[N idiomatic.Number](x N) N {
	return N(math.Tanh(float64(x)))
}
