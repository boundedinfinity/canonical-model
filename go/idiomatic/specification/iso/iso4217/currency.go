package iso4217

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Currency struct {
	Country string `json:"country"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Numeric string `json:"numeric"`
	Units   string `json:"units"`
}

type currencies struct {
	all      []Currency
	names    map[string]*Currency
	codes    map[string]*Currency
	numerics map[string]*Currency
}

func (this currencies) All() []Currency {
	return this.all[:]
}

func (this currencies) FindOk(term string) ([]Currency, bool) {
	lc := stringer.ToLower(term)
	var found []Currency
	var currency *Currency
	var ok bool

	if currency, ok = this.names[lc]; ok {
		found = append(found, *currency)
	}

	if !ok {
		if currency, ok = this.codes[lc]; ok {
			found = append(found, *currency)
		}
	}

	if !ok {
		if currency, ok = this.numerics[lc]; ok {
			found = append(found, *currency)
		}
	}

	if !ok {
		var name string

		for name, currency = range this.names {
			if stringer.Contains(name, lc) {
				found = append(found, *currency)
			}
		}

		for name, currency = range this.codes {
			if stringer.Contains(name, lc) {
				found = append(found, *currency)
			}
		}

		for name, currency = range this.numerics {
			if stringer.Contains(name, lc) {
				found = append(found, *currency)
			}
		}
	}

	return found, len(found) > 0
}

var (
	ErrIso4217Invalid   = errors.New("invalid ISO 4217 record")
	errIso4217InvalidFn = func(term string) error {
		return fmt.Errorf("%w : %s", ErrIso4217Invalid, term)
	}
)

func (this currencies) FindErr(term string) ([]Currency, error) {
	found, ok := this.FindOk(term)
	var err error

	if !ok {
		err = errIso4217InvalidFn(term)
	}

	return found, err
}

func init() {
	for i := range Currencies.all {
		name := stringer.ToLower(Currencies.all[i].Name)
		Currencies.names[stringer.ToLower(Currencies.all[i].Name)] = &Currencies.all[i]

		if idx := stringer.Index(name, "("); idx != -1 {
			name2 := name[:idx]
			name2 = stringer.TrimSpaceRight(name2)
			if name2 != name {
				Currencies.names[name2] = &Currencies.all[i]
			}
		}

		Currencies.codes[stringer.ToLower(Currencies.all[i].Code)] = &Currencies.all[i]

		numeric := stringer.ToLower(Currencies.all[i].Numeric)
		Currencies.numerics[numeric] = &Currencies.all[i]

		numeric2 := stringer.TrimPrefix(numeric, "0")
		if numeric2 != numeric {
			Currencies.numerics[numeric2] = &Currencies.all[i]
		}
	}
}
