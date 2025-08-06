package people

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed prefix.data.yaml
var prefixYaml string

var (
	Prefixes = []Prefix{}
)

func init() {
	type dataContainer struct {
		Data []Prefix
	}

	var data dataContainer

	if err := yaml.Unmarshal([]byte(prefixYaml), &data); err != nil {
		panic(err)
	}

	Prefixes = append(Prefixes, data.Data...)
}
