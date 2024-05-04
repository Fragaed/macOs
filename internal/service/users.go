package service

import (
	todo "Fragaed"
	"Fragaed/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser(userId int) (todo.User, error) {
	return s.repo.GetUser(userId)
}

func (s *UserService) UpdateUser(user todo.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(userId int) (todo.User, error) {
	return s.repo.DeleteUser(userId)
}

func (s *UserService) ListAllUsers(c todo.Conditions) ([]todo.User, error) {
	return s.repo.ListAllUsers(c)
}

func generatePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash)
}
