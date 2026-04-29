package length

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/imperial"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/metric"
)

type ImperialLength struct {
	Kind   imperial.Length
	Amount float64
}

func (this ImperialLength) GetAmount() float64 {
	return this.Amount
}

func (this ImperialLength) Imperial(kind imperial.Length) Length {
	return &ImperialLength{
		Kind:   kind,
		Amount: imperial.Lengths.Convert(this.Amount, this.Kind, kind),
	}
}

func (this ImperialLength) Metric(prefix metric.Prefix) Length {
	ammount := imperial.Lengths.Convert(this.Amount, this.Kind, imperial.Lengths.Inch)
	ammount *= inchesToMillimeters
	return &metricLength{
		Prefix: prefix,
		Amount: ammount,
	}
}
