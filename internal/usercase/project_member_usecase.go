package usercase

import (
	"time"

	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/interface/repository"
)

type ProjectMemberUseCase struct {
	repo        repository.ProjectMemberRepostiroy
	projectRepo repository.ProjectRepository
}

func NewProjectMemberUseCase(repo repository.ProjectMemberRepostiroy, projectRepo repository.ProjectRepository) *ProjectMemberUseCase {
	return &ProjectMemberUseCase{
		projectRepo: projectRepo,
		repo:        repo,
	}
}

func (pmu *ProjectMemberUseCase) GetByProjectId(projectId int) (*[]entity.ProjectMember, error) {
	projects, err := pmu.repo.GetByProjectId(projectId)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (pmu *ProjectMemberUseCase) AddMemberToProject(role string, ownerId, userId, projectId int) error {
	err := pmu.projectRepo.IsItOwner(ownerId, projectId)
	if err != nil {
		return helper.ErrNoAccess
	}

	member := entity.ProjectMember{
		ProjectId: projectId,
		UserId:    userId,
		Role:      role,
		CreatedAt: time.Now().UTC(),
	}
	err = pmu.repo.Create(&member)
	return err
}

//FIXME: IS IT OWNER ??????????

func (pmu *ProjectMemberUseCase) DeleteMemberFromProject(ownerId, userId, projectId int) error {
	err := pmu.projectRepo.IsItOwner(ownerId, projectId)
	if err != nil {
		return helper.ErrNoAccess
	}
	err = pmu.repo.Delete(userId)
	return err
}
