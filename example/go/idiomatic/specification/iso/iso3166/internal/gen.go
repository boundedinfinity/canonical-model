package main

// https://www.iso.org/iso-3166-country-codes.html
// https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
// https://en.wikipedia.org/wiki/ISO_3166-1
// https://github.com/PuerkitoBio/goquery?tab=readme-ov-file#examples

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"net/http"
	gurl "net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/specification/iso/iso3166"
)

func main() {
	fmt.Println("Generating ISO 3166 : Country Codes")
	url := "https://en.wikipedia.org/wiki/ISO_3166-1"
	jsonPath := "iso-3166.json"
	genPath := "../iso-3166.gen.go"

	_, err := os.Stat(jsonPath)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	var infos iso_3166

	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Downloading data from %v\n", url)

		infos = iso_3166{
			GeneratedAt: rfc3339date.NewDateTime(time.Now()),
			References: []string{
				"https://www.iso.org/iso-3166-country-codes.html",
				"https://en.wikipedia.org/wiki/ISO_3166-1",
			},
		}

		var dataM dataMatcher
		if err := readJson("matchers.json", &dataM); err != nil {
			panic(err)
		}

		if err := get_3166_1(dataM, &infos); err != nil {
			panic(err)
		}

		if err := writeJson(jsonPath, infos); err != nil {
			panic(err)
		}
	} else {
		bs, err := os.ReadFile(jsonPath)

		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(bs, &infos); err != nil {
			panic(err)
		}
	}

	if err := gen(genPath, infos, false); err != nil {
		panic(err)
	}
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Generate
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var content = `
package iso3166

var Iso3166s = iso3166s{}

type iso3166s struct {
	records  	[]Iso3166_1
	lc  	 	[]Iso3166_1
	alpha2s  	[]Iso3166_1_Alpha2
	alpha3s  	[]Iso3166_1_Alpha3
	numerics 	[]Iso3166_1_Numeric
	list2_codes []Iso3166_2_Code

{{- range .Records }}
	{{ .Alpha2 }} Iso3166_1_Alpha2
{{- end }}

{{- range .Records }}
	{{ .Alpha3 }} Iso3166_1_Alpha3
{{- end }}

{{- range .Records }}
	N_{{ .Numeric }} Iso3166_1_Numeric
{{- end }}

{{- range .Records }}
{{- range .Subdivisions }}
	{{ kebab2snake .Code }} Iso3166_2
{{- end }}
{{- end }}
}

func init() {
{{- range .Records }}
	Iso3166s.{{ .Alpha2 }} = Iso3166_1_Alpha2("{{ .Alpha2 }}")
{{- end }}

{{ range .Records }}
	Iso3166s.{{ .Alpha3 }} = Iso3166_1_Alpha3("{{ .Alpha3 }}")
{{- end }}

{{ range .Records }}
	Iso3166s.N_{{ .Numeric }} = Iso3166_1_Numeric("{{ .Numeric }}")
{{- end }}

	Iso3166s.alpha2s = []Iso3166_1_Alpha2{
{{- range .Records }}
		Iso3166s.{{ .Alpha2 }},
{{- end }}
	}

	Iso3166s.alpha3s = []Iso3166_1_Alpha3{
{{- range .Records }}
		Iso3166s.{{ .Alpha3 }},
{{- end }}
	}

	Iso3166s.numerics = []Iso3166_1_Numeric{
{{- range .Records }}
		Iso3166s.N_{{ .Numeric }},
{{- end }}
	}

	Iso3166s.records = []Iso3166_1{
{{- range .Records }}
		{
			Name: "{{ .Name }}",
			Alpha2: "{{ .Alpha2 }}",
			Alpha3: "{{ .Alpha3 }}",
			Numeric: "{{ .Numeric }}",
			Independent: "{{ .Independent }}",
			Reference: "{{ .Reference }}",
			Subdivisions: []Iso3166_2{
{{- range .Subdivisions }}
				{
					Name: "{{ .Name }}",
					Code: "{{ .Code }}",
					Reference: "{{ .Reference }}",
				},
{{- end }}
			},
		},
{{- end }}
	}

	Iso3166s.lc = []Iso3166_1{
{{- range .Records }}
		{
			Name: "{{ lower .Name }}",
			Alpha2: "{{ lower .Alpha2 }}",
			Alpha3: "{{ lower .Alpha3 }}",
			Numeric: "{{ lower .Numeric }}",
			Independent: "{{ lower .Independent }}",
			Reference: "{{ .Reference }}",
			Subdivisions: []Iso3166_2{
{{- range .Subdivisions }}
				{
					Name: "{{ lower .Name }}",
					Code: "{{ lower .Code }}",
					Reference: "{{ .Reference }}",
				},
{{- end }}
			},
		},
{{- end }}
	}
}
`

type templateData struct {
	Records  []iso3166.Iso3166_1
	Alpha2   []string
	Alpha3   []string
	Numerics []string
}

var (
	errColumnNotFound  = errors.New("column not found")
	errColumnNotFoundf = func(n string) error { return fmt.Errorf("%w, %s", errColumnNotFound, n) }
)

func gen(path string, jsonData iso_3166, formatSource bool) error {
	templateData := templateData{
		Records: jsonData.Data,
	}

	fm := template.FuncMap{
		"lower":       func(s string) string { return strings.ToLower(s) },
		"kebab2snake": func(s string) string { return strings.ReplaceAll(s, "-", "_") },
	}

	tmpl, err := template.New("").Funcs(fm).Parse(content)

	if err != nil {
		panic(err)
	}

	var bs []byte
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, templateData); err != nil {
		return err
	}

	bs = buf.Bytes()

	if formatSource {
		bs, err = format.Source(bs)
		if err != nil {
			return err
		}
	}

	if err := os.WriteFile(path, bs, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Utilities
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func writeJson(path string, data any) error {
	bs, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	if err := os.WriteFile(path, bs, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

func readJson(path string, data any) error {
	bs, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(bs, data)
}

type iso_3166 struct {
	GeneratedAt rfc3339date.Rfc3339DateTime `json:"generated-at"`
	References  []string                    `json:"references,omitempty,omitzero"`
	Data        []iso3166.Iso3166_1         `json:"data,omitempty,omitzero"`
}

type dataMatcher struct {
	List1 pageMatcher   `json:"list-1"`
	List2 []pageMatcher `json:"list-2"`
}

type pageMatcher struct {
	Source string         `json:"source"`
	Tables []tableMatcher `json:"tables"`
}

type tableMatcher struct {
	Table        int             `json:"table"`
	FirstDataRow int             `json:"first-data-row"`
	Columns      []columnMatcher `json:"columns"`
}

type columnMatcher struct {
	Name   string `json:"name"`
	Column int    `json:"column"`
	Value  string `json:"value"`
}

func getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Golang_Canonical_Model/1.0")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bss := string(bs)
	reader := strings.NewReader(bss)
	doc, err := goquery.NewDocumentFromReader(reader)

	return doc, err
}

func normalize(t *goquery.Selection) string {
	s := t.Text()
	s = strings.TrimSpace(s)

	if strings.Contains(s, "[") && strings.Contains(s, "]") {
		i := strings.Index(s, "[")
		s = s[:i]
	}

	return s
}

func getTableSelection(doc *goquery.Document, tableM tableMatcher) *goquery.Selection {
	var selection *goquery.Selection
	next := true

	doc.Find("table").EachWithBreak(func(i int, table *goquery.Selection) bool {
		if tableM.Table != i || !next {
			return next
		}

		selection = table
		next = false

		return next
	})

	return selection

}

func getCellSelection(row *goquery.Selection, tableM tableMatcher, columnName string) *goquery.Selection {
	var columnM columnMatcher
	var cmok bool
	var selection *goquery.Selection

	if row == nil {
		return selection
	}

	for _, current := range tableM.Columns {
		if current.Name == columnName && current.Value == "" {
			columnM = current
			cmok = true
			break
		}
	}

	if !cmok {
		return selection
	}

	next := true

	row.Find("th,td").EachWithBreak(func(i int, td *goquery.Selection) bool {
		if columnM.Column != i || !next {
			return next
		}

		selection = td
		next = false

		return next
	})

	return selection
}

func getText(row *goquery.Selection, tableM tableMatcher, columnName string, value *string) bool {
	selection := getCellSelection(row, tableM, columnName)

	if selection != nil {
		*value = normalize(selection)
	} else {
		for _, current := range tableM.Columns {
			if current.Name == columnName {
				if current.Value != "" {
					*value = current.Value
				}
			}
		}
	}

	return *value != ""
}

func getLink(row *goquery.Selection, matcher tableMatcher, name string, value *string, baseUrl string) bool {
	selection := getCellSelection(row, matcher, name)

	if selection == nil {
		return false
	}

	selection.Find("a").Each(func(_ int, anchor *goquery.Selection) {
		if found, ok := anchor.Attr("href"); ok {
			if parsed, err := gurl.Parse(baseUrl); err == nil {
				parsed.Path = found
				encoded := parsed.String()
				if decoded, err := gurl.QueryUnescape(encoded); err == nil {
					*value = decoded
				}
			}
		}
	})

	return *value != ""
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// List 1
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func get_3166_1(dataM dataMatcher, data *iso_3166) error {
	var doc *goquery.Document
	var ferr error

	doc, ferr = getPage(dataM.List1.Source)
	if ferr != nil {
		return ferr
	}

	list1matcher := dataM.List1.Tables[0]
	table := getTableSelection(doc, list1matcher)

	if table == nil {
		return ferr
	}

	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		if i < list1matcher.FirstDataRow {
			return
		}

		info := iso3166.Iso3166_1{
			Source: dataM.List1.Source,
		}

		getText(row, list1matcher, "Name", &info.Name)
		getText(row, list1matcher, "Alpha2", &info.Alpha2)
		getText(row, list1matcher, "Alpha3", &info.Alpha3)
		getText(row, list1matcher, "Numeric", &info.Numeric)
		getText(row, list1matcher, "Independent", &info.Independent)
		getLink(row, list1matcher, "Reference", &info.Reference, info.Source)

		var list2Source string
		if ok := getLink(row, list1matcher, "list-2", &list2Source, info.Source); ok {
			fn := func(info2 iso3166.Iso3166_2) {
				if info2.Name != "" {
					info.Subdivisions = append(info.Subdivisions, info2)
				}
			}

			if err := get_3166_2(dataM, list2Source, fn); err != nil {
				ferr = err
				return
			}
		}

		if info.Name != "" {
			data.Data = append(data.Data, info)
		}
	})

	return ferr
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// List 2
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func get_3166_2(dataM dataMatcher, source string, fn func(iso3166.Iso3166_2)) error {
	doc, err := getPage(source)
	if err != nil {
		return err
	}

	var pageM pageMatcher
	var pmok bool

	for _, current := range dataM.List2 {
		if current.Source == source {
			pageM = current
			pmok = true
			break
		}
	}

	if !pmok {
		for _, current := range dataM.List2 {
			if current.Source == "default" {
				pageM = current
				pmok = true
				break
			}
		}
	}

	if !pmok {
		return nil
	}

	for _, tableM := range pageM.Tables {
		table := getTableSelection(doc, tableM)

		if table == nil {
			return nil
		}

		table.Find("tr").Each(func(i int, row *goquery.Selection) {
			if i < tableM.FirstDataRow {
				return
			}

			info2 := iso3166.Iso3166_2{
				Source: source,
			}

			getText(row, tableM, "Name", &info2.Name)
			getText(row, tableM, "Code", &info2.Code)
			getLink(row, tableM, "Reference", &info2.Reference, info2.Source)
			getText(row, tableM, "Category", &info2.Category)
			getText(row, tableM, "Parent", &info2.Parent)
			getText(row, tableM, "AlternateCode", &info2.AlternateCode)
			getText(row, tableM, "PreviousCode", &info2.PreviousCode)

			if fn != nil {
				fn(info2)
			}
		})
	}

	return nil
}
