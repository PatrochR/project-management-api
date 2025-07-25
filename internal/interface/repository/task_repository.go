package repository

import "github.com/patorochr/project-management-api/internal/entity"

type TaskRepository interface {
	GetTaskByProjectId(projectId int) (*[]entity.Task, error)
	GetTaskById(id int) (*entity.Task, error)
	Create(task *entity.Task) error
	IsItOwner(ownerId, taskId int) error
	Update(task *entity.Task, id int) error
	Delete(id int) error
}
