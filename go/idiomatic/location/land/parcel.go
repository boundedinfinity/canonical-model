// Reference
// - https://id.land/
// - https://retipster.com/find-your-vacant-land/
package land

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type Parcel struct {
	Id     ider.Id `json:"id,omitempty"`
	Number string  `json:"number,omitempty"`
}
