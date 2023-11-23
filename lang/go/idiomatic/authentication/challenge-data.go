package authentication

import "github.com/boundedinfinity/schema/idiomatic/id"

// https://financialdataexchange.org/common/Uploaded%20files/OFX%20files/OFX%20Banking%20Specification%20v2.3.pdf#p=62
// 2.5.4.4.1 Enumerated <MFAPHRASEID> Meanings

var (
	ChallengeQuestions = []ChallengeQuestion{
		{
			Id:   id.MustParse("4599B6F6-393D-4329-9862-04132F940262"),
			Code: "MFA1",
			Text: "City Of birth",
		},
		{
			Id:   id.MustParse("0821F522-B504-4F89-A840-4F8754850F16"),
			Code: "MFA2",
			Text: "Date of birth, formatted MM/DD/YYYY",
		},
		{
			Id:   id.MustParse("34c7fea7-6286-43ab-8a2c-72c94013691a"),
			Code: "MFA3",
			Text: "Debit card number",
		},
		{
			Id:   id.MustParse("7c545c64-56f7-4279-a9e8-e72a599750c9"),
			Code: "MFA4",
			Text: "Father's middle name",
		},
	}
)
