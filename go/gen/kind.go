package gen

type KindKind string

var (
	Kinds = kinds{
		Language:  "language",
		Object:    "object",
		Reference: "reference",
		String:    "string",
	}
)

type kinds struct {
	Language  KindKind
	Object    KindKind
	Reference KindKind
	String    KindKind
}

type LanguageKindKind string

var (
	LanguageKinds = languageKinds{
		Go:         "go",
		Typescript: "typescript",
	}
)

type languageKinds struct {
	Go         LanguageKindKind
	Typescript LanguageKindKind
}
