package idiomatic

//go:generate enumer -config=./banking/account-number-format.enum.yaml
//go:generate enumer -config=./people/name-format.enum.yaml
//go:generate enumer -config=./people/prefix-format.enum.yaml
//go:generate enumer -config=./people/suffix-format.enum.yaml
//go:generate enumer -config=./people/sex-type.enum.yaml
//go:generate enumer -config=./audit/status.enum.yaml
//go:generate enumer -config=./phone/at-command.enum.yaml
//go:generate enumer -config=./phone/nanp-phone-separator-format.enum.yaml
//go:generate enumer -config=./phone/nanp-phone-format.enum.yaml
//go:generate enumer -config=./physical_card/donor-status.enum.yaml
//go:generate enumer -config=./specification/json_schema/type.enum.yaml
//go:generate enumer -config=./vehicle/wmi-region.enum.yaml
//go:generate enumer -config=./location/state-name.enum.yaml
//go:generate enumer -config=./location/state-ansi.enum.yaml
//go:generate enumer -config=./location/state-iso.enum.yaml
//go:generate enumer -config=./banking/federal-reserve-district.enum.yaml
