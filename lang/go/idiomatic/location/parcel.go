// Reference
// - https://id.land/
// - https://retipster.com/find-your-vacant-land/
package location

import "github.com/boundedinfinity/schema/idiomatic/id"

type Parcel struct {
	Id     id.Id  `json:"id,omitempty"`
	Number string `json:"number,omitempty"`
}
