package people

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/schema/idiomatic/internal"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Prefix struct {
	internal.WithIdDescFormat[PrefixFormat]
	Text         string   `json:"text,omitempty"`
	Abbreviation []string `json:"abbreviation,omitempty"`
}

func (t Prefix) String() string {
	var s string
	format, _ := slicer.FirstNotZero(t.Format, PrefixFormats.Abbreviation)

	if (format == PrefixFormats.Abbreviation) && len(t.Abbreviation) > 0 {
		s = t.Abbreviation[0]
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

// type prefixFn func(*Prefix) error

// type prefixBuilder struct {
// 	fns []prefixFn
// }

// func BuildPrefix() *prefixBuilder {
// 	return &prefixBuilder{
// 		fns: make([]prefixFn, 0),
// 	}
// }

// func (t *prefixBuilder) Build() (Prefix, error) {
// 	var v Prefix

// 	for _, fn := range t.fns {
// 		if err := fn(&v); err != nil {
// 			return v, err
// 		}
// 	}

// 	return v, nil
// }

// func (t *prefixBuilder) Must() Prefix {
// 	v, err := t.Build()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return v
// }

// func (t *prefixBuilder) Id(v string) *prefixBuilder {
// 	t.fns = append(t.fns, func(p *Prefix) error {
// 		id, err := uuid.Parse(v)

// 		if err != nil {
// 			return err
// 		}

// 		p.Id = id
// 		return nil
// 	})

// 	return t
// }

// func (t *prefixBuilder) Text(v string) *prefixBuilder {
// 	t.fns = append(t.fns, func(p *Prefix) error {
// 		p.Text = v
// 		return nil
// 	})

// 	return t
// }

// func (t *prefixBuilder) Abbreviation(v string) error {
// 	t.fns = append(t.fns, func(p *Prefix) error {
// 		p.Abbreviation = append(p.Abbreviation, v)
// 		return nil
// 	})

// 	return nil
// }

// func (t *prefixBuilder) Description(v string) *prefixBuilder {
// 	t.fns = append(t.fns, func(p *Prefix) error {
// 		p.Description = v
// 		return nil
// 	})

// 	return t
// }

// func (t *prefixBuilder) Format(v PrefixFormat) *prefixBuilder {
// 	t.fns = append(t.fns, func(p *Prefix) error {
// 		p.Format = v
// 		return nil
// 	})

// 	return t
// }
