package repository

import "github.com/patorochr/project-management-api/internal/entity"

type TaskRepository interface {
	GetTaskByProjectId(int) (*[]entity.Task, error)
	GetTaskById(int) (*entity.Task, error)
	Create(*entity.Task) error
	IsItOwner(ownerId, taskId int) error
	Update(*entity.Task, int) error
	Delete(int) error
}
