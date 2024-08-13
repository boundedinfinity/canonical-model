module github.com/boundedinfinity/schema

go 1.21

require (
	github.com/boundedinfinity/enumer v1.0.21
	github.com/boundedinfinity/go-commoner v1.0.30
	github.com/boundedinfinity/go-mimetyper v1.0.17
	github.com/boundedinfinity/rfc3339date v1.0.7
	github.com/dave/jennifer v1.7.0
	github.com/gertd/go-pluralize v0.2.1
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.9.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
)

replace github.com/boundedinfinity/go-commoner => ../../../go-commoner/
