package vcard

import "github.com/boundedinfinity/schema/idiomatic/label"

var VCardType = vCardType{}

type vCardType struct{}

func (t vCardType) FromLabel(labels ...label.ValueLabel) VCardTypeList {
	var list VCardTypeList

	for _, item := range labels {
		switch val := item.(type) {
		case label.NameOnlyLabel:
			list = append(list, val.Name.Name)
		case label.DateLabel:
			list = append(list, val.Name.Name)
		}
	}

	return list
}

type VCardTypeList []string
