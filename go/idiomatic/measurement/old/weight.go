package measurement

import "github.com/boundedinfinity/canonical_model/go/idiomatic/mather/rational"

type Weight[T ~int] struct {
	Unit   rational.Rational[T]  `json:"unit,omitempty"`
	System MeasurementSystem     `json:"system,omitempty"`
	Format MeasurementFormatType `json:"format,omitempty"`
}

func (t Weight[T]) Metric() Weight[T] {
	return Weight[T]{}
}

func (t Weight[T]) Imperial() Weight[T] {
	return Weight[T]{}
}
