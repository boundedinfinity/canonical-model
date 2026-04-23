package authentication

import "github.com/boundedinfinity/canonical_model/idiomatic/id"

// https://en.wikipedia.org/wiki/Pretty_Good_Privacy

type PgpKeypair struct {
	Id          id.Id  `json:"id,omitempty"`
	Public      string `json:"public,omitempty"`
	Private     string `json:"private,omitempty"`
	Description string `json:"description,omitempty"`
}
