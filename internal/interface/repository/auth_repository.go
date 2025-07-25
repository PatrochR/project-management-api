package repository

import "github.com/patorochr/project-management-api/internal/entity"

type AuthRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}
