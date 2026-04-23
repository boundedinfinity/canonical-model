package numberer

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather"
)

func Fraction(numerator, denominator int) Number {
	return &FractionNumber{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

var _ Number = &FractionNumber{}

type FractionNumber struct {
	Numerator   int
	Denominator int
}

// ToFloat implements [Number].
func (this *FractionNumber) ToFloat() Number {
	panic("unimplemented")
}

func (this FractionNumber) String() string {
	return fmt.Sprintf("%d/%d", this.Numerator, this.Denominator)
}

func (this FractionNumber) Copy() FractionNumber {
	return FractionNumber{
		Numerator:   this.Numerator,
		Denominator: this.Denominator,
	}
}

func (this FractionNumber) ToFraction() Number {
	return &this
}

func (this FractionNumber) ToNormal() Number {
	return &this
}

func (this FractionNumber) ToMixed() Number {
	if this.IsProper() {
		return &MixedNumber{Fraction: this.Copy()}
	} else {
		return &MixedNumber{
			Whole: this.Numerator / this.Denominator,
			Fraction: FractionNumber{
				Numerator:   this.Numerator % this.Denominator,
				Denominator: this.Denominator,
			},
		}
	}
}

func (this FractionNumber) Reduce() FractionNumber {
	gcd := GreatestCommonFactor(this.Numerator, this.Denominator)

	return FractionNumber{
		Numerator:   this.Numerator / gcd,
		Denominator: this.Denominator / gcd,
	}
}

func (this FractionNumber) IsZero() bool {
	var zero FractionNumber
	return this == zero
}

func (this *FractionNumber) IsImproper() bool {
	return this.Numerator > this.Denominator
}

func (this *FractionNumber) IsProper() bool {
	return this.Denominator > this.Numerator
}

func (this *FractionNumber) Reciprocal() Number {
	return &FractionNumber{
		Numerator:   this.Denominator,
		Denominator: this.Numerator,
	}
}

func (t FractionNumber) Enumerate(l, h int) []FractionNumber {
	var items []FractionNumber
	item := t.Reduce()
	l = mather.Max(l, item.Denominator)

	for i := l; i <= h; i <<= 1 {
		items = append(items, item)
		item = *Fraction(item.Numerator*2, item.Denominator*2).(*FractionNumber)
	}

	return items
}

func (this FractionNumber) Float() float64 {
	return float64(this.Numerator) / float64(this.Denominator)
}

func (this *FractionNumber) ToImproper() Number {
	return this
}
