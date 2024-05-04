package repository

import (
	todo "Fragaed"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return id, nil
}

func (r *UserPostgres) GetUser(userId int) (todo.User, error) {
	query := fmt.Sprintf("SELECT id, name, username, email, password_hash, deleted, time_deleted FROM %s WHERE id = $1", userTable)
	row := r.db.QueryRow(query, userId)

	var user todo.User
	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Deleted, &user.TimeDeleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user with ID %d not found", userId)
		}
		return user, fmt.Errorf("error scanning user: %w", err)
	}

	return user, nil
}

func (r *UserPostgres) UpdateUser(user todo.User) error {
	query := `
        UPDATE users
        SET name = $1,
            username = $2,
            email = $3,
            password_hash = $4
        WHERE id = $5
    `
	_, err := r.db.Exec(query, user.Name, user.Username, user.Email, user.Password, user.Id)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (r *UserPostgres) DeleteUser(userId int) (todo.User, error) {
	// Подготавливаем SQL запрос для обновления записи пользователя
	query := `
    UPDATE users
    SET deleted = $1, time_deleted = $2
    WHERE id = $3
    RETURNING id, name, username, email, password_hash, deleted, time_deleted
`
	// Текущее время для установки времени удаления
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Выполняем SQL запрос с помощью Prepare и Exec
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return todo.User{}, fmt.Errorf("failed to prepare delete query: %v", err)
	}
	defer stmt.Close()

	// Выполняем запрос с передачей значений
	var deleted bool = true
	_, err = stmt.Exec(deleted, currentTime, userId)
	if err != nil {
		return todo.User{}, fmt.Errorf("failed to execute delete query: %v", err)
	}

	// Возвращаем удаленного пользователя
	deletedUser := todo.User{Id: userId,
		Deleted:     true,
		TimeDeleted: currentTime,
	}
	return deletedUser, nil
}

func (r *UserPostgres) ListAllUsers(c todo.Conditions) ([]todo.User, error) {
	limit := c.Limit
	offset := c.Offset
	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", userTable)
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error listing all users: %w", err)
	}
	defer rows.Close()

	var users []todo.User

	for rows.Next() {
		var user todo.User
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Deleted, &user.TimeDeleted)
		if err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %w", err)
	}

	return users, nil
}
