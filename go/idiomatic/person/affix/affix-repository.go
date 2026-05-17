package affix

import (
	"context"
	"database/sql"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type AffixRepository interface {
	Save(ctx context.Context, affix *Affix) error
	Get(ctx context.Context, id ider.Id) (Affix, error)
	Delete(ctx context.Context, id ider.Id) error
}

func Sql(db *sql.DB, categoryRepository CategoryRepository) AffixRepository {
	return &sqlAffixRepository{db: db, category: categoryRepository}
}

type sqlAffixRepository struct {
	db       *sql.DB
	category CategoryRepository
}

const (
	_SQL_INSERT_AFFIX = "INSERT INTO affix (id, name, kind, description, category_id) VALUES (?, ?, ?, ?, ?)"
	_SQL_DELETE_AFFIX = "DELETE FROM affix WHERE id = ?"
	_SQL_UPDATE_AFFIX = "UPDATE affix SET name = ?, kind = ?, description = ?, category_id = ? WHERE id = ?"
	_SQL_INSERT_ABBR  = "INSERT INTO affix_abbreviation (affix_id, index, abbreviation) VALUES (?, ?, ?)"
	_SQL_DELETE_ABBR  = "DELETE FROM affix_abbreviation WHERE affix_id = ?"
)

func (this sqlAffixRepository) Save(ctx context.Context, affix *Affix) error {
	if err := this.category.Save(ctx, &affix.Category); err != nil {
		return err
	}

	if affix.Id.IsZero() {
		affix.Id = ider.New()

		if _, err := this.db.ExecContext(ctx, _SQL_INSERT_AFFIX, affix.Id, affix.Name, affix.Kind, affix.Description, affix.Category.Id); err != nil {
			return err
		}
	} else {
		if _, err := this.db.ExecContext(ctx, _SQL_UPDATE_AFFIX, affix.Name, affix.Kind, affix.Description, affix.Category.Id, affix.Id); err != nil {
			return err
		}
	}

	if _, err := this.db.ExecContext(ctx, _SQL_DELETE_ABBR, affix.Id); err != nil {
		return err
	}

	for i, abbr := range affix.Abbreviations {
		if _, err := this.db.ExecContext(ctx, _SQL_INSERT_ABBR, affix.Id, i, abbr); err != nil {
			break
		}
	}

	return nil
}

func (this sqlAffixRepository) Get(ctx context.Context, id ider.Id) (Affix, error) {
	panic("implement me")
}

func (this sqlAffixRepository) Delete(ctx context.Context, id ider.Id) error {
	panic("implement me")
}
