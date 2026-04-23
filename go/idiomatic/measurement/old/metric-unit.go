package measurement

var (
	metricUnitsCf   map[MetricUnit]float64
	metricUnitsAbbr map[MetricUnit][]string
)

func init() {
	metricUnitsCf = map[MetricUnit]float64{
		MetricUnits.Tera:  10e12,
		MetricUnits.Giga:  10e9,
		MetricUnits.Mega:  10e6,
		MetricUnits.Kilo:  10e3,
		MetricUnits.Hecto: 10e2,
		MetricUnits.Deca:  10e1,
		MetricUnits.Unit:  1.0,
		MetricUnits.Deci:  10e-1,
		MetricUnits.Centi: 10e-2,
		MetricUnits.Milli: 10e-3,
		MetricUnits.Micro: 10e-6,
		MetricUnits.Nano:  10e-9,
		MetricUnits.Pico:  10e-12,
	}
	metricUnitsAbbr = map[MetricUnit][]string{
		MetricUnits.Tera:  {MetricUnits.Tera.String(), "T"},
		MetricUnits.Giga:  {MetricUnits.Giga.String(), "G"},
		MetricUnits.Mega:  {MetricUnits.Mega.String(), "M"},
		MetricUnits.Kilo:  {MetricUnits.Kilo.String(), "k"},
		MetricUnits.Hecto: {MetricUnits.Hecto.String(), "h"},
		MetricUnits.Deca:  {MetricUnits.Deca.String(), "da"},
		MetricUnits.Unit:  {},
		MetricUnits.Deci:  {MetricUnits.Deci.String(), "d"},
		MetricUnits.Centi: {MetricUnits.Centi.String(), "c"},
		MetricUnits.Milli: {MetricUnits.Milli.String(), "m"},
		MetricUnits.Micro: {MetricUnits.Micro.String(), "u", "μ"},
		MetricUnits.Nano:  {MetricUnits.Nano.String(), "n"},
		MetricUnits.Pico:  {MetricUnits.Pico.String(), "p"},
	}
}

// func (t metricUnits) Convert(number rational.Rational, from, to MetricUnit) rational.Rational {
// 	return rational.FromFloat(t.convert(number.Float(), from, to))
// }

func (t metricUnits) convert(number float64, from, to MetricUnit) float64 {
	var float float64
	fromCf, fromOk := metricUnitsCf[from]
	toCf, toOk := metricUnitsCf[to]

	if fromOk && toOk {
		float = number * fromCf * toCf
	}

	return float
}
