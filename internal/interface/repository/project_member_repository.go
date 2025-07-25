package repository

import "github.com/patorochr/project-management-api/internal/entity"

type ProjectMemberRepostiroy interface {
	GetByProjectId(projectId int) (*[]entity.ProjectMember, error)
	Create(projectMember *entity.ProjectMember) error
	ChangeRole(role string, id int) error
	Delete(userId, projectId int) error
	CanUseProject(projectId, userId int) error
}
