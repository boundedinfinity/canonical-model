//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package phone

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

// /////////////////////////////////////////////////////////////////
//  NapaPhoneSeparatorFormat Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t NapaPhoneSeparatorFormat) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  NapaPhoneSeparatorFormat JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t NapaPhoneSeparatorFormat) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *NapaPhoneSeparatorFormat) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, NapaPhoneSeparatorFormats.Parse)
}

// /////////////////////////////////////////////////////////////////
//  NapaPhoneSeparatorFormat YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t NapaPhoneSeparatorFormat) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *NapaPhoneSeparatorFormat) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, NapaPhoneSeparatorFormats.Parse)
}

// /////////////////////////////////////////////////////////////////
//  NapaPhoneSeparatorFormat SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t NapaPhoneSeparatorFormat) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *NapaPhoneSeparatorFormat) Scan(value interface{}) error {
	return enumer.Scan(value, t, NapaPhoneSeparatorFormats.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

var NapaPhoneSeparatorFormats = struct {
	napaPhoneSeparatorFormats
	Err    error
	Values []NapaPhoneSeparatorFormat
}{
	napaPhoneSeparatorFormats: napaPhoneSeparatorFormats{
		DotSeparated:         NapaPhoneSeparatorFormat("DotSeparated"),
		DashSeparated:        NapaPhoneSeparatorFormat("DashSeparated"),
		ParenthesesAndDashes: NapaPhoneSeparatorFormat("ParenthesesAndDashes"),
		None:                 NapaPhoneSeparatorFormat("None"),
	},
	Err: fmt.Errorf("invalid NapaPhoneSeparatorFormat"),
}

func init() {
	NapaPhoneSeparatorFormats.Values = []NapaPhoneSeparatorFormat{
		NapaPhoneSeparatorFormats.DotSeparated,
		NapaPhoneSeparatorFormats.DashSeparated,
		NapaPhoneSeparatorFormats.ParenthesesAndDashes,
		NapaPhoneSeparatorFormats.None,
	}
}

func (t napaPhoneSeparatorFormats) newErr(a any, values ...NapaPhoneSeparatorFormat) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		NapaPhoneSeparatorFormats.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t napaPhoneSeparatorFormats) ParseFrom(v string, values ...NapaPhoneSeparatorFormat) (NapaPhoneSeparatorFormat, error) {
	var found NapaPhoneSeparatorFormat
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, NapaPhoneSeparatorFormat](v)(value) {
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

func (t napaPhoneSeparatorFormats) Parse(v string) (NapaPhoneSeparatorFormat, error) {
	return t.ParseFrom(v, NapaPhoneSeparatorFormats.Values...)
}

func (t napaPhoneSeparatorFormats) IsFrom(v string, values ...NapaPhoneSeparatorFormat) bool {
	for _, value := range values {
		if enumer.IsEq[string, NapaPhoneSeparatorFormat](v)(value) {
			return true
		}
	}
	return false
}

func (t napaPhoneSeparatorFormats) Is(v string) bool {
	return t.IsFrom(v, NapaPhoneSeparatorFormats.Values...)
}
