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
	panic("unimplemented")
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
	panic("unimplemented")
}

func (this *FloatNumber) ToImproper() Number {
	panic("unimplemented")
}
