//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package people

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type NameFormat string

// /////////////////////////////////////////////////////////////////
//  NameFormat Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t NameFormat) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  NameFormat JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t NameFormat) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *NameFormat) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, NameFormats.Parse)
}

// /////////////////////////////////////////////////////////////////
//  NameFormat YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t NameFormat) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *NameFormat) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, NameFormats.Parse)
}

// /////////////////////////////////////////////////////////////////
//  NameFormat SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t NameFormat) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *NameFormat) Scan(value interface{}) error {
	return enumer.Scan(value, t, NameFormats.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type nameFormats struct {
	GivenNameFamilyName NameFormat
	FamilyNameGivenName NameFormat
	Values              []NameFormat
	Err                 error
}

var NameFormats = nameFormats{
	GivenNameFamilyName: NameFormat("given-name-family-name"),
	FamilyNameGivenName: NameFormat("family-name-given-name"),
	Err:                 fmt.Errorf("invalid NameFormat"),
}

func init() {
	NameFormats.Values = []NameFormat{
		NameFormats.GivenNameFamilyName,
		NameFormats.FamilyNameGivenName,
	}
}

func (t nameFormats) newErr(a any, values ...NameFormat) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		NameFormats.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t nameFormats) ParseFrom(v string, values ...NameFormat) (NameFormat, error) {
	var found NameFormat
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, NameFormat](v)(value) {
			found = value
			ok = true
			break
		}
	}

	if !ok {
		return found, t.newErr(v, values...)
	}

	return found, nil
}

func (t nameFormats) Parse(v string) (NameFormat, error) {
	return t.ParseFrom(v, NameFormats.Values...)
}

func (t nameFormats) IsFrom(v string, values ...NameFormat) bool {
	for _, value := range values {
		if enumer.IsEq[string, NameFormat](v)(value) {
			return true
		}
	}
	return false
}

func (t nameFormats) Is(v string) bool {
	return t.IsFrom(v, NameFormats.Values...)
}
