package impl

import (
	"database/sql"
	"gotodo/pkg/domain"
)

type SqliteAuthRepository struct {
	db *sql.DB
}

func NewSqliteAuthRepository(db *sql.DB) *SqliteAuthRepository {
	return &SqliteAuthRepository{
		db: db,
	}
}

func (ar *SqliteAuthRepository) Save(auth *domain.Authentication) error {
	tx, err := ar.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO authentications(id, username, password_hash) VALUES (?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt.Exec(auth.ID, auth.Username, auth.PasswordHash)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (ar *SqliteAuthRepository) GetByUsername(username string) (*domain.Authentication, error) {
	var auth domain.Authentication
	err := ar.db.QueryRow("SELECT id, username, password_hash FROM authentications WHERE username = ?", username).
		Scan(&auth.ID, &auth.Username, &auth.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAuthNotFound
		}
		return nil, err
	}

	return &auth, nil
}

func (ar *SqliteAuthRepository) GetByUsernameAndPassword(username string, password string, passwordMatches domain.PasswordMatches) (*domain.Authentication, error) {
	var auth domain.Authentication

	err := ar.db.QueryRow("SELECT id, username, password_hash FROM authentications WHERE username = ?", username).
		Scan(&auth.ID, &auth.Username, &auth.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrAuthNotFound
		}
		return nil, err
	}

	if !passwordMatches(password, auth.PasswordHash) {
		return nil, domain.ErrUnauthorized
	}

	return &auth, nil
}
