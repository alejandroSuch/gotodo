package impl

import (
	"database/sql"
	"gotodo/pkg/domain"
)

type SqliteUserRepository struct {
	db *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) *SqliteUserRepository {
	return &SqliteUserRepository{
		db: db,
	}
}

func (ur *SqliteUserRepository) Create(user *domain.User) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO users(id, name, auth_id) VALUES (?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt.Exec(user.ID, user.Name, user.AuthId)
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

func (ur *SqliteUserRepository) Save(user *domain.User, nextIdentity domain.NextIdentity) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, todo := range user.Todos {
		if todo.ID == "" {
			err = ur.createTodoInTransaction(tx, &todo, user.ID, nextIdentity)
		} else {
			err = ur.updateTodoInTransaction(tx, &todo)
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (ur *SqliteUserRepository) createTodoInTransaction(tx *sql.Tx, todo *domain.Todo, userID string, nextIdentity domain.NextIdentity) error {
	_, err := tx.Exec("INSERT INTO todos (id, user_id, description, status) VALUES (?, ?, ?, ?)", nextIdentity(), userID, todo.Description, domain.TodoStatusPending)
	if err != nil {
		return err
	}

	return nil
}

func (ur *SqliteUserRepository) updateTodoInTransaction(tx *sql.Tx, todo *domain.Todo) error {
	_, err := tx.Exec("UPDATE todos SET description = ?, status = ? WHERE id = ?", todo.Description, todo.Status, todo.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *SqliteUserRepository) FindById(id string) (*domain.User, error) {
	var user domain.User

	err := ur.db.QueryRow("SELECT id, name, auth_id FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.AuthId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return ur.appendTodosTo(user)
}

func (ur *SqliteUserRepository) appendTodosTo(user domain.User) (*domain.User, error) {
	rows, err := ur.db.Query("SELECT id ,description, status FROM todos WHERE user_id = ?", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Status); err != nil {
			return nil, err
		}
		user.Todos = append(user.Todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *SqliteUserRepository) FindByAuthenticationID(authID string) (*domain.User, error) {
	var user domain.User

	err := ur.db.QueryRow("SELECT id, name, auth_id FROM users WHERE auth_id = ?", authID).
		Scan(&user.ID, &user.Name, &user.AuthId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return ur.appendTodosTo(user)
}
