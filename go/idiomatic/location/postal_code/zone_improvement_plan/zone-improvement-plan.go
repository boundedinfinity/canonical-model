package zone_improvement_plan

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/postal_code"
)

// https://en.wikipedia.org/wiki/ZIP_Code

var _ postal_code.PostalCode = &ZoneImprovementPlan{}

type ZoneImprovementPlan struct {
	Id        ider.Id `json:"id"`
	Code      string  `json:"code"`
	Plus4Code string  `json:"plus4-code"`
	Type      Type    `json:"type"`
}

func (this *ZoneImprovementPlan) GetId() ider.Id {
	return this.Id
}

func (this *ZoneImprovementPlan) GetKind() postal_code.Kind {
	return postal_code.Kinds.ZoneImprovementPlan
}

func (this *ZoneImprovementPlan) String() string {
	return this.Code + "-" + this.Plus4Code
}
