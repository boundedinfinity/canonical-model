package passport

type Kind string

type kinds struct {
	Unknown      Kind
	UnitedStates Kind
}

var Kinds = kinds{
	Unknown:      "government.passport.unknown",
	UnitedStates: "government.passport.united-states",
}
