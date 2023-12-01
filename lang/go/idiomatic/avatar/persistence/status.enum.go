//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Manual changes will be overwritten.                                              *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package persistence

import (
	"database/sql/driver"
	"fmt"

	"github.com/boundedinfinity/enumer"
)

type Status string

// /////////////////////////////////////////////////////////////////
//  Status Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t Status) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  Status JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Status) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *Status) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, Statuses.Parse)
}

// /////////////////////////////////////////////////////////////////
//  Status YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Status) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *Status) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, Statuses.Parse)
}

// /////////////////////////////////////////////////////////////////
//  Status SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t Status) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *Status) Scan(value interface{}) error {
	return enumer.Scan(value, t, Statuses.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type statuses struct {
	Active       Status
	Inactive     Status
	Discontinued Status
	Values       []Status
	Err          error
}

var Statuses = statuses{
	Active:       Status("active"),
	Inactive:     Status("inactive"),
	Discontinued: Status("discontinued"),
	Err:          fmt.Errorf("invalid Status"),
}

func init() {
	Statuses.Values = []Status{
		Statuses.Active,
		Statuses.Inactive,
		Statuses.Discontinued,
	}
}

func (t statuses) newErr(a any, values ...Status) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		Statuses.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t statuses) ParseFrom(v string, values ...Status) (Status, error) {
	var found Status
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, Status](v)(value) {
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

func (t statuses) Parse(v string) (Status, error) {
	return t.ParseFrom(v, Statuses.Values...)
}

func (t statuses) IsFrom(v string, values ...Status) bool {
	for _, value := range values {
		if enumer.IsEq[string, Status](v)(value) {
			return true
		}
	}
	return false
}

func (t statuses) Is(v string) bool {
	return t.IsFrom(v, Statuses.Values...)
}