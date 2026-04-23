package metric

// https://en.wikipedia.org/wiki/Metric_prefix

type Prefix string

type prefixes struct {
	Unknown Prefix
	Quetta  Prefix
	Ronna   Prefix
	Yotta   Prefix
	Zetta   Prefix
	Exa     Prefix
	Peta    Prefix
	Tera    Prefix
	Giga    Prefix
	Mega    Prefix
	Kilo    Prefix
	Hecto   Prefix
	Deca    Prefix
	Unit    Prefix
	Deci    Prefix
	Centi   Prefix
	Milli   Prefix
	Micro   Prefix
	Nano    Prefix
	Pico    Prefix
	Femto   Prefix
	Atto    Prefix
	Zepto   Prefix
	Yocto   Prefix
	Quecto  Prefix
}

var Prefixes = prefixes{
	Unknown: "measurement.metric.prefix.unknown",
	Unit:    "measurement.metric.prefix.unit",
	Tera:    "measurement.metric.prefix.tera",
	Giga:    "measurement.metric.prefix.giga",
	Mega:    "measurement.metric.prefix.mega",
	Kilo:    "measurement.metric.prefix.kilo",
	Hecto:   "measurement.metric.prefix.hecto",
	Deca:    "measurement.metric.prefix.deca",
	Deci:    "measurement.metric.prefix.deci",
	Centi:   "measurement.metric.prefix.centi",
	Milli:   "measurement.metric.prefix.milli",
	Micro:   "measurement.metric.prefix.micro",
	Nano:    "measurement.metric.prefix.nano",
	Pico:    "measurement.metric.prefix.pico",
	Femto:   "measurement.metric.prefix.femto",
	Atto:    "measurement.metric.prefix.atto",
	Zepto:   "measurement.metric.prefix.zepto",
	Yocto:   "measurement.metric.prefix.yocto",
	Quecto:  "measurement.metric.prefix.quecto",
	Quetta:  "measurement.metric.prefix.quetta",
	Ronna:   "measurement.metric.prefix.ronna",
	Yotta:   "measurement.metric.prefix.yotta",
	Peta:    "measurement.metric.prefix.peta",
	Exa:     "measurement.metric.prefix.exa",
	Zetta:   "measurement.metric.prefix.zetta",
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this prefixes) All() []Prefix {
	return []Prefix{
		this.Unit,
		this.Tera,
		this.Giga,
		this.Mega,
		this.Kilo,
		this.Hecto,
		this.Deca,
		this.Deci,
		this.Centi,
		this.Milli,
		this.Micro,
		this.Nano,
		this.Pico,
		this.Femto,
		this.Atto,
		this.Zepto,
		this.Yocto,
		this.Quecto,
		this.Quetta,
		this.Ronna,
		this.Yotta,
		this.Peta,
		this.Exa,
		this.Zetta,
	}
}

func (this prefixes) Factor(kind Prefix) float64 {
	var factor float64

	switch kind {
	case Prefixes.Unit:
		factor = 1
	case Prefixes.Tera:
		factor = 1e12
	case Prefixes.Giga:
		factor = 1e9
	case Prefixes.Mega:
		factor = 1e6
	case Prefixes.Kilo:
		factor = 1e3
	case Prefixes.Hecto:
		factor = 1e2
	case Prefixes.Deca:
		factor = 1e1
	case Prefixes.Deci:
		factor = 1e-1
	case Prefixes.Centi:
		factor = 1e-2
	case Prefixes.Milli:
		factor = 1e-3
	case Prefixes.Micro:
		factor = 1e-6
	case Prefixes.Nano:
		factor = 1e-9
	case Prefixes.Pico:
		factor = 1e-12
	case Prefixes.Femto:
		factor = 1e-15
	case Prefixes.Atto:
		factor = 1e-18
	case Prefixes.Zepto:
		factor = 1e-21
	case Prefixes.Yocto:
		factor = 1e-24
	case Prefixes.Quecto:
		factor = 1e-27
	case Prefixes.Quetta:
		factor = 1e30
	case Prefixes.Ronna:
		factor = 1e27
	case Prefixes.Yotta:
		factor = 1e24
	case Prefixes.Zetta:
		factor = 1e21
	case Prefixes.Exa:
		factor = 1e18
	case Prefixes.Peta:
		factor = 1e15
	default:
		return 0
	}

	return factor
}

func (this prefixes) Name(kind Prefix) string {
	var name string

	switch kind {
	case Prefixes.Unit:
		name = ""
	case Prefixes.Tera:
		name = "Tera"
	case Prefixes.Giga:
		name = "Giga"
	case Prefixes.Mega:
		name = "Mega"
	case Prefixes.Kilo:
		name = "Kilo"
	case Prefixes.Hecto:
		name = "Hecto"
	case Prefixes.Deca:
		name = "Deca"
	case Prefixes.Deci:
		name = "Deci"
	case Prefixes.Centi:
		name = "Centi"
	case Prefixes.Milli:
		name = "Milli"
	case Prefixes.Micro:
		name = "Micro"
	case Prefixes.Nano:
		name = "Nano"
	case Prefixes.Pico:
		name = "Pico"
	case Prefixes.Femto:
		name = "Femto"
	case Prefixes.Atto:
		name = "Atto"
	case Prefixes.Zepto:
		name = "Zepto"
	case Prefixes.Yocto:
		name = "Yocto"
	case Prefixes.Quecto:
		name = "Quecto"
	case Prefixes.Quetta:
		name = "Quetta"
	case Prefixes.Ronna:
		name = "Ronna"
	case Prefixes.Yotta:
		name = "Yotta"
	case Prefixes.Zetta:
		name = "Zetta"
	case Prefixes.Exa:
		name = "Exa"
	case Prefixes.Peta:
		name = "Peta"
	default:
		return ""
	}

	return name
}

func (this prefixes) Symbol(kind Prefix) string {
	var name string

	switch kind {
	case Prefixes.Unit:
		name = ""
	case Prefixes.Tera:
		name = "T"
	case Prefixes.Giga:
		name = "G"
	case Prefixes.Mega:
		name = "M"
	case Prefixes.Kilo:
		name = "k"
	case Prefixes.Hecto:
		name = "h"
	case Prefixes.Deca:
		name = "da"
	case Prefixes.Deci:
		name = "d"
	case Prefixes.Centi:
		name = "c"
	case Prefixes.Milli:
		name = "m"
	case Prefixes.Micro:
		name = "μ"
	case Prefixes.Nano:
		name = "n"
	case Prefixes.Pico:
		name = "p"
	case Prefixes.Femto:
		name = "f"
	case Prefixes.Atto:
		name = "a"
	case Prefixes.Zepto:
		name = "z"
	case Prefixes.Yocto:
		name = "y"
	case Prefixes.Quecto:
		name = "q"
	case Prefixes.Quetta:
		name = "Q"
	case Prefixes.Ronna:
		name = "R"
	case Prefixes.Yotta:
		name = "Y"
	case Prefixes.Zetta:
		name = "Z"
	case Prefixes.Exa:
		name = "E"
	case Prefixes.Peta:
		name = "P"
	default:
		return ""
	}

	return name
}
