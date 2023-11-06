package authentication

import "github.com/boundedinfinity/schema/idiomatic/id"

// https://financialdataexchange.org/common/Uploaded%20files/OFX%20files/OFX%20Banking%20Specification%20v2.3.pdf#p=62

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
	}
)
