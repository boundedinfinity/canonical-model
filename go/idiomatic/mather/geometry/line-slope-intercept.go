package geometry

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
)

type SlopeInterceptLine struct {
	M numberer.Number
	B numberer.Number
}

func (t SlopeInterceptLine) Y(x numberer.Number) numberer.Number {
	return numberer.Float(t.M.Float()*x.Float() + t.B.Float())
}

func (t SlopeInterceptLine) Segment(x1, x2 numberer.Number) LineSegment {
	return NewLineSegmentXY(x1, t.Y(x1), x2, t.Y(x2))
}
