package repository

import "github.com/patorochr/project-management-api/internal/entity"

type ProjectRepository interface {
	GetByOwnerId(ownerId int) (*[]entity.Project, error)
	GetById(id int) (*entity.Project, error)
	Create(project *entity.Project) (*int, error)
	Update(project *entity.Project, id int) error
	Delete(id int) error
	IsItOwner(ownerId, projectId int) error
}
