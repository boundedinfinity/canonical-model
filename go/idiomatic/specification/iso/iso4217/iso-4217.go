package iso4217

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Iso4217Code string
type Iso4217Numeric string

type Iso4217 struct {
	Country string `json:"country"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Numeric string `json:"numeric"`
	Units   string `json:"units"`
}

var Iso4217s = iso4217s{}

func (this iso4217s) Records() []Iso4217 {
	return this.records
}

func (this iso4217s) Codes() []Iso4217Code {
	return this.codes
}

func (this iso4217s) Numerics() []Iso4217Numeric {
	return this.numerics
}

func (this iso4217s) FindRecords(term string) []Iso4217 {
	items, _ := this.FindRecordsErr(term)
	return items
}

var (
	ErrIso4217Invalid   = errors.New("invalid ISO 4217 record")
	errIso4217InvalidFn = func(term string) error {
		return fmt.Errorf("%w : %s", ErrIso4217Invalid, term)
	}
)

func (this iso4217s) FindRecordsErr(term string) ([]Iso4217, error) {
	var items []Iso4217
	var err error
	lc := stringer.Lowercase(term)

	for i, item := range this.lc {
		if item.Code == lc || item.Name == lc || item.Numeric == lc || item.Country == lc {
			items = append(items, this.records[i])
		}
	}

	if len(items) == 0 {
		err = errIso4217InvalidFn(term)
	}

	return items, err
}
