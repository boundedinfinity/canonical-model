package imperial

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Length string

func (this Length) Name() string {
	return Lengths.Name(this)
}

func (this Length) Symbol() string {
	return Lengths.Symbol(this)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var Lengths = lengths{
	Unknown:      "measurement.imperial.length.unknown",
	Inch:         "measurement.imperial.length.inch",
	Foot:         "measurement.imperial.length.foot",
	Twip:         "measurement.imperial.length.twip",
	Thou:         "measurement.imperial.length.thou",
	Barleycorn:   "measurement.imperial.length.barleycorn",
	Hand:         "measurement.imperial.length.hand",
	Yard:         "measurement.imperial.length.yard",
	Chain:        "measurement.imperial.length.chain",
	Furlong:      "measurement.imperial.length.furlong",
	Mile:         "measurement.imperial.length.mile",
	League:       "measurement.imperial.length.league",
	Fathom:       "measurement.imperial.length.fathom",
	Cable:        "measurement.imperial.length.cable",
	NauticalMile: "measurement.imperial.length.nautical-mile",
	Link:         "measurement.imperial.length.link",
	Rod:          "measurement.imperial.length.rod",
}

type lengths struct {
	Unknown      Length
	Inch         Length
	Foot         Length
	Twip         Length
	Thou         Length
	Barleycorn   Length
	Hand         Length
	Yard         Length
	Chain        Length
	Furlong      Length
	Mile         Length
	League       Length
	Fathom       Length
	Cable        Length
	NauticalMile Length
	Link         Length
	Rod          Length
}

func (this lengths) All() []Length {
	return []Length{
		this.Inch,
		this.Foot,
		this.Twip,
		this.Thou,
		this.Barleycorn,
		this.Hand,
		this.Yard,
		this.Chain,
		this.Furlong,
		this.Mile,
		this.League,
		this.Fathom,
		this.Cable,
		this.NauticalMile,
		this.Link,
		this.Rod,
	}
}

func (this lengths) Convert(amount float64, from Length, to Length) float64 {
	return amount * this.Factor(from) / this.Factor(to)
}

const (
	thoFactor        = 1.0 / 17280
	barleycornFactor = 1.0 / 36
	incheFactor      = 1.0 / 12
	handFactor       = 1.0 / 3
	linkFactor       = 66.0 / 100
	rodFactor        = 66.0 / 4
)

func (this lengths) Factor(kind Length) float64 {
	var factor float64

	switch kind {
	case Lengths.Foot:
		factor = 1
	case Lengths.Thou:
		factor = thoFactor
	case Lengths.Barleycorn:
		factor = barleycornFactor
	case Lengths.Inch:
		factor = incheFactor
	case Lengths.Hand:
		factor = handFactor
	case Lengths.Yard:
		factor = 3
	case Lengths.Chain:
		factor = 66
	case Lengths.Furlong:
		factor = 660
	case Lengths.Mile:
		factor = 5280
	case Lengths.League:
		factor = 15840
	case Lengths.Fathom:
		factor = 6.0761
	case Lengths.Cable:
		factor = 607.61
	case Lengths.NauticalMile:
		factor = 6076.1
	case Lengths.Link:
		factor = linkFactor
	case Lengths.Rod:
		factor = rodFactor
	default:
		return 0
	}

	return factor
}

func (this lengths) Name(kind Length) string {
	var name string

	switch kind {
	case Lengths.Foot:
		name = "foot"
	case Lengths.Thou:
		name = "thou"
	case Lengths.Barleycorn:
		name = "barleycorn"
	case Lengths.Inch:
		name = "inch"
	case Lengths.Hand:
		name = "hand"
	case Lengths.Yard:
		name = "yard"
	case Lengths.Chain:
		name = "chain"
	case Lengths.Furlong:
		name = "furlong"
	case Lengths.Mile:
		name = "mile"
	case Lengths.League:
		name = "league"
	case Lengths.Fathom:
		name = "fathom"
	case Lengths.Cable:
		name = "cable"
	case Lengths.NauticalMile:
		name = "nautical mile"
	case Lengths.Link:
		name = "link"
	case Lengths.Rod:
		name = "rod"
	default:
		return ""
	}

	return name
}

func (this lengths) Symbol(kind Length) string {
	var symbol string

	switch kind {
	case Lengths.Foot:
		symbol = "ft"
	case Lengths.Thou:
		symbol = "thou"
	case Lengths.Barleycorn:
		symbol = "barleycorn"
	case Lengths.Inch:
		symbol = "in"
	case Lengths.Hand:
		symbol = "hand"
	case Lengths.Yard:
		symbol = "yd"
	case Lengths.Chain:
		symbol = "ch"
	case Lengths.Furlong:
		symbol = "fur"
	case Lengths.Mile:
		symbol = "mi"
	case Lengths.League:
		symbol = "lea"
	case Lengths.Fathom:
		symbol = "fathom"
	case Lengths.Cable:
		symbol = "cable"
	case Lengths.NauticalMile:
		symbol = "nmi"
	case Lengths.Link:
		symbol = "link"
	case Lengths.Rod:
		symbol = "rod"
	default:
		return ""
	}

	return symbol
}

func (this lengths) ParseOk(s string) (Length, bool) {
	kind := Lengths.Unknown
	var ok bool
	s = stringer.Chain(s, stringer.TrimSpace, stringer.CompactSpace, stringer.ToLower)

	for _, kind = range Lengths.All() {
		if s == stringer.ToLower(Lengths.Symbol(kind)) {
			ok = true
			break
		}

		if s == stringer.ToLower(Lengths.Name(kind)) {
			ok = true
			break
		}
	}

	if kind != Lengths.Unknown && s == `"` {
		kind = Lengths.Inch
		ok = true
	}

	if kind != Lengths.Unknown && s == `'` {
		kind = Lengths.Foot
		ok = true
	}

	return kind, ok
}
