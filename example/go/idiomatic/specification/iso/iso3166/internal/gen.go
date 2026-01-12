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

		if err := get_3166_1(url, &infos); err != nil {
			panic(err)
		}

		if err := write_iso_3166(infos, jsonPath); err != nil {
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

	if err := gen(genPath, infos); err != nil {
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

func gen(path string, jsonData iso_3166) error {
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

	bs, err = format.Source(bs)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, bs, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func write_iso_3166(data iso_3166, path string) error {
	bs, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	if err := os.WriteFile(path, bs, os.FileMode(0755)); err != nil {
		return err
	}

	return nil
}

type iso_3166 struct {
	GeneratedAt rfc3339date.Rfc3339DateTime `json:"generated-at"`
	References  []string                    `json:"references,omitempty,omitzero"`
	Data        []iso3166.Iso3166_1         `json:"data,omitempty,omitzero"`
}

type tableMatcher struct {
	Source  string          `json:"source"`
	Table   int             `json:"table"`
	Row     int             `json:"row"`
	Columns []columnMatcher `json:"columns"`
}

type columnMatcher struct {
	Name   string `json:"name"`
	Header string `json:"header"`
	Index  int    `json:"index"`
	ok     bool   `json:"-"`
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

func getText(doc *goquery.Document, matcher tableMatcher, name string) (string, bool) {
	selection := getSelection(doc, matcher, name)
	var value string

	if selection != nil {
		value = normalize(selection)
	}

	return value, value != ""
}

func getLink(url string, tr *goquery.Selection, value *string, matcher columnMatcher) {
	if matcher.ok {
		tr.Find("th,td").Each(func(i2 int, td *goquery.Selection) {
			if matcher.Index == i2 {
				td.Find("a").Each(func(_ int, anchor *goquery.Selection) {
					if found, ok := anchor.Attr("href"); ok {
						if parsed, err := gurl.Parse(url); err == nil {
							parsed.Path = found
							encoded := parsed.String()
							if decoded, err := gurl.QueryUnescape(encoded); err == nil {
								*value = decoded
							}
						}
					}
				})
			}
		})
	}
}

func findLink(url string, tr *goquery.Selection, value *string) {
	var found bool

	tr.Find("th,td").EachWithBreak(func(i2 int, td *goquery.Selection) bool {
		td.Find("a").EachWithBreak(func(_ int, anchor *goquery.Selection) bool {
			if href, ok := anchor.Attr("href"); ok {
				if parsed, err := gurl.Parse(url); err == nil {
					parsed.Path = href
					encoded := parsed.String()
					if decoded, err := gurl.QueryUnescape(encoded); err == nil {
						*value = decoded
						found = true
					}
				}
			}

			return !found
		})

		return !found
	})
}

func get_3166_2(matcher *tableMatcher, info *iso3166.Iso3166_1, info2 *iso3166.Iso3166_2) error {
	if info != nil && info2.Source == "" {
		return nil
	}

	doc, err := getPage(info2.Source)
	if err != nil {
		return err
	}

	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if i != 0 {
			return
		}

		code := getHeader(table, []string{"Code"})
		name := getHeader(table, []string{
			"Subdivision name (ja)",
			"Subdivision name in English",
			"Subdivision name (en)",
			"Subdivision",
		})
		// reference := getHeader(table, []string{"Subdivision name"})
		category := getHeader(table, []string{
			"Subdivision category",
			"Subd. cat.",
		})

		matcher.Columns = append(matcher.Columns, code)
		matcher.Columns = append(matcher.Columns, name)
		matcher.Columns = append(matcher.Columns, category)

		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				return
			}

			getText(tr, &info2.Code, code)
			// getLink(info2.Source, tr, &info2.Reference, reference)
			findLink(info2.Source, tr, &info2.Reference)
			getText(tr, &info2.Name, name)
			getText(tr, &info2.Category, category)

			ignores := func(s string) bool {
				items := []string{
					"conventional names",
					"BGN/PCGN",
					"Name",
				}

				for _, item := range items {
					if strings.Contains(s, item) {
						return true
					}
				}

				return false
			}

			if info2.Code != "" && !ignores(info2.Code) {
				if info2.Reference == "" {
					fmt.Printf("no reference for: %s\n", info2.Source)
				}

				info.Subdivisions = append(info.Subdivisions, *info2)
			}
		})
	})

	return nil
}

func getSelection(doc *goquery.Document, matcher tableMatcher, name string) *goquery.Selection {
	var found columnMatcher
	var ok bool
	var selection *goquery.Selection

	for _, columnMatcher := range matcher.Columns {
		if columnMatcher.Name == name {
			found = columnMatcher
			ok = true
			break
		}
	}

	if !ok {
		return selection
	}

	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if matcher.Table != i {
			return
		}

		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if matcher.Row != i {
				return
			}

			tr.Find("th,td").Each(func(i int, td *goquery.Selection) {
				if found.Index != i {
					return
				}

				selection = td
			})
		})
	})

	return selection
}

func get_3166_1(url string, data *iso_3166) error {
	list1matcher := tableMatcher{
		Source: "https://en.wikipedia.org/wiki/ISO_3166-1",
		Table:  1,
		Row:    1,
		Columns: []columnMatcher{
			{
				Name:  "Name",
				Index: 0,
			},
			{
				Name:  "Alpha2",
				Index: 1,
			},
			{
				Name:  "Alpha3",
				Index: 2,
			},
			{
				Name:  "Numeric",
				Index: 3,
			},
			{
				Name:  "Source",
				Index: 4,
			},
			{
				Name:  "Reference",
				Index: 0,
			},
			{
				Name:  "Independent",
				Index: 5,
			},
		},
	}

	matchers := []tableMatcher{}

	doc, ferr := getPage(url)
	if ferr != nil {
		return ferr
	}

	doc.Find("table:has(caption)").Each(func(_ int, table *goquery.Selection) {
		matcher := tableMatcher{
			Table:   0,
			Row:     1,
			Columns: []columnMatcher{},
		}

		name := getHeader(table, []string{"short name"})
		alpha2 := getHeader(table, []string{"Alpha-2 code"})
		alpha3 := getHeader(table, []string{"Alpha-3 code"})
		numeric := getHeader(table, []string{"Numeric code"})
		list2 := getHeader(table, []string{"Link to"})
		independent := getHeader(table, []string{"Independent"})

		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				return
			}

			info := iso3166.Iso3166_1{
				Source: url,
			}

			getText(tr, &info.Name, name)
			// getLink(url, tr, &info.Reference, name)
			findLink(url, tr, &info.Reference)
			getText(tr, &info.Alpha2, alpha2)
			getText(tr, &info.Alpha3, alpha3)
			getText(tr, &info.Numeric, numeric)
			getText(tr, &info.Independent, independent)

			var list2Source string
			getLink(url, tr, &list2Source, list2)
			if list2Source != "" {
				matcher.Source = list2Source

				info2 := iso3166.Iso3166_2{Source: list2Source}
				if err := get_3166_2(&matcher, &info, &info2); err != nil {
					ferr = err
					return
				}

				matchers = append(matchers, matcher)
			}

			if info.Name != "" {
				data.Data = append(data.Data, info)
			}
		})
	})

	if bs, err := json.MarshalIndent(matchers, "", "    "); err != nil {
		return err
	} else {
		if err := os.WriteFile("matchers.json", bs, os.FileMode(0755)); err != nil {
			return err
		}
	}

	return ferr
}
