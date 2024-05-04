package service

import (
	todo "Fragaed"
	"Fragaed/internal/repository"
	"log"
)

type Users interface {
	CreateUser(user todo.User) (int, error)
	GetUser(userId int) (todo.User, error)
	UpdateUser(user todo.User) error
	DeleteUser(userId int) (todo.User, error)
	ListAllUsers(c todo.Conditions) ([]todo.User, error)
}

type Service struct {
	Users
}

func NewService(repos *repository.Repository) *Service {
	service := &Service{
		Users: NewUserService(repos.Users),
	}
	log.Println("Service initialized")
	return service
}
