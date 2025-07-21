package repository

import "github.com/patorochr/project-management-api/internal/entity"

type ProjectMemberRepostiroy interface {
	GetByProjectId(int) (*[]entity.ProjectMember, error)
	Create(*entity.ProjectMember) error
	ChangeRole(string, int) error
	Delete(userId, projectId int) error
	CanUseProject(projectId, userId int) error
}
