package angle

import (
	"math"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/numberer"
)

type Angle struct {
	Kind      Kind
	Direction Direction
	Magnitude numberer.Number
}

func (t Angle) ToRadian() Angle {
	switch t.Kind {
	case Kinds.Radian:
		return t
	default:
		return Angle{
			Kind:      Kinds.Radian,
			Direction: t.Direction,
			Magnitude: DegreeToRadian(t.Magnitude),
		}
	}
}

func (t Angle) ToDegree() Angle {
	switch t.Kind {
	case Kinds.Degree:
		return t
	default:
		return Angle{
			Kind:      Kinds.Degree,
			Direction: t.Direction,
			Magnitude: RadianToDegree(t.Magnitude),
		}
	}
}

func DegreeToRadian(angle numberer.Number) numberer.Number {
	return numberer.Float(angle.Float() * math.Pi / 180)
}

func RadianToDegree(angle numberer.Number) numberer.Number {
	return numberer.Float(angle.Float() * 180 / math.Pi)
}
