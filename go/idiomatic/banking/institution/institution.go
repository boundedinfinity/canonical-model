package institution

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/banking/routing_transit_number"
	"github.com/boundedinfinity/canonical_model/idiomatic/digital/bookmark"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

type InstitutionModel struct {
	Id                   ider.Id                                     `json:"id"`
	Name                 string                                      `json:"name"`
	RoutingTransitNumber routing_transit_number.RoutingTransitNumber `json:"routing-transit-number"`
	Site                 bookmark.Bookmark                           `json:"site"`
}
