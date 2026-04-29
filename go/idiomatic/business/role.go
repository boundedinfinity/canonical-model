package business

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type RoleModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}
