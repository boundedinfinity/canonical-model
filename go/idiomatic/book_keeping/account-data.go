package book_keeping

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

var StandardAccounts = []AccountModel{
	{
		Id:   ider.MustParse("7620081B-D5DD-484C-BA8F-1BA0C1005914"),
		Name: "Assets",
	},
	{
		Id:   ider.MustParse("F22B3C38-D1EA-4BEB-8FD8-E46109E06434"),
		Name: "Liabilities",
	},
	{
		Id:   ider.MustParse("6C245CEC-9541-48F3-B4F8-F645DAA2B6FF"),
		Name: "Equity",
	},
	{
		Id:   ider.MustParse("E22BD1AB-AEAB-45FC-80E1-57641D952DBC"),
		Name: "Income",
	},
	{
		Id:   ider.MustParse("A55F2BF5-18C1-4DB7-92BE-FBD26E4F4300"),
		Name: "Expenses",
	},
}
