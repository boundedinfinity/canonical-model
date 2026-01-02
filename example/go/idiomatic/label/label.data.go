package label

import (
	"encoding/json"
)

var (
	LabelNames *LabelManager

	labelJson = `
		[
			{
				"kind": "label",
				"id": "B12E7628-35B0-4AE5-B2D9-0F34F4828229",
				"abbreviation": "dob",
				"name": "Date of Birth"
			},
			{
				"kind": "label",
				"id": "453F12F8-90C3-487C-BF96-53E281ABE143",
				"abbreviation": "dod",
				"name": "Date of Death"
			},
			{
				"kind": "label",
				"id": "9102C190-4093-4C73-BFE5-1FBA1CA41DE7",
				"name": "Expiration Date"
			},
			{
				"kind": "label",
				"id": "B64458ED-7B79-483E-9087-613BABB7A165",
				"name": "Formation Date"
			},
			{
				"kind": "label",
				"id": "FA64EF80-C3BC-4197-9A87-1FE1728A1100",
				"name": "Dissolution Date"
			},
			{
				"kind": "label",
				"id": "51ED6E00-D565-471F-87FD-479773C1382B",
				"name": "End of Life",
				"abbreviation": "eol"
			},
			{
				"kind": "label",
				"id": "51ED6E00-D565-471F-87FD-479773C1382B",
				"name": "Warranty Begin"
			},
			{
				"kind": "label",
				"id": "B6CB6E80-6179-42AA-B582-755FE640EA02",
				"name": "Warranty End"
			}
		]`
)

func init() {
	var labels []Label

	if err := json.Unmarshal([]byte(labelJson), &labels); err != nil {
		panic(err)
	}

	LabelNames = NewLabelManager(LabelManagerOptions{}, labels...)
}
