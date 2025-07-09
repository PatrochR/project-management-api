package repository

import "github.com/patorochr/project-management-api/internal/entity"

type ProjectMemberRepostiroy interface {
	GetByProjectId(int) (*[]entity.ProjectMember, error)
	Create(*entity.ProjectMember) error
	ChangeRole(string, int) error
	Delete(userId, projectId int) error
	//FIXME: this only check if user is the member of the project but if user=owener this well not work
	CanUseProject(projectId, userId int) error
}
