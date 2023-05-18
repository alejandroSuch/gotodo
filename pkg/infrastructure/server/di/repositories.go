package di

import (
	"gotodo/pkg/domain"
	"gotodo/pkg/domain/impl"
)

func InitializeAuthRepository() domain.AuthenticationRepository {
	sqlDB := InitializeDB()
	sqliteAuthRepository := impl.NewSqliteAuthRepository(sqlDB)
	return sqliteAuthRepository
}

func InitializeUserRepository() domain.UserRepository {
	sqlDB := InitializeDB()
	sqliteUserRepository := impl.NewSqliteUserRepository(sqlDB)
	return sqliteUserRepository
}
