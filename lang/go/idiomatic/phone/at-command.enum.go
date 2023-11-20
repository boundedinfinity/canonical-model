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

type AtCommand string

// /////////////////////////////////////////////////////////////////
//  AtCommand Stringer implemenation
// /////////////////////////////////////////////////////////////////

func (t AtCommand) String() string {
	return string(t)
}

// /////////////////////////////////////////////////////////////////
//  AtCommand JSON marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t AtCommand) MarshalJSON() ([]byte, error) {
	return enumer.MarshalJSON(t)
}

func (t *AtCommand) UnmarshalJSON(data []byte) error {
	return enumer.UnmarshalJSON(data, t, AtCommands.Parse)
}

// /////////////////////////////////////////////////////////////////
//  AtCommand YAML marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t AtCommand) MarshalYAML() (interface{}, error) {
	return enumer.MarshalYAML(t)
}

func (t *AtCommand) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumer.UnmarshalYAML(unmarshal, t, AtCommands.Parse)
}

// /////////////////////////////////////////////////////////////////
//  AtCommand SQL Database marshal/unmarshal implemenation
// /////////////////////////////////////////////////////////////////

func (t AtCommand) Value() (driver.Value, error) {
	return enumer.Value(t)
}

func (t *AtCommand) Scan(value interface{}) error {
	return enumer.Scan(value, t, AtCommands.Parse)
}

// /////////////////////////////////////////////////////////////////
//
//  Enumeration
//
// /////////////////////////////////////////////////////////////////

type atCommands struct {
	Pause2Seconds   AtCommand
	PauseIndefinite AtCommand
	Values          []AtCommand
	Err             error
}

var AtCommands = atCommands{
	Pause2Seconds:   AtCommand("Pause2Seconds"),
	PauseIndefinite: AtCommand("PauseIndefinite"),
	Err:             fmt.Errorf("invalid AtCommand"),
}

func init() {
	AtCommands.Values = []AtCommand{
		AtCommands.Pause2Seconds,
		AtCommands.PauseIndefinite,
	}
}

func (t atCommands) newErr(a any, values ...AtCommand) error {
	return fmt.Errorf(
		"invalid %w value '%v'. Must be one of %v",
		AtCommands.Err,
		a,
		enumer.Join(values, ", "),
	)
}

func (t atCommands) ParseFrom(v string, values ...AtCommand) (AtCommand, error) {
	var found AtCommand
	var ok bool

	for _, value := range values {
		if enumer.IsEq[string, AtCommand](v)(value) {
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

func (t atCommands) Parse(v string) (AtCommand, error) {
	return t.ParseFrom(v, AtCommands.Values...)
}

func (t atCommands) IsFrom(v string, values ...AtCommand) bool {
	for _, value := range values {
		if enumer.IsEq[string, AtCommand](v)(value) {
			return true
		}
	}
	return false
}

func (t atCommands) Is(v string) bool {
	return t.IsFrom(v, AtCommands.Values...)
}
