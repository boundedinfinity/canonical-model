package model

type Kind string

var (
	Kinds = kinds{
		Array:       "array",
		Boolean:     "boolean",
		Enumeration: "enumeration",
		Float:       "float",
		Integer:     "integer",
		Language:    "language",
		Object:      "object",
		Reference:   "reference",
		String:      "string",
	}
)

type kinds struct {
	Array       Kind
	Boolean     Kind
	Enumeration Kind
	Float       Kind
	Integer     Kind
	Language    Kind
	Object      Kind
	Reference   Kind
	String      Kind
}
