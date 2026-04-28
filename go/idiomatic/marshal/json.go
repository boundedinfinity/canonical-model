package marshal

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func UnmarshalJsonFromSlice[K ~string](bytes []byte, valids []K, errFn func(format string, a ...any) error) (K, error) {
	var found K
	var s string
	err := json.Unmarshal(bytes, &s)

	if err != nil {
		return found, errFn("unmarshal error: %w", err)
	}

	var ok bool

	for _, value := range valids {
		if s == string(value) {
			found = value
			ok = true
			break
		}
	}

	if !ok {
		valid := stringer.Join(", ", valids...)
		return found, errFn("unknown: %s : must be one of %s", s, valid)
	}

	return found, nil
}
