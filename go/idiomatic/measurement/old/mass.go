package measurement

import "github.com/boundedinfinity/go-commoner/idiomatic/mather/rational"

type Mass[T ~int] struct {
	Value  rational.Rational[T]  `json:"value,omitempty"`
	System MeasurementSystem     `json:"system,omitempty"`
	Unit   ImperialMassUnit      `json:"unit,omitempty"`
	Format MeasurementFormatType `json:"format,omitempty"`
}

func (t Mass[T]) Metric() Mass[T] {
	return Mass[T]{}
}

func (t Mass[T]) Imperial() Mass[T] {
	return Mass[T]{}
}
