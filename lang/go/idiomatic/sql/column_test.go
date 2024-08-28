package sql_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/sql"
	"github.com/stretchr/testify/assert"
)

func create_column_table() *sql.TableSchema {
	return sql.Table().Name("person").Build()
}

func Test_Column_Generate(t *testing.T) {
	tcs := []struct {
		name     string
		input    *sql.ColumnSchema
		expected string
		err      error
	}{
		{
			name:     "case 1",
			input:    sql.Column().Name("first_name").Table(create_column_table()).Build(),
			expected: ``,
			err:      sql.ErrColumnSchemaNoType,
		},
		{
			name: "case 2",
			input: sql.Column().
				Table(create_column_table()).
				Name("first_name").
				Type(sql.ColumnTypes.TEXT).
				Build(),
			expected: `"first_name" TEXT`,
		},
		{
			name: "case 3",
			input: sql.Column().
				Table(create_column_table()).
				Name("first_name").
				Type(sql.ColumnTypes.TEXT).
				Build(),
			expected: `"first_name" TEXT`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := tc.input.Generate()

			assert.ErrorIs(tt, err, tc.err)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
