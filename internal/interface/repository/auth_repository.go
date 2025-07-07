package repository

import "github.com/patorochr/project-management-api/internal/entity"

type AuthRepository interface {
	GetUserByEmail(string) (*entity.User, error)
	CreateUser(*entity.User) error
}
