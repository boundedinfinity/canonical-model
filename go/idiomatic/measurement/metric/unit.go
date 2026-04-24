package metric

import "github.com/boundedinfinity/canonical_model/go/idiomatic/measurement/quantity"

type Unit struct {
	Name     string
	Symbol   string
	Quantity quantity.Kind
}

func (this Unit) String() string {
	return this.Symbol
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type units struct {
	Unknown Unit
	Second  Unit
	Meter   Unit
	Gram    Unit
	Ampere  Unit
	Kelvin  Unit
	Mole    Unit
	Candela Unit
}

var Units = units{
	Unknown: Unit{Quantity: quantity.Kinds.Unknown},
	Second:  Unit{Name: "Second", Symbol: "s", Quantity: quantity.Kinds.Time},
	Meter:   Unit{Name: "Meter", Symbol: "m", Quantity: quantity.Kinds.Length},
	Gram:    Unit{Name: "Gram", Symbol: "g", Quantity: quantity.Kinds.Mass},
	Ampere:  Unit{Name: "Ampere", Symbol: "A", Quantity: quantity.Kinds.ElectricCurrent},
	Kelvin:  Unit{Name: "Kelvin", Symbol: "K", Quantity: quantity.Kinds.Temperature},
	Mole:    Unit{Name: "Mole", Symbol: "mol", Quantity: quantity.Kinds.AmountOfSubstance},
	Candela: Unit{Name: "Candela", Symbol: "cd", Quantity: quantity.Kinds.LuminousIntensity},
}
