package geometry

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
)

func NewRectangle(topLeft CartesianCoordinate, dimensions Dimension2d) Rectangle {
	return Rectangle{
		TopLeft:    topLeft,
		Dimensions: dimensions,
	}
}

func NewRectangleXYHW(x, y, height, width numberer.Number) Rectangle {
	return NewRectangle(
		CartesianCoordinate{X: x, Y: y},
		Dimension2d{Height: height, Width: width},
	)
}

type Rectangle struct {
	TopLeft    CartesianCoordinate
	Dimensions Dimension2d
}

func (t Rectangle) Area() numberer.Number {
	return numberer.Float(t.Dimensions.Height.Float() * t.Dimensions.Width.Float())
}

func (t Rectangle) Perimeter() numberer.Number {
	return numberer.Float(2 * (t.Dimensions.Height.Float() + t.Dimensions.Width.Float()))
}

func (t Rectangle) BottomLeft() CartesianCoordinate {
	return CartesianCoordinate{
		X: t.TopLeft.X,
		Y: numberer.Float(t.TopLeft.Y.Float() + t.Dimensions.Height.Float()),
	}
}

func (t Rectangle) TopRight() CartesianCoordinate {
	return CartesianCoordinate{
		X: numberer.Float(t.TopLeft.X.Float() + t.Dimensions.Width.Float()),
		Y: t.TopLeft.Y,
	}
}

func (t Rectangle) BottomRight() CartesianCoordinate {
	return CartesianCoordinate{
		X: numberer.Float(t.TopLeft.X.Float() + t.Dimensions.Width.Float()),
		Y: numberer.Float(t.TopLeft.Y.Float() + t.Dimensions.Height.Float()),
	}
}

func (t Rectangle) TopMidpoint() CartesianCoordinate {
	return NewLineSegmentCoords(t.TopLeft, t.BottomRight()).Midpoint()
}

func (t Rectangle) RightMidpoint() CartesianCoordinate {
	return NewLineSegmentCoords(t.TopRight(), t.BottomRight()).Midpoint()
}

func (t Rectangle) BottomMidpoint() CartesianCoordinate {
	return NewLineSegmentCoords(t.BottomLeft(), t.BottomRight()).Midpoint()
}

func (t Rectangle) LeftMidpoint() CartesianCoordinate {
	return NewLineSegmentCoords(t.TopLeft, t.BottomLeft()).Midpoint()
}

func (t Rectangle) Center() CartesianCoordinate {
	return NewLineSegmentCoords(t.TopMidpoint(), t.BottomMidpoint()).Midpoint()
}

// func (t Rectangle) PointOnPeremiter(angle Angle[A]) CartesianCoordinate {
// 	center := t.Center()
// 	centerToTopLeft := NewLineSegmentCoords(center, t.TopLeft)
// 	circle := NewCircle(center, centerToTopLeft.Length())
// 	circlePoint := circle.PointOnCircumference[A](angle)
// 	circleLine := NewLineSegmentCoords(center, circlePoint)
// 	topLine := NewLineSegmentCoords(t.TopLeft, t.TopRight())

// 	return topLine.IntersectAt(circleLine)
// }
