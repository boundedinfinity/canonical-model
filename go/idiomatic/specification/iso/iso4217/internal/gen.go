package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

// https://www.iso.org/iso-4217-currency-codes.html
// https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml

type xmlData struct {
	XMLName   xml.Name `xml:"ISO_4217"`
	Published string   `xml:"Pblshd,attr"`
	Table     xmlTable `xml:"CcyTbl"`
}

type xmlTable struct {
	XMLName xml.Name  `xml:"CcyTbl"`
	Items   []iso4217 `xml:"CcyNtry"`
}

type iso4217 struct {
	XMLName xml.Name `xml:"CcyNtry"`
	Country string   `xml:"CtryNm"`
	Name    string   `xml:"CcyNm"`
	Code    string   `xml:"Ccy"`
	Numeric string   `xml:"CcyNbr"`
	Units   string   `xml:"CcyMnrUnts"`
}

type templateData struct {
	Items    []iso4217
	Codes    []string
	Numerics []string
}

func main() {
	source := "https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml"
	path := "../iso-4217.gen.go"
	xmlPath := "list-one.xml"
	var data xmlData

	if err := getData(source, xmlPath, &data); err != nil {
		panic(data)
	}

	if err := gen(path, data); err != nil {
		panic(data)
	}
}

var content = `
package iso4217

type iso4217s struct {
	records []Iso4217
	lc  []Iso4217
	codes []Iso4217Code
	numerics []Iso4217Numeric

{{- range .Codes }}
	{{ . }} Iso4217Code
{{- end }}

{{- range .Numerics }}
	N_{{ . }} Iso4217Numeric
{{- end }}
}

func init() {
{{- range .Codes }}
	Iso4217s.{{ . }} = Iso4217Code("{{ . }}")
{{- end }}
{{ range .Numerics }}
	Iso4217s.N_{{ . }} = Iso4217Numeric("{{ . }}")
{{- end }}

	Iso4217s.codes = []Iso4217Code{
{{- range .Codes }}
		Iso4217s.{{ . }},
{{- end }}
	}

	Iso4217s.numerics = []Iso4217Numeric{
{{- range .Numerics }}
		Iso4217s.N_{{ . }},
{{- end }}
	}

	Iso4217s.records = []Iso4217{
{{- range .Items }}
		{
			Country: "{{ .Country }}",
			Name: "{{ .Name }}",
			Code: "{{ .Code }}",
			Numeric: "{{ .Numeric }}",
			Units: "{{ .Units }}",
		},
{{- end }}
	}

	Iso4217s.lc = []Iso4217{
{{- range .Items }}
		{
			Country: "{{ lower .Country }}",
			Name: "{{ lower .Name }}",
			Code: "{{ lower .Code }}",
			Numeric: "{{ lower .Numeric }}",
			Units: "{{ lower .Units }}",
		},
{{- end }}
	}
}
`

func gen(path string, xmlData xmlData) error {
	add := func(m map[string]bool, value string) {
		if _, ok := m[value]; !ok && value != "" {
			m[value] = true
		}
	}

	codes := map[string]bool{}
	numerics := map[string]bool{}

	for _, item := range xmlData.Table.Items {
		add(codes, item.Code)
		add(numerics, item.Numeric)
	}

	data := templateData{
		Items:    xmlData.Table.Items,
		Codes:    mapper.Keys(codes),
		Numerics: mapper.Keys(numerics),
	}

	fm := template.FuncMap{
		"lower": func(s string) string { return strings.ToLower(s) },
	}

	tmpl, err := template.New("").Funcs(fm).Parse(content)

	if err != nil {
		panic(err)
	}

	var bs []byte
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	bs = buf.Bytes()
	bs, err = format.Source(bs)

	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, bs, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

func getData(url string, path string, data *xmlData) error {
	_, err := os.Stat(path)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if errors.Is(err, os.ErrNotExist) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req.Header.Set("User-Agent", "Golang_Canonical_Model/1.0")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		file, err := os.Create(path)

		if err != nil {
			return err
		}

		defer file.Close()

		if _, err := io.Copy(file, res.Body); err != nil {
			return err
		}
	}

	bs, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	if err := xml.Unmarshal(bs, data); err != nil {
		return err
	}

	return nil
}
