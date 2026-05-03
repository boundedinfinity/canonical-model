package geometry

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/angle"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/trigonometry"
)

type CartesianCoordinate struct {
	X numberer.Number
	Y numberer.Number
}

type PolarCoordinate struct {
	Radius numberer.Number
	Angle  angle.Angle
}

func CartesianToPolar(coordinate CartesianCoordinate) PolarCoordinate {
	x := coordinate.X.Float()
	y := coordinate.Y.Float()
	return PolarCoordinate{
		Radius: numberer.Float(mather.Sqrt(mather.Pow(x, 2) + mather.Pow(y, 2))),
		Angle: angle.Angle{
			Magnitude: numberer.Float(trigonometry.ArcTangent(y / x)),
			Kind:      angle.Kinds.Radian,
			Direction: angle.Directions.CounterClockwise,
		},
	}
}

func PolarToCartesian(coordinate PolarCoordinate) CartesianCoordinate {
	angle := coordinate.Angle.ToRadian()
	radius := coordinate.Radius.Float()
	magnitude := angle.Magnitude.Float()

	return CartesianCoordinate{
		X: numberer.Float(radius * trigonometry.Cosine(magnitude)),
		Y: numberer.Float(radius * trigonometry.Sine(magnitude)),
	}
}
