package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"
)

func NewRectangle[T idiomatic.Number, A AngleNumber](topLeft CartesianCoordinate[T], dimensions Dimension2d[T]) Rectangle[T, A] {
	return Rectangle[T, A]{
		TopLeft:    topLeft,
		Dimensions: dimensions,
	}
}

func NewRectangleXYHW[T idiomatic.Number, A AngleNumber](x, y, height, width T) Rectangle[T, A] {
	return NewRectangle[T, A](
		CartesianCoordinate[T]{X: x, Y: y},
		Dimension2d[T]{Height: height, Width: width},
	)
}

type Rectangle[T idiomatic.Number, A AngleNumber] struct {
	TopLeft    CartesianCoordinate[T]
	Dimensions Dimension2d[T]
}

func (t Rectangle[T, A]) Area() T {
	return t.Dimensions.Height * t.Dimensions.Width
}

func (t Rectangle[T, A]) Perimeter() T {
	return 2 * (t.Dimensions.Height + t.Dimensions.Width)
}

func (t Rectangle[T, A]) BottomLeft() CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.TopLeft.X,
		Y: t.TopLeft.Y + t.Dimensions.Height,
	}
}

func (t Rectangle[T, A]) TopRight() CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.TopLeft.X + t.Dimensions.Width,
		Y: t.TopLeft.Y,
	}
}

func (t Rectangle[T, A]) BottomRight() CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.TopLeft.X + t.Dimensions.Width,
		Y: t.TopLeft.Y + t.Dimensions.Height,
	}
}

func (t Rectangle[T, A]) TopMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopLeft, t.BottomRight()).Midpoint()
}

func (t Rectangle[T, A]) RightMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopRight(), t.BottomRight()).Midpoint()
}

func (t Rectangle[T, A]) BottomMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.BottomLeft(), t.BottomRight()).Midpoint()
}

func (t Rectangle[T, A]) LeftMidpoint() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopLeft, t.BottomLeft()).Midpoint()
}

func (t Rectangle[T, A]) Center() CartesianCoordinate[T] {
	return NewLineSegmentCoords(t.TopMidpoint(), t.BottomMidpoint()).Midpoint()
}

// func (t Rectangle[T, A]) PointOnPeremiter(angle Angle[A]) CartesianCoordinate[T] {
// 	center := t.Center()
// 	centerToTopLeft := NewLineSegmentCoords(center, t.TopLeft)
// 	circle := NewCircle[T, A](center, centerToTopLeft.Length())
// 	circlePoint := circle.PointOnCircumference[A](angle)
// 	circleLine := NewLineSegmentCoords(center, circlePoint)
// 	topLine := NewLineSegmentCoords(t.TopLeft, t.TopRight())

// 	return topLine.IntersectAt(circleLine)
// }
