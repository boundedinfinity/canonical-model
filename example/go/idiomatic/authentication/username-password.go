package authentication

import "github.com/boundedinfinity/schema/idiomatic/id"

type UsernamePassword struct {
	Id          id.Id  `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Description string `json:"description,omitempty"`
}
