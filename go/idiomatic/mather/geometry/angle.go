package geometry

import (
	gomath "math"

	"github.com/boundedinfinity/go-commoner/idiomatic"
)

type AngleNumber interface {
	~int | ~int16 | ~int32 | ~int64 | ~uint | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

type Angle[T AngleNumber] struct {
	Magnitude T
	Type      AngleType
	Direction AngleDirection
}

func (t Angle[T]) ToRadian() Angle[T] {
	switch t.Type {
	case AngleTypes.Radians:
		return t
	default:
		return Angle[T]{
			Type:      AngleTypes.Radians,
			Direction: t.Direction,
			Magnitude: DegreeToRadian(t.Magnitude),
		}
	}
}

func (t Angle[T]) ToDegree() Angle[T] {
	switch t.Type {
	case AngleTypes.Degrees:
		return t
	default:
		return Angle[T]{
			Type:      AngleTypes.Degrees,
			Direction: t.Direction,
			Magnitude: RadianToDegree(t.Magnitude),
		}
	}
}

func DegreeToRadian[T idiomatic.Number](angle T) T {
	return T(float64(angle) * gomath.Pi / 180)
}

func RadianToDegree[T idiomatic.Number](angle T) T {
	return T(float64(angle) * 180 / gomath.Pi)
}
