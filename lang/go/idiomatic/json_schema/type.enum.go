package json_schema

//go:generate enumer -path=./type.enum.go

type JsonSchemaType string

type jsonSchemaTypes struct {
	String  JsonSchemaType `enum:"string"`
	Number  JsonSchemaType `enum:"number"`
	Integer JsonSchemaType `enum:"integer"`
	Object  JsonSchemaType `enum:"object"`
	Ref     JsonSchemaType `enum:"ref"`
}
