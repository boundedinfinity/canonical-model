package jurisdiction

type Kind string

type kinds struct {
	Unknown       Kind
	Federal       Kind
	State         Kind
	County        Kind
	City          Kind
	Judicial      Kind
	Senatorial    Kind
	Congressional Kind
	SchoolZone    Kind
}

var Kinds = kinds{
	Unknown:       "government.us.jurisdiction.unknown",
	Federal:       "government.us.jurisdiction.federal",
	State:         "government.us.jurisdiction.state",
	County:        "government.us.jurisdiction.county",
	City:          "government.us.jurisdiction.city",
	Judicial:      "government.us.jurisdiction.judicial",
	Senatorial:    "government.us.jurisdiction.senatorial",
	Congressional: "government.us.jurisdiction.congressional",
	SchoolZone:    "government.us.jurisdiction.school-zone",
}
