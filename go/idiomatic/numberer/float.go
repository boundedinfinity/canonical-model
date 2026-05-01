package numberer

func Float(value float64) Number {
	return &FloatNumber{
		value: value,
	}
}

var _ Number = &FloatNumber{}

type FloatNumber struct {
	value float64
}

func (this *FloatNumber) ToFloat() Number {
	return this
}

func (this *FloatNumber) ToFraction() Number {
	return &FractionNumber{
		Numerator:   int(this.value * 1000000),
		Denominator: 1000000,
	}
}

func (this *FloatNumber) IsZero() bool {
	return this.value == 0
}

func (this FloatNumber) Float() float64 {
	return this.value
}

func (this *FloatNumber) IsImproper() bool {
	return this.value > 1
}

func (this *FloatNumber) IsProper() bool {
	return this.value < 1
}

func (this *FloatNumber) Reciprocal() Number {
	return &FloatNumber{value: 1 / this.value}
}

func (this *FloatNumber) ToMixed() Number {
	return &MixedNumber{
		Whole: int(this.value),
		Fraction: FractionNumber{
			Numerator:   int((this.value - float64(int(this.value))) * 1000000),
			Denominator: 1000000,
		},
	}
}

func (this *FloatNumber) ToImproper() Number {
	return &FractionNumber{
		Numerator:   int(this.value * 1000000),
		Denominator: 1000000,
	}
}
