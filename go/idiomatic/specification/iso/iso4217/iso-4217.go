package iso4217

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Iso4217 struct {
	Country string `json:"country"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Numeric string `json:"numeric"`
	Units   string `json:"units"`
}

var Iso4217s = iso4217s{}

type iso4217s struct {
	records []Iso4217
}

func (this iso4217s) Records() []Iso4217 {
	return this.records
}
func (this iso4217s) FindOk(term string) ([]Iso4217, bool) {
	lc := stringer.ToLower(term)
	match := func(item Iso4217) bool {
		return stringer.Contains(item.Name, lc) ||
			stringer.Contains(item.Numeric, lc) ||
			stringer.Contains(item.Country, lc)
	}

	found := slicer.FilterFunc(match, this.records...)
	return found, len(found) > 0
}

var (
	ErrIso4217Invalid   = errors.New("invalid ISO 4217 record")
	errIso4217InvalidFn = func(term string) error {
		return fmt.Errorf("%w : %s", ErrIso4217Invalid, term)
	}
)

func (this iso4217s) FindErr(term string) ([]Iso4217, error) {
	found, ok := this.FindOk(term)
	var err error

	if !ok {
		err = errIso4217InvalidFn(term)
	}

	return found, err
}
