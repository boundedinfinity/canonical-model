package geometry

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/angle"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather/trigonometry"
)

func NewCircle(center CartesianCoordinate, radius numberer.Number) Circle {
	return Circle{
		Center: center,
		Radius: radius,
	}
}

func NewCircleXY(x, y, radius numberer.Number) Circle {
	return NewCircle(CartesianCoordinate{X: x, Y: y}, radius)
}

type Circle struct {
	Center CartesianCoordinate
	Radius numberer.Number
}

func (t Circle) Diameter() numberer.Number {
	return numberer.Float(t.Radius.Float() * 2)
}

func (t Circle) PointOnCircumference(a angle.Angle) CartesianCoordinate {
	var theta float64
	radius := t.Radius.Float()

	switch a.Direction {
	case angle.Directions.CounterClockwise:
		theta = 360 - a.Magnitude.Float()
	default:
		theta = a.Magnitude.Float()
	}

	return CartesianCoordinate{
		X: numberer.Float(radius*trigonometry.Cosine(theta) + t.Center.X.Float()),
		Y: numberer.Float(radius*trigonometry.Sine(theta) + t.Center.Y.Float()),
	}
}

// func (t Circle) PointOnDiameter(percent T) CartesianCoordinate {
// 	leftPoint := t.PointOnCircumference(Angle{Magnitude: 180})
// 	x := leftPoint.X.Float() * t.Diameter().Float() * percent
// 	return t.YCoordinate(x)
// }

// (x – h)^2 + (y – k)^2 = r^2
//
//	(y – k)^2 = r^2 - (x – h)^2
//	y – k = √(r^2 - (x – h)^2)
//	y = k + √(r^2 - (x – h)^2)
func (t Circle) YCoordinate(x numberer.Number) CartesianCoordinate {
	h := t.Center.X.Float()
	k := t.Center.Y.Float()
	r := t.Radius.Float()
	return CartesianCoordinate{
		X: x,
		Y: numberer.Float(k + mather.Sqrt(mather.Square(r)-mather.Square(x.Float()-h))),
	}
}

// (x – h)^2 + (y – k)^2 = r^2
//
// (x – h)^2 = r^2 - (y – k)^2
// x - h = √(r^2 - (y – k)^2)
// x = h + √(r^2 - (y – k)^2)
func (t Circle) XCoordinate(y numberer.Number) CartesianCoordinate {
	h := t.Center.X.Float()
	k := t.Center.Y.Float()
	r := t.Radius.Float()
	return CartesianCoordinate{
		X: numberer.Float(h + mather.Sqrt(mather.Square(r)-mather.Square(y.Float()-k))),
		Y: y,
	}
}
