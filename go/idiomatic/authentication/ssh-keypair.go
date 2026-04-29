package authentication

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

// https://www.ssh.com/academy

type SshKeypair struct {
	Id          ider.Id `json:"id,omitempty"`
	Public      string  `json:"public,omitempty"`
	Private     string  `json:"private,omitempty"`
	Description string  `json:"description,omitempty"`
	Bits        int     `json:"bits,omitempty"`
	PassPhrase  int     `json:"pass-phrase,omitempty"`
}
