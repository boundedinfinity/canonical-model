package tax_identification_number

// https://www.ssa.gov/

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

var _ TaxIdentificationNumber = &SocialSecurityNumber{}

type SocialSecurityNumber struct {
	Id   ider.Id `json:"id"`
	Name string  `json:"name"`
}

func (this SocialSecurityNumber) GetName() string {
	return this.Name
}

func (_ SocialSecurityNumber) GetKind() Kind {
	return Kinds.SocialSecurityNumber
}

// import (
// 	"fmt"

// 	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
// 	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
// )

// // SSN,
// // https://www.ssa.gov/number-card
// // https://www.ssa.gov/policy/docs/ssb/v45n11/v45n11p29.pdf
// type SocialSecurityNumber struct {
// 	Id     ider.Id `json:"id,omitempty"`
// 	Area   int   `json:"area,omitempty"`
// 	Group  int   `json:"group,omitempty"`
// 	Serial int   `json:"serial,omitempty"`
// }

// func (t SocialSecurityNumber) Last4() string {
// 	return fmt.Sprintf("%d", t.Serial)
// }

// func (t SocialSecurityNumber) String() string {
// 	return fmt.Sprintf("%d-%d-%d", t.Area, t.Group, t.Serial)
// }

// func (t SocialSecurityNumber) Type() string {
// 	return "social-security-number"
// }

// var _ TaxIdentificationNumber = &SocialSecurityNumber{}

// ////////////////////////////////////////////////////////////
// // Companion
// ////////////////////////////////////////////////////////////

// var SocialSecurityNumbers = socialSecurityNumbers{}

// type socialSecurityNumbers struct{}

// func (t socialSecurityNumbers) Parse(input string) (SocialSecurityNumber, error) {
// 	var ssn SocialSecurityNumber

// 	return ssn, nil
// }

// func (t socialSecurityNumbers) MustParse(input string) SocialSecurityNumber {
// 	ssn, err := t.Parse(input)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return ssn
// }

// ////////////////////////////////////////////////////////////
// // Card
// ////////////////////////////////////////////////////////////

// type SocialSecurityCard struct {
// 	Id     ider.Id                `json:"id,omitempty"`
// 	Name   people.Name          `json:"name,omitempty"`
// 	Number SocialSecurityNumber `json:"number,omitempty"`
// }
