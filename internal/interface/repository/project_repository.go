package repository

import "github.com/patorochr/project-management-api/internal/entity"

type ProjectRepository interface {
	GetByOwnerId(int) (*[]entity.Project, error)
	GetById(int) (*entity.Project, error)
	Create(*entity.Project) (*int, error)
	Update(*entity.Project, int) error
	Delete(int) error
	IsItOwner(ownerId, projectId int) error
}
