package sql

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// https://sqlite.org/datatype3.html
type Kind string

func (this Kind) String() string {
	return string(this)
}

var Kinds = kinds{
	Null:    "NULL",
	Integer: "INTEGER",
	Text:    "TEXT",
	Real:    "REAL",
	Blob:    "BLOB",
}

type kinds struct {
	Null    Kind
	Integer Kind
	Text    Kind
	Real    Kind
	Blob    Kind
}
