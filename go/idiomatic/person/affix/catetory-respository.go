package affix

import (
	"context"
	"database/sql"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *Category) error
	Get(ctx context.Context, id ider.Id) (Category, error)
	Delete(ctx context.Context, id ider.Id) error
}

func CategorySql(db *sql.DB) CategoryRepository {
	return &sqlCategoryRepository{db: db}
}

type sqlCategoryRepository struct {
	db *sql.DB
}

const (
	_SQL_INSERT_CATEGORY = "INSERT INTO category (id, name, description) VALUES (?, ?, ?)"
	_SQL_DELETE_CATEGORY = "DELETE FROM category WHERE id = ?"
	_SQL_UPDATE_CATEGORY = "UPDATE category SET name = ?, description = ? WHERE id = ?"
)

func (this sqlCategoryRepository) Save(ctx context.Context, category *Category) error {
	var err error

	if category.Id.IsZero() {
		category.Id = ider.New()

		if _, err = this.db.ExecContext(ctx, _SQL_INSERT_CATEGORY, category.Id, category.Name, category.Description); err != nil {
			return err
		}
	} else {
		if _, err = this.db.ExecContext(ctx, _SQL_UPDATE_CATEGORY, category.Name, category.Description, category.Id); err != nil {
			return err
		}
	}

	return err
}

func (this sqlCategoryRepository) Get(ctx context.Context, id ider.Id) (Category, error) {
	panic("implement me")
}

func (this sqlCategoryRepository) Delete(ctx context.Context, id ider.Id) error {
	panic("implement me")
}
