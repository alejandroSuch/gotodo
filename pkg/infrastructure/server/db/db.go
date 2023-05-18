package db

import (
	"database/sql"
	"sync"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func ProvideDB() *sql.DB {
	dbOnce.Do(func() {
		_db, err := sql.Open("sqlite3", "db/todos.db")
		if err != nil {
			panic(err)
		}

		db = _db
	})

	return db
}
