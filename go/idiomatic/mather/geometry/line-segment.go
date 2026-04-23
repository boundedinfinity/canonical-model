package geometry

import (
	"github.com/boundedinfinity/go-commoner/idiomatic"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

func NewLineSegmentXY[T idiomatic.Number](x1, y1, x2, y2 T) LineSegment[T] {
	return NewLineSegmentCoords[T](
		CartesianCoordinate[T]{X: x1, Y: y2},
		CartesianCoordinate[T]{X: x2, Y: y2},
	)
}

func NewLineSegmentCoords[T idiomatic.Number](start, end CartesianCoordinate[T]) LineSegment[T] {
	return LineSegment[T]{
		Start: start,
		End:   end,
	}
}

type LineSegment[T idiomatic.Number] struct {
	Start CartesianCoordinate[T]
	End   CartesianCoordinate[T]
}

func (t LineSegment[T]) Length() T {
	xSqrd := mather.Square(t.End.X - t.Start.X)
	ySqrd := mather.Square(t.End.Y - t.Start.Y)

	return mather.Sqrt(xSqrd + ySqrd)
}

func (t LineSegment[T]) PointAtPercent(percentage T) CartesianCoordinate[T] {
	return CartesianCoordinate[T]{
		X: t.Start.X + t.End.X*percentage/100,
		Y: t.Start.Y + t.End.Y*percentage/100,
	}
}

func (t LineSegment[T]) Midpoint() CartesianCoordinate[T] {
	return t.PointAtPercent(50)
}

func (t LineSegment[T]) Slope() T {
	return (t.Start.X - t.Start.Y) / (t.End.X - t.End.Y)
}

func (t LineSegment[T]) SlopeIntercept() SlopeInterceptLine[T] {
	slope := t.Slope()

	return SlopeInterceptLine[T]{
		M: slope,
		B: t.Start.Y / slope * t.Start.X,
	}
}

func (t LineSegment[T]) Standard() StandardLine[T] {
	sif := t.SlopeIntercept()

	return StandardLine[T]{
		A: -sif.M,
		B: 1,
		C: -sif.B,
	}
}

// https://www.cuemath.com/geometry/intersection-of-two-lines/
// https://www.cuemath.com/algebra/cross-multiplication-method/
func (t LineSegment[T]) IntersectAt(segment LineSegment[T]) CartesianCoordinate[T] {
	l1 := t.Standard()
	l2 := segment.Standard()

	return CartesianCoordinate[T]{
		X: (l1.B*l2.C - l2.B*l1.C) / (l1.A*l2.B - l2.A*l1.B),
		Y: (l1.C*l2.A - l2.A*l1.A) / (l1.A*l2.B - l2.A*l1.B),
	}
}
