package zone_improvement_plan

// https://en.wikipedia.org/wiki/ZIP_Code#Types

type Type string

type types struct {
	Unknown       Type
	Unique        Type
	PostOfficeBox Type
	Military      Type
	Standard      Type
}

var Types = types{
	Unknown:       "location.postal-codes.zone-improvement-plan.unknown",
	Unique:        "location.postal-codes.zone-improvement-plan.unique",
	PostOfficeBox: "location.postal-codes.zone-improvement-plan.post-office-box",
	Military:      "location.postal-codes.zone-improvement-plan.military",
	Standard:      "location.postal-codes.zone-improvement-plan.standard",
}
