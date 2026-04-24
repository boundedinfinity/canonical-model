package authentication

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

type UsernamePassword struct {
	Id          ider.Id `json:"id,omitempty"`
	Username    string  `json:"username,omitempty"`
	Password    string  `json:"password,omitempty"`
	Description string  `json:"description,omitempty"`
}
