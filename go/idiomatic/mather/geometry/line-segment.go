package geometry

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

func NewLineSegmentXY(x1, y1, x2, y2 numberer.Number) LineSegment {
	return NewLineSegmentCoords(
		CartesianCoordinate{X: x1, Y: y2},
		CartesianCoordinate{X: x2, Y: y2},
	)
}

func NewLineSegmentCoords(start, end CartesianCoordinate) LineSegment {
	return LineSegment{
		Start: start,
		End:   end,
	}
}

type LineSegment struct {
	Start CartesianCoordinate
	End   CartesianCoordinate
}

func (this LineSegment) Length() numberer.Number {
	xSqrd := mather.Square(this.End.X.Float() - this.Start.X.Float())
	ySqrd := mather.Square(this.End.Y.Float() - this.Start.Y.Float())
	return numberer.Float(mather.Sqrt(xSqrd + ySqrd))
}

func (this LineSegment) PointAtPercent(percentage numberer.Number) CartesianCoordinate {
	p := percentage.Float()
	return CartesianCoordinate{
		X: numberer.Float(this.Start.X.Float() + this.End.X.Float()*p/100),
		Y: numberer.Float(this.Start.Y.Float() + this.End.Y.Float()*p/100),
	}
}

func (this LineSegment) Midpoint() CartesianCoordinate {
	return this.PointAtPercent(numberer.Float(50))
}

func (this LineSegment) Slope() numberer.Number {
	return numberer.Float((this.Start.X.Float() - this.Start.Y.Float()) / (this.End.X.Float() - this.End.Y.Float()))
}

func (this LineSegment) SlopeIntercept() SlopeInterceptLine {
	slope := this.Slope()

	return SlopeInterceptLine{
		M: slope,
		B: numberer.Float(this.Start.Y.Float() / slope.Float() * this.Start.X.Float()),
	}
}

func (this LineSegment) Standard() StandardLine {
	sif := this.SlopeIntercept()

	return StandardLine{
		A: numberer.Float(-sif.M.Float()),
		B: numberer.Float(1),
		C: numberer.Float(-sif.B.Float()),
	}
}

// https://www.cuemath.com/geometry/intersection-of-two-lines/
// https://www.cuemath.com/algebra/cross-multiplication-method/
func (this LineSegment) IntersectAt(segment LineSegment) CartesianCoordinate {
	l1 := this.Standard()
	l2 := segment.Standard()
	l1a := l1.A.Float()
	l1b := l1.B.Float()
	l1c := l1.C.Float()
	l2a := l1.A.Float()
	l2b := l2.B.Float()
	l2c := l2.C.Float()

	return CartesianCoordinate{
		X: numberer.Float((l1b*l2c - l2b*l1c) / (l1a*l2b - l2a*l1b)),
		Y: numberer.Float((l1c*l2a - l2a*l1a) / (l1a*l2b - l2a*l1b)),
	}
}
