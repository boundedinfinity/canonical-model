package name

import (
	"context"
	"database/sql"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person/affix"
)

type NameRespository interface {
	Save(ctx context.Context, name *Name) error
	Get(ctx context.Context, id ider.Id) (Name, error)
	Delete(ctx context.Context, id ider.Id) error
}

func Sql(db *sql.DB, affixRepository affix.AffixRepository) NameRespository {
	return &sqlNameRepository{db: db, affixRepository: affixRepository}
}

type sqlNameRepository struct {
	db              *sql.DB
	affixRepository affix.AffixRepository
}

const (
	_SQL_INSERT_NAME   = "INSERT INTO names (id, kind) VALUES (?, ?)"
	_SQL_DELETE_NAME   = "DELETE FROM names WHERE id = ?"
	_SQL_UPDATE_NAME   = "UPDATE names SET kind = ? WHERE id = ?"
	_SQL_INSERT_FIRST  = "INSERT INTO name_given (name_id, index, name) VALUES (?, ?, ?)"
	_SQL_DELETE_FIRST  = "DELETE FROM name_given WHERE name_id = ?"
	_SQL_INSERT_MIDDLE = "INSERT INTO name_middle (name_id, index, name) VALUES (?, ?, ?)"
	_SQL_DELETE_MIDDLE = "DELETE FROM name_middle WHERE name_id = ?"
	_SQL_INSERT_LAST   = "INSERT INTO name_family (name_id, index, name) VALUES (?, ?, ?)"
	_SQL_DELETE_LAST   = "DELETE FROM name_family WHERE name_id = ?"
	_SQL_INSERT_PREFIX = "INSERT INTO name_prefix (name_id, prefix_id, index ) VALUES (?, ?, ?)"
	_SQL_DELETE_PREFIX = "DELETE FROM name_prefix WHERE name_id = ?"
	_SQL_INSERT_SUFFIX = "INSERT INTO name_suffix (name_id, suffix_id, index ) VALUES (?, ?, ?)"
	_SQL_DELETE_SUFFIX = "DELETE FROM name_suffix WHERE name_id = ?"
	_SQL_SELECT_NAME   = "SELECT id, kind FROM names WHERE id = ?"
	_SQL_SELECT_FIRST  = "SELECT name FROM name_given WHERE name_id = ? ORDER BY index"
	_SQL_SELECT_MIDDLE = "SELECT name FROM name_middle WHERE name_id = ? ORDER BY index"
	_SQL_SELECT_LAST   = "SELECT name FROM name_family WHERE name_id = ? ORDER BY index"
	_SQL_SELECT_PREFIX = "SELECT prefix_id FROM name_prefix WHERE name_id = ? ORDER BY index"
	_SQL_SELECT_SUFFIX = "SELECT suffix_id FROM name_suffix WHERE name_id = ? ORDER BY index"
)

func (this sqlNameRepository) Save(ctx context.Context, name *Name) error {
	var err error

	if name.Id.IsZero() {
		name.Id = ider.New()

		if _, err = this.db.ExecContext(ctx, _SQL_INSERT_NAME, name.Id, name.Kind); err != nil {
			return err
		}
	} else {
		if _, err = this.db.ExecContext(ctx, _SQL_UPDATE_NAME, name.Kind, name.Id); err != nil {
			return err
		}
	}

	doNames := func(ns []string, deleteSql, insertSql string) error {
		if _, err = this.db.ExecContext(ctx, deleteSql, name.Id); err != nil {
			return err
		}

		for i, n := range ns {
			if _, err := this.db.ExecContext(ctx, insertSql, name.Id, i, n); err != nil {
				return err
			}
		}
		return nil
	}

	if err = doNames(name.Given, _SQL_DELETE_FIRST, _SQL_INSERT_FIRST); err != nil {
		return err
	}

	if err = doNames(name.Middle, _SQL_DELETE_MIDDLE, _SQL_INSERT_MIDDLE); err != nil {
		return err
	}

	if err = doNames(name.Family, _SQL_DELETE_LAST, _SQL_INSERT_LAST); err != nil {
		return err
	}

	doAffixes := func(affixes []affix.Affix, deleteSql, insertSql string) error {
		if _, err = this.db.ExecContext(ctx, deleteSql, name.Id); err != nil {
			return err
		}

		for i, affix := range affixes {
			if _, err := this.db.ExecContext(ctx, insertSql, name.Id, affix.Id, i); err != nil {
				return err
			}
		}
		return nil
	}

	if err = doAffixes(name.Prefix, _SQL_DELETE_PREFIX, _SQL_INSERT_PREFIX); err != nil {
		return err
	}

	if err = doAffixes(name.Suffix, _SQL_DELETE_SUFFIX, _SQL_INSERT_SUFFIX); err != nil {
		return err
	}

	return nil
}

func (this sqlNameRepository) Get(ctx context.Context, id ider.Id) (Name, error) {
	name := Name{Id: id}

	row := this.db.QueryRowContext(ctx, _SQL_SELECT_NAME, id)
	if err := row.Scan(&name.Id, &name.Kind); err != nil {
		return name, err
	}

	doNames := func(ns *[]string, sql string) error {
		rows, err := this.db.QueryContext(ctx, sql, id)

		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			var n string

			if err := rows.Scan(&n); err != nil {
				return err
			}

			*ns = append(*ns, n)
		}

		return nil
	}

	if err := doNames(&name.Given, _SQL_SELECT_FIRST); err != nil {
		return name, err
	}

	if err := doNames(&name.Middle, _SQL_SELECT_MIDDLE); err != nil {
		return name, err
	}

	if err := doNames(&name.Family, _SQL_SELECT_LAST); err != nil {
		return name, err
	}

	doAffixes := func(affixes *[]affix.Affix, sql string) error {
		rows, err := this.db.QueryContext(ctx, sql, id)

		if err != nil {
			return err
		}

		defer rows.Close()

		for rows.Next() {
			var affixId ider.Id

			if err := rows.Scan(&affixId); err != nil {
				return err
			}

			affix, err := this.affixRepository.Get(ctx, affixId)

			if err != nil {
				return err
			}

			*affixes = append(*affixes, affix)
		}

		return nil
	}

	if err := doAffixes(&name.Prefix, _SQL_SELECT_PREFIX); err != nil {
		return name, err
	}

	if err := doAffixes(&name.Suffix, _SQL_SELECT_SUFFIX); err != nil {
		return name, err
	}

	return name, nil
}

func (this sqlNameRepository) Delete(ctx context.Context, id ider.Id) error {
	if _, err := this.db.ExecContext(ctx, _SQL_DELETE_NAME, id); err != nil {
		return err
	}

	return nil
}
