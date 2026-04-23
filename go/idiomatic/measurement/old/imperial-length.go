package measurement

var (
	imperialLengthUnitsCf   map[ImperialLengthUnit]float64
	imperialLengthUnitsAbbr map[ImperialLengthUnit][]string
)

// https://en.wikipedia.org/wiki/Imperial_units

func init() {
	imperialLengthUnitsCf = map[ImperialLengthUnit]float64{
		ImperialLengthUnits.Twip:         1.0 / 17280.0,
		ImperialLengthUnits.Thou:         1.0 / 12000.0,
		ImperialLengthUnits.Barleycorn:   1.0 / 36.0,
		ImperialLengthUnits.Inch:         1.0 / 12.0,
		ImperialLengthUnits.Hand:         1.0 / 3.0,
		ImperialLengthUnits.Foot:         1.0,
		ImperialLengthUnits.Yard:         3.0,
		ImperialLengthUnits.Chain:        66.0,
		ImperialLengthUnits.Furlong:      660.0,
		ImperialLengthUnits.Mile:         5280.0,
		ImperialLengthUnits.League:       15840.0,
		ImperialLengthUnits.Fathom:       6.0761,
		ImperialLengthUnits.Cable:        607.61,
		ImperialLengthUnits.NauticalMile: 6076.1,
		ImperialLengthUnits.Link:         66.0 / 100.0,
		ImperialLengthUnits.Rod:          66.0 / 4.0,
	}

	imperialLengthUnitsAbbr = map[ImperialLengthUnit][]string{
		ImperialLengthUnits.Twip:         {ImperialLengthUnits.Twip.String()},
		ImperialLengthUnits.Thou:         {ImperialLengthUnits.Thou.String(), "th"},
		ImperialLengthUnits.Barleycorn:   {ImperialLengthUnits.Barleycorn.String()},
		ImperialLengthUnits.Inch:         {ImperialLengthUnits.Hand.String(), "in", `"`},
		ImperialLengthUnits.Hand:         {ImperialLengthUnits.Hand.String(), "hh"},
		ImperialLengthUnits.Foot:         {ImperialLengthUnits.Foot.String(), "ft", "'"},
		ImperialLengthUnits.Yard:         {ImperialLengthUnits.Yard.String(), "yd"},
		ImperialLengthUnits.Chain:        {ImperialLengthUnits.Chain.String(), "ch"},
		ImperialLengthUnits.Furlong:      {ImperialLengthUnits.Furlong.String(), "fur"},
		ImperialLengthUnits.Mile:         {ImperialLengthUnits.Mile.String(), "mi"},
		ImperialLengthUnits.League:       {ImperialLengthUnits.League.String(), "lea"},
		ImperialLengthUnits.Fathom:       {ImperialLengthUnits.Fathom.String(), "ftm"},
		ImperialLengthUnits.Cable:        {ImperialLengthUnits.Cable.String()},
		ImperialLengthUnits.NauticalMile: {ImperialLengthUnits.NauticalMile.String(), "nmi"},
		ImperialLengthUnits.Link:         {ImperialLengthUnits.Link.String()},
		ImperialLengthUnits.Rod:          {ImperialLengthUnits.Rod.String()},
	}
}

// func (t imperialLengthUnits) Convert(number rational.Rational, from, to ImperialLengthUnit) rational.Rational {
// 	return rational.FromFloat(t.convert(number.Float(), from, to))
// }

func (t imperialLengthUnits) convert(number float64, from, to ImperialLengthUnit) float64 {
	float := number
	fromCf, fromOk := imperialLengthUnitsCf[from]
	toCf, toOk := imperialLengthUnitsCf[to]

	if fromOk && toOk {
		float = number * fromCf * toCf
	}

	return float
}
