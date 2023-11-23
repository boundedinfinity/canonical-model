package phone

import "github.com/boundedinfinity/schema/idiomatic/id"

type Fax struct {
	Id     id.Id      `json:"id,omitempty"`
	Number NanpNumber `json:"number,omitempty"`
}

func (t Fax) String() string {
	return t.Number.String()
}
