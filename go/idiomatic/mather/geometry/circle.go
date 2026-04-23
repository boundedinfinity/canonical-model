package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/trigonometry"
)

func NewCircle[T AngleNumber](center CartesianCoordinate[T], radius T) Circle[T] {
	return Circle[T]{
		Center: center,
		Radius: radius,
	}
}

func NewCircleXY[T AngleNumber](x, y, radius T) Circle[T] {
	return NewCircle[T](CartesianCoordinate[T]{X: x, Y: y}, radius)
}

type Circle[T AngleNumber] struct {
	Center CartesianCoordinate[T]
	Radius T
}

func (t Circle[T]) Diameter() T {
	return t.Radius * T(2)
}

func (t Circle[T]) PointOnCircumference(angle Angle[T]) CartesianCoordinate[T] {
	var theta T

	switch angle.Direction {
	case AngleDirections.CounterClockwise:
		theta = T(360) - theta
	default:
		theta = angle.Magnitude
	}

	return CartesianCoordinate[T]{
		X: t.Radius*trigonometry.Cosine(theta) + t.Center.X,
		Y: t.Radius*trigonometry.Sine(theta) + t.Center.Y,
	}
}

func (t Circle[T]) PointOnDiameter(percent T) CartesianCoordinate[T] {
	leftPoint := t.PointOnCircumference(Angle[T]{Magnitude: 180})
	x := leftPoint.X * t.Diameter() * percent
	return t.YCoordinate(x)
}

// (x – h)^2 + (y – k)^2 = r^2
//
//	(y – k)^2 = r^2 - (x – h)^2
//	y – k = √(r^2 - (x – h)^2)
//	y = k + √(r^2 - (x – h)^2)
func (t Circle[T]) YCoordinate(x T) CartesianCoordinate[T] {
	h := t.Center.X
	k := t.Center.Y
	r := t.Radius
	return CartesianCoordinate[T]{
		X: x,
		Y: k + mather.Sqrt(mather.Square(r)-mather.Square(x-h)),
	}
}

// (x – h)^2 + (y – k)^2 = r^2
//
// (x – h)^2 = r^2 - (y – k)^2
// x - h = √(r^2 - (y – k)^2)
// x = h + √(r^2 - (y – k)^2)
func (t Circle[T]) XCoordinate(y T) CartesianCoordinate[T] {
	h := t.Center.X
	k := t.Center.Y
	r := t.Radius
	return CartesianCoordinate[T]{
		X: h + mather.Sqrt(mather.Square(r)-mather.Square(y-k)),
		Y: y,
	}
}
