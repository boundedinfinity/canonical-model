package geometry

import "github.com/boundedinfinity/go-commoner/idiomatic"

type SlopeInterceptLine[T idiomatic.Number] struct {
	M T
	B T
}

func (t SlopeInterceptLine[T]) Y(x T) T {
	return t.M*x + t.B
}

func (t SlopeInterceptLine[T]) Segment(x1, x2 T) LineSegment[T] {
	return NewLineSegmentXY(x1, t.Y(x1), x2, t.Y(x2))
}
