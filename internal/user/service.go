package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

type Service interface {
	GetUsers(_ context.Context) ([]User, error)
	GetUserByID(_ context.Context, id int) (*User, error)
	GetUserByEmail(_ context.Context, email string) (*User, error)
	CreateUser(_ context.Context, user User) (User, error)
}

type service struct {
	repo UserRepository
}

func (s *service) GetUsers(_ context.Context) ([]User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, err
	}

	return users, nil
}

func (s *service) GetUserByID(_ context.Context, id int) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *service) GetUserByEmail(_ context.Context, email string) (*User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *service) CreateUser(_ context.Context, user User) (User, error) {
	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	if existingUser != nil {
		return User{}, errors.New("user with the provided email already exists")
	}

	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func NewService(repo UserRepository) Service {
	return &service{repo: repo}
}
