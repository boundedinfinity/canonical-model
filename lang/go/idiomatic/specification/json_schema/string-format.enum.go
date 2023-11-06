package json_schema

//go:generate enumer -path=./string-format.enum.go

type JsonStringFormat string

// https://json-schema.org/understanding-json-schema/reference/string#format

type jsonStringFormats struct {
	DateTime            JsonStringFormat `enum:"date-time"`
	Time                JsonStringFormat `enum:"time"`
	Date                JsonStringFormat `enum:"date"`
	Duration            JsonStringFormat `enum:"duration"`
	Email               JsonStringFormat `enum:"email"`
	IdnEmail            JsonStringFormat `enum:"idn-email"`
	Hostmane            JsonStringFormat `enum:"hostname"`
	IdnHostmane         JsonStringFormat `enum:"idn-hostname"`
	Ipv4                JsonStringFormat `enum:"ipv4"`
	Ipv6                JsonStringFormat `enum:"ipv6"`
	Uuid                JsonStringFormat `enum:"uuid"`
	Uri                 JsonStringFormat `enum:"uri"`
	UriReference        JsonStringFormat `enum:"uri-reference"`
	Iri                 JsonStringFormat `enum:"iri"`
	IriReference        JsonStringFormat `enum:"iri-reference"`
	UriTemplate         JsonStringFormat `enum:"uri-template"`
	JsonPointer         JsonStringFormat `enum:"json-pointer"`
	RelativeJsonPointer JsonStringFormat `enum:"relative-json-pointer"`
	Regex               JsonStringFormat `enum:"regex"`
}
