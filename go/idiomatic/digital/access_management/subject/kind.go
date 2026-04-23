package subject

type Kind string

type kinds struct {
	Unknown Kind
	Person  Kind
}

var Kinds = kinds{
	Unknown: "digital.access-management.subject.unknown",
	Person:  "digital.access-management.subject.person",
}
