package main

import "github.com/boundedinfinity/optioner"

type Suffix struct {
	Id           string
	Text         string
	Abbreviation optioner.Option[string]
	Description  optioner.Option[string]
}

func (t Suffix) Validate() error {
	return nil
}

func (t Suffix) String() string {
	if t.Abbreviation.Defined() {
		return t.Abbreviation.Get()
	} else {
		return t.Text
	}
}
