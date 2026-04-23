package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	path := "../digit.gen.go"

	file := he1(os.Create(path))
	defer file.Close()

	m := map[string][]string{
		"2": {"A", "B", "C"},
		"3": {"D", "E", "F"},
		"4": {"G", "H", "I"},
		"5": {"J", "K", "L"},
		"6": {"M", "Q", "R", "S"},
		"7": {"P", "N", "O"},
		"8": {"T", "U", "V"},
		"9": {"W", "X", "Y", "Z"},
	}

	fm := template.FuncMap{
		"lower": func(s string) string { return strings.ToLower(s) },
		"upper": func(s string) string { return strings.ToUpper(s) },
	}

	tmpl := he1(template.New("").Funcs(fm).Parse(content))
	he0(tmpl.Execute(file, m))
}

var content = `
package phone

type digits struct {
	dm map[string]Digit
{{- range $k, $v := . }}
	n{{ lower $k }} Digit
{{- end }}
}

var Digits = digits{
{{- range $k, $v := . }}
	n{{ lower $k }}: Digit{Number: {{ $k }}, Mnemonic: []string{ {{- range $v }}"{{ . }}", {{ end -}}}},
{{- end }}
}

func (this digits) All() []Digit {
	return []Digit {
{{- range $k, $v := . }}
		this.n{{ lower $k }},
{{- end }}
	}
}

func init() {
	Digits.dm = map[string]Digit{
{{- range $k, $v := . }}
		"{{ $k }}": Digits.n{{ lower $k }},
{{- range $v }}
		"{{ upper . }}": Digits.n{{ lower $k }},
		"{{ lower . }}": Digits.n{{ lower $k }},
{{- end }}
{{- end }}
	}
}
`

func he0(err error) {
	if err != nil {
		panic(err)
	}
}

func he1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
