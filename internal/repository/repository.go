package repository

import (
	todo "Fragaed"
	"github.com/jmoiron/sqlx"
	"log"
)

type Users interface {
	CreateUser(user todo.User) (int, error)
	GetUser(userId int) (todo.User, error)
	UpdateUser(user todo.User) error
	DeleteUser(userId int) (todo.User, error)
	ListAllUsers(c todo.Conditions) ([]todo.User, error)
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	r := &Repository{
		Users: NewUserPostgres(db),
	}
	log.Println("Repository initialized")
	return r
}
