package test

import (
	"time"

	"github.com/patorochr/project-management-api/internal/entity"
)

type fakeAuthRepository struct {
}

func NewFakeAuthRepository() *fakeAuthRepository {
	return &fakeAuthRepository{}
}

func (f *fakeAuthRepository) GetUserByEmail(email string) (*entity.User, error) {
	if email == "" {
		return &entity.User{}, nil
	}
	return &entity.User{
		Id:           1,
		Email:        "reza@gmail.com",
		HashPassword: "$2a$10$N8sekyH6vTSs7pRzBa8e6.f1Ln03eXZmjfW1Usl0hXxOF9T5Q/K6i",
		IsAdmin:      false,
		CreatedAt:    time.Time{},
	}, nil
}

func (f *fakeAuthRepository) CreateUser(*entity.User) error {
	return nil
}
