// https://en.wikipedia.org/wiki/Universal_Product_Code
// https://en.wikipedia.org/wiki/GS1#Standards
// https://en.wikipedia.org/wiki/Luhn_algorithm
// https://en.wikipedia.org/wiki/Global_Trade_Item_Number
// https://en.wikipedia.org/wiki/List_of_GS1_country_codes

package id

type BarCodeType string

type barCodeTypes struct {
	UpcA  BarCodeType
	UpcB  BarCodeType
	UpcC  BarCodeType
	UpcD  BarCodeType
	UpcE  BarCodeType
	Upc2  BarCodeType
	Upc5  BarCodeType
	Ean2  BarCodeType
	Ean8  BarCodeType
	Ean13 BarCodeType
}

func (_ barCodeTypes) Parse(raw string) BarCode {
	return nil
}

var BarCodeTypes = barCodeTypes{
	UpcA:  "UPC-A",
	UpcB:  "UPC-B",
	UpcC:  "UPC-C",
	UpcD:  "UPC-D",
	UpcE:  "UPC-E",
	Upc2:  "UPC-2",
	Upc5:  "UPC-5",
	Ean2:  "EAN-2",
	Ean8:  "EAN-8",
	Ean13: "EAN-13",
}

type BarCode interface {
	Type() BarCodeType
	Manufacturer() string
	Product() string
	CheckDigit() int
}

// var _ BarCode = &UpcA{}

type UpcA struct {
}

type UpcB struct {
}

type UpcC struct {
}

type UpcD struct {
}

type UpcE struct {
}

type Upc2 struct {
}

type Upc5 struct {
}

type Ean2 struct {
}

type Ean8 struct {
}

type Ean13 struct {
}
