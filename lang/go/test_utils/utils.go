package test_utils

import (
	"database/sql"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	_ "github.com/tursodatabase/go-libsql"
)

var (
	_TEST_DIR = "./test-db"
)

func DbName(name string) string {
	dbname := name
	dbname = stringer.Replace(dbname, "_", " ", "-")
	dbname = stringer.Lowercase(dbname)
	return dbname
}

func insertSql(name string, fields map[string]any) (string, []any) {
	tfields := []string{}
	tvalues := []any{}

	for fname, fvalue := range fields {
		tfields = append(tfields, DbName(fname))
		tvalues = append(tvalues, fvalue)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		DbName(name),
		stringer.Join(", ", tfields...),
		stringer.Join(", ", stringer.Duplicate(len(tvalues), "?")...),
	)

	return query, tvalues
}

func TableInsert(db *sql.DB, name string, fields map[string]any) error {
	query, values := insertSql(name, fields)
	_, err := db.Exec(query, values...)
	return err
}

func createSql(name string, fields map[string]string) string {
	tfields := []string{}

	for fname, ftype := range fields {
		tfields = append(tfields, DbName(fname)+" "+ftype)
	}

	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s);",
		DbName(name),
		stringer.Join(", ", tfields...),
	)
	return query
}

func TableCreate(db *sql.DB, name string, fields map[string]string) error {
	query := createSql(name, fields)
	_, err := db.Exec(query)
	return err
}

func GetDb(name string) (*sql.DB, func()) {
	tname := DbName(name)
	testDir := pather.Dirs.Abs(_TEST_DIR)
	pather.Dirs.Ensure(testDir)
	path := pather.Paths.Join(testDir, fmt.Sprintf("%s.db", tname))

	db, err := sql.Open("libsql", fmt.Sprintf("file:%s", path))

	if err != nil {
		panic(err)
	}

	closer := func() {
		if err := db.Close(); err != nil {
			panic(err)
		}

		if _, err := pather.Files.RemoveErr(path); err != nil {
			panic(err)
		}
	}

	return db, closer
}
