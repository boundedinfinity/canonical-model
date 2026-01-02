package main

// https://www.iso.org/iso-3166-country-codes.html
// https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
// https://en.wikipedia.org/wiki/ISO_3166-1
// https://github.com/PuerkitoBio/goquery?tab=readme-ov-file#examples

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	gurl "net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/rfc3339date"
)

func main() {
	fmt.Println("Generating ISO 3166 : Country Codes")
	url := "https://en.wikipedia.org/wiki/ISO_3166-1"

	infos := iso_3166{
		GeneratedAt: rfc3339date.NewDateTime(time.Now()),
		References: []string{
			"https://www.iso.org/iso-3166-country-codes.html",
			"https://en.wikipedia.org/wiki/ISO_3166-1",
		},
	}

	if err := get_3166_1(url, &infos); err != nil {
		panic(err)
	}

	if err := write_iso_3166(infos, "../iso-3166.json"); err != nil {
		panic(err)
	}
}

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
	Data        []iso_3166_1                `json:"data,omitempty,omitzero"`
}

type iso_3166_1 struct {
	Source       string       `json:"source,omitempty,omitzero"`
	Name         string       `json:"name,omitempty,omitzero"`
	Alpha2       string       `json:"alpha-2,omitempty,omitzero"`
	Alpha3       string       `json:"alpha-3,omitempty,omitzero"`
	Numeric      string       `json:"numeric,omitempty,omitzero"`
	Independent  string       `json:"independent,omitempty,omitzero"`
	Reference    string       `json:"reference,omitempty,omitzero"`
	Subdivisions []iso_3166_2 `json:"subdivisions,omitempty,omitzero"`
}

type iso_3166_2 struct {
	Source    string `json:"source,omitempty,omitzero"`
	Name      string `json:"name,omitempty,omitzero"`
	Code      string `json:"code,omitempty,omitzero"`
	Category  string `json:"cagetory,omitempty,omitzero"`
	Reference string `json:"reference,omitempty,omitzero"`
}

type columnMatcher struct {
	name   string
	header string
	index  int
	ok     bool
}

func getPage(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

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

func getHeader(table *goquery.Selection, names []string) columnMatcher {
	var matcher columnMatcher
	var headers []string

	table.Find("tr th").Each(func(i2 int, th *goquery.Selection) {
		headers = append(headers, normalize(th))
	})

FOUND:
	for _, name := range names {
		for i2, header := range headers {
			if stringer.Contains(header, name) {
				matcher.name = name
				matcher.header = header
				matcher.index = i2
				matcher.ok = true
				break FOUND
			}
		}
	}

	return matcher
}

func getText(s *goquery.Selection, value *string, matcher columnMatcher) {
	if matcher.ok {
		s.Find("th,td").Each(func(i2 int, td *goquery.Selection) {
			if matcher.index == i2 {
				*value = normalize(td)
			}
		})
	}
}

func getLink(url string, tr *goquery.Selection, value *string, matcher columnMatcher) {
	if matcher.ok {
		tr.Find("th,td").Each(func(i2 int, td *goquery.Selection) {
			if matcher.index == i2 {
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

func get_3166_2(info *iso_3166_1, info2 *iso_3166_2) error {
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
			"Subdivision name in English",
			"Subdivision name (en)",
			"Subdivision",
		})
		// reference := getHeader(table, []string{"Subdivision name"})
		category := getHeader(table, []string{
			"Subdivision category",
			"Subd. cat.",
		})

		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				return
			}

			getText(tr, &info2.Code, code)
			// getLink(info2.Source, tr, &info2.Reference, reference)
			findLink(info2.Source, tr, &info2.Reference)
			getText(tr, &info2.Name, name)
			getText(tr, &info2.Category, category)

			if info2.Code != "" {
				if info2.Reference == "" {
					fmt.Printf("no reference for: %s\n", info2.Source)
				}

				info.Subdivisions = append(info.Subdivisions, *info2)
			}
		})
	})

	return nil
}

func get_3166_1(url string, data *iso_3166) error {
	doc, ferr := getPage(url)
	if ferr != nil {
		return ferr
	}

	doc.Find("table:has(caption)").Each(func(_ int, table *goquery.Selection) {
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

			info := iso_3166_1{
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
				info2 := iso_3166_2{Source: list2Source}
				if err := get_3166_2(&info, &info2); err != nil {
					ferr = err
					return
				}
			}

			if info.Name != "" {
				data.Data = append(data.Data, info)
			}
		})
	})

	return ferr
}
