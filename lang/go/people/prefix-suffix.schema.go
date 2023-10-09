package people

import (
	"github.com/boundedinfinity/optioner"
	"github.com/google/uuid"
)

type Prefix struct {
	Id           optioner.Option[uuid.UUID]          `json:"id,omitempty"`
	Text         string                              `json:"text,omitempty"`
	Abbreviation optioner.Option[[]string]           `json:"abbreviation,omitempty"`
	Description  optioner.Option[string]             `json:"description,omitempty"`
	Format       optioner.Option[PrefixSuffixFormat] `json:"format,omitempty"`
}

func (t Prefix) String() string {
	var s string

	if (t.Format.Empty() || t.Format.Get() == Abbreviation) && len(t.Abbreviation.Get()) > 0 {
		s = t.Abbreviation.Get()[0]
	} else {
		s = t.Text
	}

	return s
}

func (t Prefix) Validate(groups ...string) error {
	return nil
}

type Suffix struct {
	Id           optioner.Option[uuid.UUID]          `json:"id,omitempty"`
	Text         string                              `json:"text,omitempty"`
	Abbreviation optioner.Option[[]string]           `json:"abbreviation,omitempty"`
	Description  optioner.Option[string]             `json:"description,omitempty"`
	Format       optioner.Option[PrefixSuffixFormat] `json:"format,omitempty"`
}

func (t Suffix) Validate(groups ...string) error {
	return nil
}

func (t Suffix) String() string {
	var s string

	format := t.Format.OrElse(Abbreviation)

	if format == Abbreviation && len(t.Abbreviation.Get()) > 0 {
		s = t.Abbreviation.Get()[0]
	} else {
		s = t.Text
	}

	return s
}
