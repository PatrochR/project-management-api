package test

import "github.com/patorochr/project-management-api/internal/entity"

type fakeProjectMemberRepository struct {
}

func NewFakeProjectMemberRepository() *fakeProjectMemberRepository {
	return &fakeProjectMemberRepository{}
}

func (f *fakeProjectMemberRepository) GetByProjectId(int) (*[]entity.ProjectMember, error) {
	return &[]entity.ProjectMember{}, nil
}
func (f *fakeProjectMemberRepository) Create(*entity.ProjectMember) error {
	return nil
}
func (f *fakeProjectMemberRepository) ChangeRole(string, int) error {
	return nil
}
func (f *fakeProjectMemberRepository) Delete(userId, projectId int) error {
	return nil
}
func (f *fakeProjectMemberRepository) CanUseProject(projectId, userId int) error {
	return nil
}
