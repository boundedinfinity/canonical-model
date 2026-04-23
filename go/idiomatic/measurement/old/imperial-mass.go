package measurement

var (
	imperialMassUnitsCf   map[ImperialMassUnit]float64
	imperialMassUnitsAbbr map[ImperialMassUnit][]string
	unitToSystem          map[ImperialMassUnit]MeasurementSystem
)

// https://en.wikipedia.org/wiki/Imperial_units

func init() {
	unitToSystem = map[ImperialMassUnit]MeasurementSystem{
		ImperialMassUnits.Grain:         MeasurementSystems.Imperial,
		ImperialMassUnits.Drachm:        MeasurementSystems.Imperial,
		ImperialMassUnits.Ounce:         MeasurementSystems.Imperial,
		ImperialMassUnits.Pound:         MeasurementSystems.Imperial,
		ImperialMassUnits.Stone:         MeasurementSystems.Imperial,
		ImperialMassUnits.Quarter:       MeasurementSystems.Imperial,
		ImperialMassUnits.HundredWeight: MeasurementSystems.Imperial,
		ImperialMassUnits.Ton:           MeasurementSystems.Imperial,
		ImperialMassUnits.Slug:          MeasurementSystems.Imperial,
	}

	imperialMassUnitsCf = map[ImperialMassUnit]float64{
		ImperialMassUnits.Grain:         1.0 / 7000.0,
		ImperialMassUnits.Drachm:        1.0 / 256.0,
		ImperialMassUnits.Ounce:         1.0 / 16.0,
		ImperialMassUnits.Pound:         1.0 / 1.0,
		ImperialMassUnits.Stone:         14.0,
		ImperialMassUnits.Quarter:       28.0,
		ImperialMassUnits.HundredWeight: 112,
		ImperialMassUnits.Ton:           2240,
		ImperialMassUnits.Slug:          32.17404856,
	}

	imperialMassUnitsAbbr = map[ImperialMassUnit][]string{
		ImperialMassUnits.Grain:         {ImperialMassUnits.Grain.String(), "gr"},
		ImperialMassUnits.Drachm:        {ImperialMassUnits.Drachm.String(), "dr"},
		ImperialMassUnits.Ounce:         {ImperialMassUnits.Ounce.String(), "oz"},
		ImperialMassUnits.Pound:         {ImperialMassUnits.Pound.String(), "lb"},
		ImperialMassUnits.Stone:         {ImperialMassUnits.Stone.String(), "st"},
		ImperialMassUnits.Quarter:       {ImperialMassUnits.Quarter.String(), "gtr", "qr"},
		ImperialMassUnits.HundredWeight: {ImperialMassUnits.HundredWeight.String(), "cwt"},
		ImperialMassUnits.Ton:           {ImperialMassUnits.Ton.String()},
		ImperialMassUnits.Slug:          {ImperialMassUnits.Slug.String()},
	}
}

// func (t imperialMassUnits) Convert(number rational.Rational, from, to ImperialMassUnit) rational.Rational {
// 	return rational.FromFloat(t.convert(number.Float(), from, to))
// }

func (t imperialMassUnits) convert(number float64, from, to ImperialMassUnit) float64 {
	float := number
	fromCf, fromOk := imperialMassUnitsCf[from]
	toCf, toOk := imperialMassUnitsCf[to]

	if fromOk && toOk {
		float = number * fromCf * toCf
	}

	return float
}

// func (t imperialMassUnits) Pounds(unit float64) Mass {
// 	return Mass{
// 		Value:  rational.FromFloat(unit),
// 		System: unitToSystem[ImperialMassUnits.Pound],
// 		Unit:   ImperialMassUnits.Pound,
// 	}
// }

// func (t imperialMassUnits) Tons(unit float64) Mass {
// 	return Mass{
// 		Value:  rational.FromFloat(unit),
// 		System: unitToSystem[ImperialMassUnits.Ton],
// 		Unit:   ImperialMassUnits.Ton,
// 	}
// }
