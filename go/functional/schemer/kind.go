package schemer

type Kind string

var Kinds = kinds{
	Float:      "schemer.kind.float",
	Enumerated: "schemer.kind.enunerated",
	Integer:    "schemer.kind.integer",
	Object:     "schemer.kind.object",
	String:     "schemer.kind.string",
	Reference:  "schemer.kind.reference",
}

type kinds struct {
	Integer    Kind
	Float      Kind
	String     Kind
	Enumerated Kind
	Object     Kind
	Reference  Kind
}
