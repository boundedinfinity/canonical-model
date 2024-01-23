// Reference
//  - https://en.wikipedia.org/wiki/ZIP_Code
//  - https://www.smarty.com/articles/zip-4-code

package mailing_address

import (
	"fmt"

	"github.com/boundedinfinity/schema/idiomatic/id"
)

type Zip struct {
	Id                       id.Id `json:"id,omitempty"`
	NationalArea             int   `json:"nattional-area,omitempty"`
	SectionalCenterFacility  int   `json:"sectional-center-facility,omitempty"`
	PostOfficeOrDeliveryArea int   `json:"post-office-or-delivery-area,omitempty"`
	Plus4Code                int   `json:"plus-4-code,omitempty"`
}

func (t Zip) String() string {
	s := fmt.Sprintf("%d%d%d", t.NationalArea, t.SectionalCenterFacility, t.PostOfficeOrDeliveryArea)

	if t.Is9Digit() {
		s += fmt.Sprintf("-%d", t.Plus4Code)
	}

	return s
}

func (t Zip) Is9Digit() bool {
	if t.Plus4Code != 0 {
		return true
	}

	return false
}
