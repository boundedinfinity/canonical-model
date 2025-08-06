package model

import "github.com/boundedinfinity/schema/idiomatic/avatar/audit"

type Item[T any] struct {
	Value T
	Audit audit.Audit
}
