package people

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed prefix.data.yaml
var prefixYaml string

// func (t prefixes) Find(s string) (Prefix, bool) {
// 	fn := func(prefix Prefix) bool {
// 		return stringer.EqualIgnoreCase(prefix.Text, s) ||
// 			stringer.ContainsAnyIgnoreCase(s, prefix.Abbreviation...)
// 	}

// 	return slicer.FindFn(fn, t...)
// }

// func (t prefixes) MustFind(s string) Prefix {
// 	if found, ok := t.Find(s); !ok {
// 		panic(fmt.Errorf("can't find prefix %v", s))
// 	} else {
// 		return found
// 	}
// }

var (
	Prefixes          = []Prefix{}
	PrefixDescriptors = []PrefixDescriptor{}
)

func init() {
	if err := yaml.Unmarshal([]byte(prefixYaml), &PrefixDescriptors); err != nil {
		panic(err)
	}

	for _, prefix := range Prefixes {
		PrefixDescriptors = append(PrefixDescriptors, prefix.Descriptor)
		Prefixes = append(Prefixes, prefix)
	}
}
