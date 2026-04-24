package length

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/measurement/imperial"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/measurement/metric"
)

type imperialLength struct {
	Kind   imperial.Length
	Amount float64
}

func (this imperialLength) GetAmount() float64 {
	return this.Amount
}

func (this imperialLength) Imperial(kind imperial.Length) Length {
	return this
}

func (this imperialLength) Metric(kind metric.Prefix) Length {
	ammount := this.Amount
	imperialConversionFactor := imperial.Lengths.Factor(this.Kind)
	metricConversionFactor := metric.Prefixes.Factor(kind)
	ammount *= imperialConversionFactor * inchToMillimeters * metricConversionFactor
	return &metricLength{Amount: ammount, Prefix: kind}
}
