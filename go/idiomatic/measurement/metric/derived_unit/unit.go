package derived_unit

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/boundedinfinity/canonical_model/idiomatic/measurement/metric/base_unit"
// 	"github.com/boundedinfinity/canonical_model/idiomatic/measurement/metric/prefix"
// 	"github.com/boundedinfinity/canonical_model/idiomatic/measurement/quantity"
// 	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
// )

// type Unit struct {
// 	Name     string
// 	Symbol   string
// 	Quantity quantity.Kind
// 	Base     []Base
// }

// func (this Unit) String() string {
// 	return this.Symbol
// }

// func (this Unit) BaseSymbol() string {
// 	components := make([]string, len(this.Base))

// 	for i, element := range this.Base {
// 		components[i] = element.Symbol()
// 	}

// 	return stringer.Join(components, " · ")
// }

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// type Base struct {
// 	Prefix prefix.Kind
// 	Unit   base_unit.Unit
// 	Power  int
// }

// func (this Base) Symbol() string {
// 	var buf strings.Builder

// 	buf.WriteString(prefix.Symbol(this.Prefix))
// 	buf.WriteString(this.Unit.String())

// 	if this.Power != 0 && this.Power != 1 {
// 		fmt.Fprintf(&buf, "^%d", this.Power)
// 	}

// 	return buf.String()
// }

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// var (
// 	Units = []Unit{
// 		{
// 			Name:     "Radian",
// 			Symbol:   "rad",
// 			Quantity: quantity.Kinds.PlaneAngle,
// 		},
// 		{
// 			Name:     "Steradian",
// 			Symbol:   "sr",
// 			Quantity: quantity.Kinds.SolidAngle,
// 		},
// 		{
// 			Name:     "Hertz",
// 			Symbol:   "Hz",
// 			Quantity: quantity.Kinds.Frequency,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Second, Power: -1},
// 			},
// 		},
// 		{
// 			Name:     "Newton",
// 			Symbol:   "N",
// 			Quantity: quantity.Kinds.Force,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 1},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Pascal",
// 			Symbol:   "Pa",
// 			Quantity: quantity.Kinds.Pressure,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: -1},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Pascal",
// 			Symbol:   "Pa",
// 			Quantity: quantity.Kinds.Stress,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: -1},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Joule",
// 			Symbol:   "J",
// 			Quantity: quantity.Kinds.Energy,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Watt",
// 			Symbol:   "W",
// 			Quantity: quantity.Kinds.Power,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -3},
// 			},
// 		},
// 		{
// 			Name:     "Coulomb",
// 			Symbol:   "C",
// 			Quantity: quantity.Kinds.ElectricCharge,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Second, Power: 1},
// 				{Unit: base_unit.Units.Ampere, Power: 1},
// 			},
// 		},
// 		{
// 			Name:     "Volt",
// 			Symbol:   "V",
// 			Quantity: quantity.Kinds.ElectricPotential,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -3},
// 				{Unit: base_unit.Units.Ampere, Power: -1},
// 			},
// 		},
// 		{
// 			Name:     "Ohm",
// 			Symbol:   "Ω",
// 			Quantity: quantity.Kinds.ElectricResistance,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -3},
// 				{Unit: base_unit.Units.Ampere, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Siemens",
// 			Symbol:   "S",
// 			Quantity: quantity.Kinds.ElectricConductance,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: -3},
// 				{Unit: base_unit.Units.Meter, Power: -2},
// 				{Unit: base_unit.Units.Second, Power: 3},
// 				{Unit: base_unit.Units.Ampere, Power: 2},
// 			},
// 		},
// 		{
// 			Name:     "Farad",
// 			Symbol:   "F",
// 			Quantity: quantity.Kinds.ElectricCapacitance,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: -3},
// 				{Unit: base_unit.Units.Meter, Power: -2},
// 				{Unit: base_unit.Units.Second, Power: 4},
// 				{Unit: base_unit.Units.Ampere, Power: 2},
// 			},
// 		},
// 		{
// 			Name:     "Henry",
// 			Symbol:   "H",
// 			Quantity: quantity.Kinds.ElectricInductance,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 				{Unit: base_unit.Units.Ampere, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Tesla",
// 			Symbol:   "T",
// 			Quantity: quantity.Kinds.MagneticFluxDensity,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 				{Unit: base_unit.Units.Ampere, Power: -1},
// 			},
// 		},
// 		{
// 			Name:     "Weber",
// 			Symbol:   "Wb",
// 			Quantity: quantity.Kinds.MagneticFlux,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 				{Unit: base_unit.Units.Ampere, Power: -1},
// 			},
// 		},
// 		{
// 			Name:     "Degree Celsius",
// 			Symbol:   "°C",
// 			Quantity: quantity.Kinds.Temperature,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Kelvin, Power: 1},
// 			},
// 		},
// 		{
// 			Name:     "Lumen",
// 			Symbol:   "lm",
// 			Quantity: quantity.Kinds.LuminousFlux,
// 			Base:     []Base{
// 				// TODO
// 				// {Unit: base_unit.Units.Candela, Power: 1},
// 				// {Unit: base_unit.Units.Steradian, Power: 1},
// 			},
// 		},
// 		{
// 			Name:     "Lux",
// 			Symbol:   "lx",
// 			Quantity: quantity.Kinds.Illuminance,
// 			Base:     []Base{
// 				// TODO
// 				// {Unit: base_unit.Units.Candela, Power: 1},
// 				// {Unit: base_unit.Units.Steradian, Power: 1},
// 				// {Unit: base_unit.Units.Meter, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Becquerel",
// 			Symbol:   "Bq",
// 			Quantity: quantity.Kinds.ActivityOfARadioactiveSubstance,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Second, Power: -1},
// 			},
// 		},
// 		{
// 			Name:     "Gray",
// 			Symbol:   "Gy",
// 			Quantity: quantity.Kinds.AbsorbedDose,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Sievert",
// 			Symbol:   "Sv",
// 			Quantity: quantity.Kinds.DoseEquivalent,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Gram, Power: 3},
// 				{Unit: base_unit.Units.Meter, Power: 2},
// 				{Unit: base_unit.Units.Second, Power: -2},
// 			},
// 		},
// 		{
// 			Name:     "Katal",
// 			Symbol:   "kat",
// 			Quantity: quantity.Kinds.AcatalyticActivity,
// 			Base: []Base{
// 				{Unit: base_unit.Units.Mole, Power: 1},
// 				{Unit: base_unit.Units.Second, Power: -1},
// 			},
// 		},
// 	}
// )

// func init() {

// }
