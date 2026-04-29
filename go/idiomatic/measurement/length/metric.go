package length

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/imperial"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/metric"
)

type metricLength struct {
	Prefix metric.Prefix
	Amount float64
}

func (this metricLength) GetAmount() float64 {
	return this.Amount
}

func (this metricLength) Imperial(kind imperial.Length) Length {
	ammount := this.Amount
	metricConversionFactor := metric.Prefixes.Factor(this.Prefix)
	imperialConversionFactor := imperial.Lengths.Factor(kind)
	ammount *= metricConversionFactor * millimetersToInches * imperialConversionFactor
	return &ImperialLength{Amount: ammount, Kind: kind}
}

func (this metricLength) Metric(kind metric.Prefix) Length {
	return this
}
