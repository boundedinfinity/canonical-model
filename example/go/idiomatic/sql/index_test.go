package sql_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/sql"
	"github.com/stretchr/testify/assert"
)

func create_index_table() *sql.TableSchema {
	return sql.Table().Name("person").Build()
}

func Test_Index_Generate(t *testing.T) {

	first_name1 := sql.Column().Name("first_name").Table(create_index_table()).Build()
	first_name2 := sql.Column().Name("first_name").Unique(true).Table(create_index_table()).Build()

	tcs := []struct {
		name     string
		input    *sql.IndexSchema
		expected string
		err      error
	}{
		{
			name:     "case 1",
			input:    sql.Index().Column(first_name1).Build(),
			expected: `CREATE INDEX "idx_person_first_name" ON ("first_name");`,
		},
		{
			name:     "case 2",
			input:    sql.Index().Column(first_name2).Build(),
			expected: `CREATE UNIQUE INDEX "idx_person_first_name" ON ("first_name");`,
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
