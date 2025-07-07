package usercase

import (
	"log"
	"time"

	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/interface/repository"
)

type ProjectUseCase struct {
	repo repository.ProjectRepository
}

func NewProjectUseCase(repo repository.ProjectRepository) *ProjectUseCase {
	return &ProjectUseCase{
		repo: repo,
	}
}

func (uc *ProjectUseCase) GetByOwnerId(ownerId int) (*[]entity.Project, error) {
	return uc.repo.GetByOwnerId(ownerId)
}

func (uc *ProjectUseCase) GetById(ownerId, id int) (*entity.Project, error) {
	project, err := uc.repo.GetById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, helper.ErrNotFound
	}
	if project.Owner != ownerId {
		return nil, helper.ErrNoAccess
	}
	return project, err
}

func (uc *ProjectUseCase) Create(name, description string, ownerId int) error {
	project := entity.Project{
		Name:        name,
		Description: description,
		Owner:       ownerId,
		CreatedAt:   time.Now().UTC(),
	}
	return uc.repo.Create(&project)
}

func (uc *ProjectUseCase) Update(name, description string, ownerId, id int) error {
	oldProject, err := uc.repo.GetById(id)
	if err != nil {
		return helper.ErrNotFound
	}
	if oldProject.Owner != ownerId {
		return helper.ErrNoAccess
	}

	project := entity.Project{
		Name:        name,
		Description: description,
	}
	if err := uc.repo.Update(&project, id); err != nil {
		return helper.ErrDb
	}
	return nil
}

func (uc *ProjectUseCase) Delete(ownerId, id int) error {
	oldProject, err := uc.repo.GetById(id)
	if err != nil {
		log.Println(err.Error())
		return helper.ErrNotFound
	}
	if oldProject.Owner != ownerId {
		return helper.ErrNoAccess
	}
	if err := uc.repo.Delete(id); err != nil {
		return helper.ErrDb
	}
	return nil
}
