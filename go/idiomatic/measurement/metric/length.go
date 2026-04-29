package metric

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

var Lengths = lengths{}

type lengths struct{}

func (this lengths) Symbol(prefix Prefix) string {
	return Prefixes.Symbol(prefix) + Units.Meter.Symbol
}

func (this lengths) ParseOk(s string) (Prefix, bool) {
	var ok bool
	prefix := Prefixes.Unknown
	m := stringer.ToLower(Units.Meter.Symbol)
	s = stringer.Chain(s, stringer.TrimSpace, stringer.CompactSpace, stringer.ToLower)

	for _, prefix = range Prefixes.All() {
		target := stringer.ToLower(prefix) + m

		if s == target {
			ok = true
			break
		}
	}

	return prefix, ok
}
