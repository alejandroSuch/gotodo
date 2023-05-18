package di

import (
	"database/sql"
	"gotodo/pkg/infrastructure/server/db"
)

func InitializeDB() *sql.DB {
	sqlDB := db.ProvideDB()
	return sqlDB
}
