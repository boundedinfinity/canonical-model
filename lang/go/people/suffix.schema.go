package people

import (
	"github.com/boundedinfinity/optioner"
	"github.com/google/uuid"
)

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

// ///////////////////////////////////////////////////
// Builder
// ///////////////////////////////////////////////////

type suffixFn func(*Suffix) error

type suffixBuilder struct {
	fns []suffixFn
}

func BuildSuffix() *suffixBuilder {
	return &suffixBuilder{
		fns: make([]suffixFn, 0),
	}
}

func (t *suffixBuilder) Build() (Suffix, error) {
	var v Suffix

	for _, fn := range t.fns {
		if err := fn(&v); err != nil {
			return v, err
		}
	}

	return v, nil
}

func (t *suffixBuilder) Must() Suffix {
	v, err := t.Build()

	if err != nil {
		panic(err)
	}

	return v
}

func (t *suffixBuilder) Id(v string) *suffixBuilder {
	t.fns = append(t.fns, func(p *Suffix) error {
		id, err := uuid.Parse(v)

		if err != nil {
			return err
		}

		p.Id = optioner.Some(id)
		return nil
	})

	return t
}

func (t *suffixBuilder) Text(v string) *suffixBuilder {
	t.fns = append(t.fns, func(p *Suffix) error {
		p.Text = v
		return nil
	})

	return t
}

func (t *suffixBuilder) Abbreviation(v string) error {
	t.fns = append(t.fns, func(p *Suffix) error {
		if p.Abbreviation.Empty() {
			p.Abbreviation = optioner.Some(make([]string, 0))
		}

		p.Abbreviation = optioner.Some(append(p.Abbreviation.Get(), v))

		return nil
	})

	return nil
}

func (t *suffixBuilder) Description(v string) *suffixBuilder {
	t.fns = append(t.fns, func(p *Suffix) error {
		p.Description = optioner.Some(v)
		return nil
	})

	return t
}

func (t *suffixBuilder) Format(v PrefixSuffixFormat) *suffixBuilder {
	t.fns = append(t.fns, func(p *Suffix) error {
		p.Format = optioner.Some(v)
		return nil
	})

	return t
}
