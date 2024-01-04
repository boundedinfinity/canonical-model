package idiomatic

//go:generate enumer -config=./idiomatic/banking/account-number-format.enum.yaml
//go:generate enumer -config=./idiomatic/people/name-format.enum.yaml
//go:generate enumer -config=./idiomatic/people/prefix-format.enum.yaml
//go:generate enumer -config=./idiomatic/people/suffix-format.enum.yaml
//go:generate enumer -config=./idiomatic/people/sex-type.enum.yaml
//----go:generate enumer -config=./idiomatic/audit/status.enum.yaml
//go:generate enumer -config=./idiomatic/phone/at-command.enum.yaml
//go:generate enumer -config=./idiomatic/phone/nanp-phone-separator-format.enum.yaml
//go:generate enumer -config=./idiomatic/phone/nanp-phone-format.enum.yaml
//go:generate enumer -config=./idiomatic/physical_card/donor-status.enum.yaml
//go:generate enumer -config=./idiomatic/specification/json_schema/type.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/wmi-region.enum.yaml
//go:generate enumer -config=./idiomatic/location/state-name.enum.yaml
//go:generate enumer -config=./idiomatic/location/state-ansi.enum.yaml
//go:generate enumer -config=./idiomatic/location/state-iso.enum.yaml
//go:generate enumer -config=./idiomatic/banking/federal-reserve-district.enum.yaml
//go:generate enumer -config=./idiomatic/currency/currency-code.enum.yaml
//----go:generate enumer -config=./idiomatic/currency/currency-name.enum.yaml
//----go:generate enumer -config=./idiomatic/id/iso-639-lang.enum.yaml
//----go:generate enumer -config=./idiomatic/id/iso-639-1.enum.yaml
//----go:generate enumer -config=./idiomatic/id/iso-639-2-b.enum.yaml
//----go:generate enumer -config=./idiomatic/id/iso-639-2-t.enum.yaml
//----go:generate enumer -config=./idiomatic/id/iso-639-3.enum.yaml
//go:generate enumer -config=./idiomatic/hardware/connector-type.enum.yaml
//go:generate enumer -config=./idiomatic/finanical/interest-rate-type.enum.yaml

//go:generate enumer -config=./idiomatic/vehicle/body-style.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/body-size.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/make.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/feul-type.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/feature.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/drive-train.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/transmission-type.enum.yaml
//go:generate enumer -config=./idiomatic/vehicle/engine-type.enum.yaml

//go:generate enumer -config=./idiomatic/color/color.enum.yaml

//go:generate go run ./idiomatic/specification/iso/gen
//go:generate enumer -config=./idiomatic/location/country-alpha-2.enum.yaml
//go:generate enumer -config=./idiomatic/location/country-alpha-3.enum.yaml
