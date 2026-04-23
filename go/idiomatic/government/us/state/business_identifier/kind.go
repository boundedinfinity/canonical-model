package business_identifier

type Kind string

type kinds struct {
	Unknown     Kind
	Colorado    Kind
	Iowa        Kind
	Missouri    Kind
	Mississippi Kind
}

var Kinds = kinds{
	Unknown:     "government.us.state.business-identifier.unknown",
	Colorado:    "government.us.state.business-identifier.colorado",
	Iowa:        "government.us.state.business-identifier.iowa",
	Missouri:    "government.us.state.business-identifier.missouri",
	Mississippi: "government.us.state.business-identifier.mississippi",
}
