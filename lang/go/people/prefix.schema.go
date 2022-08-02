package main

import "github.com/boundedinfinity/optioner"

type Prefix struct {
	Id           string
	Text         string
	Abbreviation optioner.Option[string]
	Description  optioner.Option[string]
}

func (t Prefix) String() string {
	if t.Abbreviation.Defined() {
		return t.Abbreviation.Get()
	} else {
		return t.Text
	}
}

func (t Prefix) Validate() error {
	return nil
}
