package numberer

import "fmt"

func Mixed(whole int, numerator int, denominator int) Number {
	return &MixedNumber{
		Whole: whole,
		Fraction: FractionNumber{
			Numerator:   numerator,
			Denominator: denominator,
		},
	}
}

var _ Number = &MixedNumber{}

type MixedNumber struct {
	Whole    int
	Fraction FractionNumber
}

func (this MixedNumber) String() string {
	if this.Whole == 0 {
		return this.Fraction.String()
	}

	return fmt.Sprintf("%d %s", this.Whole, this.Fraction)
}

func (this MixedNumber) IsZero() bool {
	return this.Whole == 0 && this.Fraction.IsZero()
}

func (this MixedNumber) Copy() MixedNumber {
	return MixedNumber{
		Whole:    this.Whole,
		Fraction: this.Fraction.Copy(),
	}
}

func (this *MixedNumber) ToFloat() Number {
	improper := this.ToImproper().(*FractionNumber)
	return &FloatNumber{value: float64(improper.Numerator) / float64(improper.Denominator)}
}

func (this *MixedNumber) ToImproper() Number {
	if this.Whole == 0 {
		return this
	} else {
		return &FractionNumber{
			Numerator:   this.Whole*this.Fraction.Denominator + this.Fraction.Numerator,
			Denominator: this.Fraction.Denominator,
		}
	}
}

func (this *MixedNumber) ToFraction() Number {
	if this.Whole == 0 {
		return this
	} else {
		return &FractionNumber{
			Numerator:   this.Whole*this.Fraction.Denominator + this.Fraction.Numerator,
			Denominator: this.Fraction.Denominator,
		}
	}
}

func (this *MixedNumber) ToMixed() Number {
	return this
}

func (this MixedNumber) IsImproper() bool {
	if this.Whole == 0 {
		return this.Fraction.IsImproper()
	}
	return true
}

func (this MixedNumber) IsProper() bool {
	if this.Whole == 0 {
		return this.Fraction.IsProper()
	}
	return false
}

func (this MixedNumber) Reduce() MixedNumber {
	return MixedNumber{
		Whole:    this.Whole,
		Fraction: this.Fraction.Reduce(),
	}
}

func (t MixedNumber) Enumerate(l, h int) []MixedNumber {
	var items []MixedNumber
	fractions := t.Fraction.Enumerate(l, h)

	for _, fraction := range fractions {
		items = append(items, MixedNumber{
			Whole:    t.Whole,
			Fraction: fraction,
		})
	}

	return items
}

func (this MixedNumber) Reciprocal() Number {
	fraction := this.ToFraction().(*FractionNumber)
	return fraction.Reciprocal()
}

func (this MixedNumber) Float() float64 {
	fraction := this.ToFraction().(*FractionNumber)
	return float64(fraction.Numerator) / float64(fraction.Denominator)
}
