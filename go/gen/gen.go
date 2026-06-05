package gen

import (
	"strings"

	"github.com/boundedinfinity/canonical-model/go/gen/model"
	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	errKindFunc = errorer.Func(model.ErrKind)
)

type generator struct {
}

func (this generator) Validate(config model.Config) error {
	switch config.Language {
	default:
		errKindFunc("config.language not supported: %s", config.Language)
	}

	for i, k := range config.Kinds {
		switch k.(type) {
		default:
			errKindFunc("config.kinds[%d] not supported: %s", i, k.GetKind())
		}
	}

	return nil
}

func (this generator) Generate() (string, error) {
	var sb strings.Builder

	return sb.String(), nil
}
