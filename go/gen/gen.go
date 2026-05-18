package gen

import "strings"

type Generator struct {
}

func (this Generator) Validate(config Config) error {
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

func (this Generator) Generate() (string, error) {
	var sb strings.Builder

	return sb.String(), nil
}
