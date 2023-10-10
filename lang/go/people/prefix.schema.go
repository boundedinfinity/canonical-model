package people

import (
	"github.com/boundedinfinity/optioner"
	"github.com/google/uuid"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Prefix struct {
	Id           optioner.Option[uuid.UUID]    `json:"id,omitempty"`
	Text         string                        `json:"text,omitempty"`
	Abbreviation optioner.Option[[]string]     `json:"abbreviation,omitempty"`
	Description  optioner.Option[string]       `json:"description,omitempty"`
	Format       optioner.Option[PrefixFormat] `json:"format,omitempty"`
}

func (t Prefix) String() string {
	var s string

	if (t.Format.Empty() || t.Format.Get() == PrefixFormats.Abbreviation) && len(t.Abbreviation.Get()) > 0 {
		s = t.Abbreviation.Get()[0]
	} else {
		s = t.Text
	}

	return s
}

func (t Prefix) Validate(groups ...string) error {
	return nil
}

// ///////////////////////////////////////////////////
// Builder
// ///////////////////////////////////////////////////

type prefixFn func(*Prefix) error

type prefixBuilder struct {
	fns []prefixFn
}

func BuildPrefix() *prefixBuilder {
	return &prefixBuilder{
		fns: make([]prefixFn, 0),
	}
}

func (t *prefixBuilder) Build() (Prefix, error) {
	var v Prefix

	for _, fn := range t.fns {
		if err := fn(&v); err != nil {
			return v, err
		}
	}

	return v, nil
}

func (t *prefixBuilder) Must() Prefix {
	v, err := t.Build()

	if err != nil {
		panic(err)
	}

	return v
}

func (t *prefixBuilder) Id(v string) *prefixBuilder {
	t.fns = append(t.fns, func(p *Prefix) error {
		id, err := uuid.Parse(v)

		if err != nil {
			return err
		}

		p.Id = optioner.Some(id)
		return nil
	})

	return t
}

func (t *prefixBuilder) Text(v string) *prefixBuilder {
	t.fns = append(t.fns, func(p *Prefix) error {
		p.Text = v
		return nil
	})

	return t
}

func (t *prefixBuilder) Abbreviation(v string) error {
	t.fns = append(t.fns, func(p *Prefix) error {
		if p.Abbreviation.Empty() {
			p.Abbreviation = optioner.Some(make([]string, 0))
		}

		p.Abbreviation = optioner.Some(append(p.Abbreviation.Get(), v))

		return nil
	})

	return nil
}

func (t *prefixBuilder) Description(v string) *prefixBuilder {
	t.fns = append(t.fns, func(p *Prefix) error {
		p.Description = optioner.Some(v)
		return nil
	})

	return t
}

func (t *prefixBuilder) Format(v PrefixFormat) *prefixBuilder {
	t.fns = append(t.fns, func(p *Prefix) error {
		p.Format = optioner.Some(v)
		return nil
	})

	return t
}
