module github.com/boundedinfinity/canonical-model/go

go 1.25.0

require (
	github.com/boundedinfinity/go-commoner v1.0.30
	github.com/boundedinfinity/rfc3339date v1.0.7
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.9.0
	github.com/tursodatabase/go-libsql v0.0.0-20240725130945-f44f2b84c8c8
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/libsql/sqlite-antlr4-parser v0.0.0-20240327125255-dbf53b6cbf06 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/boundedinfinity/go-commoner => ../../go-commoner
