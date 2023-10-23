module github.com/boundedinfinity/schema

go 1.18

require (
	github.com/boundedinfinity/enumer v1.0.10
	github.com/boundedinfinity/go-commoner v1.0.25
	github.com/boundedinfinity/go-mimetyper v1.0.17
	github.com/boundedinfinity/rfc3339date v1.0.4
	github.com/google/uuid v1.3.1
	github.com/stretchr/testify v1.8.4
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/boundedinfinity/go-commoner => ../../../go-commoner/
