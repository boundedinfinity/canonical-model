package people

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/schema/idiomatic/audit"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Suffix struct {
	Id           id.Id        `json:"id,omitempty"`
	Description  string       `json:"description,omitempty"`
	Text         string       `json:"text,omitempty"`
	Abbreviation []string     `json:"abbreviation,omitempty"`
	Format       SuffixFormat `json:"format,omitempty"`
	Audit        audit.Audit  `json:"audit,omitempty"`
}

func (t Suffix) Validate(groups ...string) error {
	return nil
}

func (t Suffix) String() string {
	var s string
	format, _ := slicer.FirstNotZero(t.Format, SuffixFormats.Abbreviation)

	if (format == SuffixFormats.Abbreviation) && len(t.Abbreviation) > 0 {
		s = t.Abbreviation[0]
	} else {
		s = t.Text
	}

	return s
}

// ///////////////////////////////////////////////////
// Builder
// ///////////////////////////////////////////////////

// type suffixFn func(*Suffix) error

// type suffixBuilder struct {
// 	fns []suffixFn
// }

// func BuildSuffix() *suffixBuilder {
// 	return &suffixBuilder{
// 		fns: make([]suffixFn, 0),
// 	}
// }

// func (t *suffixBuilder) Build() (Suffix, error) {
// 	var v Suffix

// 	for _, fn := range t.fns {
// 		if err := fn(&v); err != nil {
// 			return v, err
// 		}
// 	}

// 	return v, nil
// }

// func (t *suffixBuilder) Must() Suffix {
// 	v, err := t.Build()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return v
// }

// func (t *suffixBuilder) Id(v string) *suffixBuilder {
// 	t.fns = append(t.fns, func(p *Suffix) error {
// 		id, err := uuid.Parse(v)

// 		if err != nil {
// 			return err
// 		}

// 		p.Id = id
// 		return nil
// 	})

// 	return t
// }

// func (t *suffixBuilder) Text(v string) *suffixBuilder {
// 	t.fns = append(t.fns, func(p *Suffix) error {
// 		p.Text = v
// 		return nil
// 	})

// 	return t
// }

// func (t *suffixBuilder) Abbreviation(v string) error {
// 	t.fns = append(t.fns, func(p *Suffix) error {
// 		p.Abbreviation = append(p.Abbreviation, v)
// 		return nil
// 	})

// 	return nil
// }

// func (t *suffixBuilder) Description(v string) *suffixBuilder {
// 	t.fns = append(t.fns, func(p *Suffix) error {
// 		p.Description = v
// 		return nil
// 	})

// 	return t
// }

// func (t *suffixBuilder) Format(v SuffixFormat) *suffixBuilder {
// 	t.fns = append(t.fns, func(p *Suffix) error {
// 		p.Format = v
// 		return nil
// 	})

// 	return t
// }
