package usecase

import (
	"database/sql"
	"time"

	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/internal/interface/helper"
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
		return nil, helper.ErrNoAccess
	}
	return uc.tr.GetTaskByProjectId(projectId)
}

func (uc *TaskUseCase) GetTaskById(taskId, userId int) (*entity.Task, error) {
	task, err := uc.tr.GetTaskById(taskId)
	if err != nil {
		return nil, err
	}
	if err := uc.pmr.CanUseProject(task.ProjectId, userId); err != nil {
		return nil, helper.ErrNoAccess
	}
	return task, nil
}

func (uc *TaskUseCase) Create(
	title, desciption, status string,
	projectId, ownerId int,
	assigneeId *int,
	deadline *time.Time) error {

	task := entity.Task{
		Title:       title,
		Description: desciption,
		Status:      status,
		ProjectId:   projectId,
		OwnerId:     ownerId,
		CreatedAt:   time.Now().UTC(),
	}
	if assigneeId != nil {
		task.AssigneeId = sql.NullInt64{
			Int64: int64(*assigneeId),
			Valid: true,
		}
	} else {
		task.AssigneeId = sql.NullInt64{
			Valid: false,
		}
	}
	if deadline != nil {
		task.Deadline = sql.NullTime{
			Time:  *deadline,
			Valid: true,
		}
	} else {
		task.Deadline = sql.NullTime{
			Valid: false,
		}
	}
	return uc.tr.Create(&task)
}

func (uc *TaskUseCase) Update(title, desciption, status *string, assigneeId *int, deadline *time.Time, taskId, userId int) error {
	task, err := uc.GetTaskById(taskId, userId)
	if task.OwnerId != userId {
		return helper.ErrNoAccess
	}
	if err != nil {
		return err
	}
	if title != nil {
		task.Title = *title
	}
	if desciption != nil {
		task.Description = *desciption
	}
	if status != nil {
		task.Status = *status
	}
	if assigneeId != nil {
		task.AssigneeId = sql.NullInt64{
			Int64: int64(*assigneeId),
			Valid: true,
		}
	}
	if deadline != nil {
		task.Deadline = sql.NullTime{
			Time:  *deadline,
			Valid: true,
		}
	}
	return uc.tr.Update(task, taskId)
}

func (uc *TaskUseCase) Delete(taskId, userId int) error {
	task, err := uc.GetTaskById(taskId, userId)
	if err != nil {
		return helper.ErrNotFound
	}
	if task.OwnerId != userId {
		return helper.ErrNoAccess
	}
	return uc.tr.Delete(taskId)
}
