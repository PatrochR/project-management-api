package test

import "github.com/patorochr/project-management-api/internal/entity"

type fakeAuthRepository struct {
}

func NewFakeAuthRepository() *fakeAuthRepository {
	return &fakeAuthRepository{}
}

func (f *fakeAuthRepository) GetUserByEmail(string) (*entity.User, error) {
	return &entity.User{}, nil
}

func (f *fakeAuthRepository) CreateUser(*entity.User) error {
	return nil
}
