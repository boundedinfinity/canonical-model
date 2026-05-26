package sql

import (
	"github.com/boundedinfinity/canonical-model/go/gen/model"
	"github.com/boundedinfinity/go-commoner/errorer"
	o "github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var (
	ErrSql      = errorer.New("sql")
	errKindFunc = errorer.Func(model.ErrKind, ErrSql)
)

type generator struct {
}

func (this generator) GenerateSql(config model.Config) (string, error) {
	database := sqlDatabase{}

	for ki, kind := range config.Kinds {
		switch k := kind.(type) {
		case *model.ObjectModel:
			table := &sqlTable{Name: k.Name}
			for pi, property := range k.Properties {
				if err := this.processSqlProperty(table, ki, kind, pi, property); err != nil {
					return "", err
				}
			}
			database.Tables = append(database.Tables, table)
		default:
			errKindFunc("config.kinds[%d]: not supported: %s", ki, kind.GetKind())
		}
	}

	var lines []string

	return stringer.Join("\n", lines...), nil
}

func (this generator) sqlTableName(kind Kind) string {

	return ""
}

func (this generator) processSqlProperty(table *sqlTable, ki int, kind model.Model, pi int, property model.Property) error {
	var names []o.Option[string]
	var types []o.Option[string]
	column := sqlColumn{}

	switch k := property.Kind.(type) {
	case *model.FloatModel:
		column.Name = k.Name
		column.Type = Kinds.Real.String()
	case *model.IntegerModel:
		column.Name = k.Name
		column.Type = Kinds.Integer.String()
	case *model.StringModel:
		column.Name = k.Name
		column.Type = Kinds.String.String()
	case *model.BooleanModel:
		column.Name = k.Name
		column.Type = Kinds.Boolean.String()
		// TODO: add CHECK constraint for boolean values
		// CHECK (column IN (0, 1))
		// CHECK (column IN (0, 1) OR column IS NULL) -- if optional
		// CHECK (column IN (0, 1) OR column IS NULL) -- if optional and not null
		// CHECK (column IN (0, 1)) -- if not optional and not null
		// CHECK (column IN (0, 1) OR column IS NULL) -- if not optional and nullable
		// CHECK (column IN (0, 1) OR column IS NULL) -- if optional and nullable
		// CHECK (column IN (0, 1) OR column IS NULL) -- if optional and nullable
	case *model.EnumerationModel:
		column.Name = k.Name
		column.Type = Kinds.String.String()
		// TODO: add CHECK constraint for enumeration values
		// CHECK (column IN ('value1', 'value2', ...))
		// CHECK (column IN ('value1', 'value2', ...) OR column IS NULL) -- if optional
		// CHECK (column IN ('value1', 'value2', ...)) -- if not optional
		// CHECK (column IN ('value1', 'value2', ...) OR column IS NULL) -- if optional and nullable
		// CHECK (column IN ('value1', 'value2', ...)) -- if not optional and not null
		// CHECK (column IN ('value1', 'value2', ...) OR column IS NULL) -- if not optional and nullable
	case *model.LanguageModel:
		if k.Database.Defined() {
			names = append(names, k.Database.Get().Name)
		}
		names = append(names, k.Name, k.Alias, o.Some(k.Type))

		if k.Database.Defined() {
			types = append(types, k.Database.Get().Type)
		}
		types = append(types, o.Some("TEXT"))
	case *model.ReferenceModel:
	default:
		errKindFunc("config.kinds[%d].properties[%d]: not supported: %s", ki, pi, property.Kind.GetKind())
	}

	column.Name = o.FirstOf(names...).Get()
	column.Type = o.FirstOf(types...).Get()
	table.Columns = append(table.Columns, &column)

	return nil
}
