package business

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type RoleModel struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}
