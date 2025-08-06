package sql_test

import (
	"testing"
)

func Test_Table_Generate(t *testing.T) {
	// person_table := sql.Table().
	// 	SetIsNotExists(true).SetName("person").
	// 	Column(&sql.ColumnSchema{Name: "id", Type: sql.ColumnTypes.TEXT, PrimaryKey: true}).
	// 	Column(&sql.ColumnSchema{Name: "first_name", Type: sql.ColumnTypes.TEXT}).
	// 	Column(&sql.ColumnSchema{Name: "last_name", Type: sql.ColumnTypes.TEXT})

	// label_table := sql.Table().SetIsNotExists(true).SetName("label").
	// 	Column(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
	// 	Column(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT))

	// _ = sql.Database().
	// 	SetForeignKeys(true).
	// 	SetFormatted(true).
	// 	Table(person_table).
	// 	Table(label_table)

	// tcs := []struct {
	// 	name     string
	// 	input    *sql.SelectSchema
	// 	expected string
	// 	err      error
	// }{
	// 	// {
	// 	// 	name:     "case 1",
	// 	// 	input:    sql.Select().All(person_table),
	// 	// 	expected: `SELECT person.id, person.first_name, person.last_name FROM person;`,
	// 	// },
	// 	// {
	// 	// 	name: "case 2",
	// 	// 	input: sql.Select().
	// 	// 		All(person_table).
	// 	// 		Where(sql.Where().
	// 	// 			Equal(person_table.GetColumn("id")).
	// 	// 			And().
	// 	// 			Equal(person_table.GetColumn("first_name")),
	// 	// 		),
	// 	// 	expected: `SELECT person.id, person.first_name, person.last_name FROM person WHERE person.id = ? AND person.first_name = ?;`,
	// 	// },
	// 	// {
	// 	// 	name: "case 3",
	// 	// 	input: sql.Select().
	// 	// 		All(person_table).
	// 	// 		Where(sql.Where().
	// 	// 			Equal(person_table.GetColumn("id")).
	// 	// 			And().Equal(person_table.GetColumn("first_name")).
	// 	// 			Or().NotEqual(person_table.GetColumn("last_name")),
	// 	// 		),
	// 	// 	expected: `SELECT person.id, person.first_name, person.last_name FROM person WHERE person.id = ? AND person.first_name = ? OR person.last_name != ?;`,
	// 	// },
	// 	// {
	// 	// 	name: "case 4",
	// 	// 	input: sql.Select().
	// 	// 		Column(person_table.GetColumn("first_name")).
	// 	// 		Where(sql.Where().
	// 	// 			Equal(person_table.GetColumn("id")).
	// 	// 			And().Equal(person_table.GetColumn("first_name")).
	// 	// 			Or().NotEqual(person_table.GetColumn("last_name")),
	// 	// 		),
	// 	// 	expected: `SELECT person.first_name FROM person WHERE person.id = ? AND person.first_name = ? OR person.last_name != ?;`,
	// 	// },
	// 	// {
	// 	// 	name:     "case 5",
	// 	// 	input:    sql.Select().All(person_table, label_table),
	// 	// 	expected: `SELECT person.id, person.first_name, person.last_name, label.id, label.name FROM person, label;`,
	// 	// },
	// 	// {
	// 	// 	name: "case 5",
	// 	// 	input: sql.Select().
	// 	// 		All(person_table, label_table).
	// 	// 		Where(sql.Where().Equal(person_table.GetColumn("first_name"))),
	// 	// 	expected: `SELECT person.id, person.first_name, person.last_name, label.id, label.name FROM person, label WHERE person.first_name = ?;`,
	// 	// },
	// 	{
	// 		name: "case 6",
	// 		input: sql.Select().
	// 			Column(person_table.GetColumn("first_name")).
	// 			InterJoin(
	// 				sql.InnerJoin(person_table.GetColumn("id")).On(label_table.GetColumn("id")),
	// 			),
	// 		expected: `SELECT person.first_name FROM person INNER JOIN label ON person.id = label.id;`,
	// 	},
	// }

	// for _, tc := range tcs {
	// 	t.Run(tc.name, func(tt *testing.T) {
	// 		actual, err := tc.input.Generate()

	// 		assert.ErrorIs(tt, err, tc.err)
	// 		assert.Equal(tt, tc.expected, actual)
	// 	})
	// }
}
