package model

type Language string

var (
	Languages = languages{
		Go:         "go",
		Typescript: "typescript",
	}
)

type languages struct {
	Go         Language
	Typescript Language
}
