package authentication

import "github.com/boundedinfinity/schema/idiomatic/id"

type SshKeypair struct {
	Id          id.Id  `json:"id,omitempty"`
	Public      string `json:"public,omitempty"`
	Private     string `json:"private,omitempty"`
	Description string `json:"description,omitempty"`
}
