package json_schema

//go:generate enumer -path=./string-format.enum.go

type JsonStringFormat string

// https://json-schema.org/understanding-json-schema/reference/string#format

type jsonStringFormats struct {
	DateTime            JsonSchemaType `enum:"date-time"`
	Time                JsonSchemaType `enum:"time"`
	Date                JsonSchemaType `enum:"date"`
	Duration            JsonSchemaType `enum:"duration"`
	Email               JsonSchemaType `enum:"email"`
	IdnEmail            JsonSchemaType `enum:"idn-email"`
	Hostmane            JsonSchemaType `enum:"hostname"`
	IdnHostmane         JsonSchemaType `enum:"idn-hostname"`
	Ipv4                JsonSchemaType `enum:"ipv4"`
	Ipv6                JsonSchemaType `enum:"ipv6"`
	Uuid                JsonSchemaType `enum:"uuid"`
	Uri                 JsonSchemaType `enum:"uri"`
	UriReference        JsonSchemaType `enum:"uri-reference"`
	Iri                 JsonSchemaType `enum:"iri"`
	IriReference        JsonSchemaType `enum:"iri-reference"`
	UriTemplate         JsonSchemaType `enum:"uri-template"`
	JsonPointer         JsonSchemaType `enum:"json-pointer"`
	RelativeJsonPointer JsonSchemaType `enum:"relative-json-pointer"`
	Regex               JsonSchemaType `enum:"regex"`
}
