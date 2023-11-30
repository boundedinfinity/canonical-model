//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package json_schema

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type Type string

// /////////////////////////////////////////////////////////////////
//  Type Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t Type) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  Type JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Type) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *Type) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, Types.Parse)
}

// /////////////////////////////////////////////////////////////////
//  Type YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Type) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *Type) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, Types.Parse)
}

// /////////////////////////////////////////////////////////////////
//  Type SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Type) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *Type) Scan(value interface{}) error {
	return enumer.Scan(value, t, Types.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type types struct {
	Array   Type
	String  Type
	Number  Type
	Integer Type
	Object  Type
	Ref     Type
	Boolean Type
	Values  []Type
	Err     error
}

var Types = types{
	Array:   Type("array"),
	String:  Type("string"),
	Number:  Type("number"),
	Integer: Type("integer"),
	Object:  Type("object"),
	Ref:     Type("ref"),
	Boolean: Type("boolean"),
	Err:     fmt.Errorf("invalid Type"),
}

func init() {
	Types.Values = []Type{
		Types.Array,
		Types.String,
		Types.Number,
		Types.Integer,
		Types.Object,
		Types.Ref,
		Types.Boolean,
	}
}

func (t types) newErr(a any, values ...Type) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		Types.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t types) ParseFrom(v string, values ...Type) (Type, error) {
	var found Type
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, Type](v)(value) {
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

func (t types) Parse(v string) (Type, error) {
	return t.ParseFrom(v, Types.Values...)
}

func (t types) IsFrom(v string, values ...Type) bool {
	for _, value := range values {
		if enumer.IsEq[string, Type](v)(value) {
			return true
		}
	}
	return false
}

func (t types) Is(v string) bool {
	return t.IsFrom(v, Types.Values...)
}
