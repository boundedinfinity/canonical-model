package zone_improvement_plan

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/location/postal_codes"
)

// https://en.wikipedia.org/wiki/ZIP_Code

var _ postal_codes.PostalCode = &ZoneImprovementPlan{}

type ZoneImprovementPlan struct {
	Id        ider.Id `json:"id"`
	Code      string  `json:"code"`
	Plus4Code string  `json:"plus4-code"`
	Type      Type    `json:"type"`
}

func (this *ZoneImprovementPlan) GetId() ider.Id {
	return this.Id
}

func (this *ZoneImprovementPlan) GetKind() postal_codes.Kind {
	return postal_codes.Kinds.ZoneImprovementPlan
}

func (this *ZoneImprovementPlan) String() string {
	return this.Code + "-" + this.Plus4Code
}
