package usercase

import (
	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/internal/interface/repository"
)

type TaskUseCase struct {
	tr  repository.TaskRepository
	pmr repository.ProjectMemberRepostiroy
}

func NewTaskUseCase(tr repository.TaskRepository, pmr repository.ProjectMemberRepostiroy) *TaskUseCase {
	return &TaskUseCase{
		tr:  tr,
		pmr: pmr,
	}
}

func (uc *TaskUseCase) GetTaskByProjectId(projectId, userId int) (*[]entity.Task, error) {
	if err := uc.pmr.CanUseProject(projectId, userId); err != nil {
		return nil, err
	}
	return uc.tr.GetTaskByProjectId(projectId)
}
