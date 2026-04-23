package temperature

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
)

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToCelsius[T idiomatic.Number](x T) T {
	return (x - 32) * 5 / 9
}

// https://byjus.com/temperidiomatic.Numbersion-formula/
func CelsiusToFahrenheit[T idiomatic.Number](x T) T {
	return x*5/9 - 32
}

// https://byjus.com/temperature-conversion-formula/
func CelsiusToKelvin[T idiomatic.Number](x T) T {
	return T(float64(x) + 273.15)
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToCelsius[T idiomatic.Number](x T) T {
	return T(float64(x) - 273.15)
}

// https://byjus.com/temperature-conversion-formula/
func FahrenheitToKelvin[T idiomatic.Number](x T) T {
	return T((float64(x)-32)*5/9 + 273.15)
}

// https://byjus.com/temperature-conversion-formula/
func KelvinToFahrenheit[T idiomatic.Number](x T) T {
	return T(((float64(x)-273.15)*5/9 - 32))
}
