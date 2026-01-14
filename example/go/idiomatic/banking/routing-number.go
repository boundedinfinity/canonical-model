package banking

import (
	"errors"
	"fmt"
	"strconv"
)

// https://en.wikipedia.org/wiki/ABA_routing_transit_number
// https://tipalti.com/resources/learn/aba-number/
// https://routingnumber.aba.com/default1.aspx

type AmericanBankersAssociationRoutingNumber string

func (this AmericanBankersAssociationRoutingNumber) FederalReserveRoutingSymbol() FederalReserveRoutingSymbol {
	return FederalReserveRoutingSymbol(string(this[0:3]))
}

func (this AmericanBankersAssociationRoutingNumber) InstitutionIdentifier() AmericanBankersAssociationInstitutionIdentifier {
	return AmericanBankersAssociationInstitutionIdentifier(string(this[4:7]))
}

func (this AmericanBankersAssociationRoutingNumber) CheckDigit() string {
	return string(this[8])
}

var (
	ErrAmericanBankersAssociationInvalid  = errors.New("American Bankers Association check sum failure")
	errAmericanBankersAssociationInvalidf = func(format string, a ...any) error {
		message := fmt.Sprintf(format, a...)
		return fmt.Errorf("%w : %s", ErrAmericanBankersAssociationInvalid, message)
	}
)

func (this AmericanBankersAssociationRoutingNumber) CheckSum() error {
	// https://en.wikipedia.org/wiki/Routing_transit_number#Check_digit

	if len(this) != 9 {
		return errAmericanBankersAssociationInvalidf(
			"input must be 9 characters, not %d characters",
			len(this),
		)
	}

	district := this.FederalReserveRoutingSymbol().District()
	if _, ok := aba2District[district]; !ok {
		return errAmericanBankersAssociationInvalidf(
			"first 2 characters '%s' must be between 00-99", district,
		)
	}

	ws := [9]int{3, 7, 1, 3, 7, 1, 3, 7, 1}
	ns := [9]int{}
	sum := 0

	for i, c := range this {
		n, err := strconv.Atoi(string(c))

		if err != nil {
			return errAmericanBankersAssociationInvalidf(
				"input digit %d '%s' must be a number",
				i, this[i],
			)
		}

		ns[i] = n
		sum += ns[i] * ws[i]
	}

	mod := sum % 10

	if mod != 0 {
		return errAmericanBankersAssociationInvalidf(
			"modulus must be 0, result was %d", mod,
		)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// Federal Reserve Routing Symbol
// ////////////////////////////////////////////////////////////////////////////////////////////////

type FederalReserveRoutingSymbol string

func (this FederalReserveRoutingSymbol) Number() int {
	return mustAtoi(this)
}

func (this FederalReserveRoutingSymbol) District() string {
	return string(this[0:1])
}

func (this FederalReserveRoutingSymbol) Office() string {
	return string(this[2])
}

func (this FederalReserveRoutingSymbol) Area() string {
	return string(this[3])
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// American Bankers Association Institution Identifier
// ////////////////////////////////////////////////////////////////////////////////////////////////

type AmericanBankersAssociationInstitutionIdentifier string

// Some examples
// Bank Name										Address							City				State		Zip Code	Routing Number
// Bellco Credit Union (Head Office)				7600 E Orchard Rd, Ste 400, N	Greenwood Village	Colorado	80111-2522	302075018
// Bellco Credit Union (Branch)						1075 S Havana St				Aurora				Colorado	80012		302075018
// Wells Fargo Bank National Association (Branch)	3910 Lincoln Way				Ames				Iowa		50014		073000228
// Wells Fargo Bank National Association (Branch)	424 Main St						Ames				Iowa		50010		073000228

type AmericanBankersAssociationInstitutionLocation struct {
	Name    string
	Address string
	City    string
	State   string
	ZipCode string
}

// ////////////////////////////////////////////////////////////////////////////////////////////////
// utils
// ////////////////////////////////////////////////////////////////////////////////////////////////

func mustAtoi[T ~string | ~rune](input T) int {
	n, err := strconv.Atoi(string(input))

	if err != nil {
		panic(err)
	}

	return n
}

var (
	aba2District = map[string]AmericanBankersAssociationInfo{}
)

type AmericanBankersAssociationInfo struct {
	Name                     string
	FederalReserveDistrict   FederalReserveDistrict
	IsFederalReserveDistrict bool
}

func init() {
	between := func(rs [][]int, info AmericanBankersAssociationInfo) {
		for _, r := range rs {
			for i := r[0]; i <= r[1]; i++ {
				aba2District[fmt.Sprintf("%02d", i)] = info
			}
		}
	}

	between(
		[][]int{{0, 0}},
		AmericanBankersAssociationInfo{Name: "United States Government"},
	)

	between(
		[][]int{{13, 20}, {33, 39}, {40, 49}, {60, 60}, {73, 79}, {81, 82}, {93, 99}},
		AmericanBankersAssociationInfo{Name: "Reserved for future use"},
	)

	between(
		[][]int{{50, 59}},
		AmericanBankersAssociationInfo{Name: "Reserved for internal process control"},
	)

	between(
		[][]int{{61, 72}},
		AmericanBankersAssociationInfo{Name: "Electronic Transaction Identifiers"},
	)

	between(
		[][]int{{80, 80}},
		AmericanBankersAssociationInfo{Name: "Travelers Checks"},
	)

	for _, district := range FederalReserveDistricts.all {
		rs := [][]int{
			{district.Number, district.Number},
			{district.Number + 20, district.Number + 20},
			{district.Number + 60, district.Number + 60},
		}

		between(rs, AmericanBankersAssociationInfo{
			Name:                     district.Name,
			FederalReserveDistrict:   district,
			IsFederalReserveDistrict: true,
		})
	}
}
