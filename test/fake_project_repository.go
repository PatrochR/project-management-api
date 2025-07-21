package test

import "github.com/patorochr/project-management-api/internal/entity"

type fakeProjectRepository struct {
}

func NewFakeProjectRepository() *fakeProjectRepository {
	return &fakeProjectRepository{}
}

func (f *fakeProjectRepository) GetByOwnerId(int) (*[]entity.Project, error) {
	return &[]entity.Project{}, nil
}
func (f *fakeProjectRepository) GetById(int) (*entity.Project, error) {
	return &entity.Project{}, nil
}
func (f *fakeProjectRepository) Create(*entity.Project) (*int, error) {
	return nil, nil
}
func (f *fakeProjectRepository) Update(*entity.Project, int) error {
	return nil
}
func (f *fakeProjectRepository) Delete(int) error {
	return nil
}
func (f *fakeProjectRepository) IsItOwner(ownerId, projectId int) error {
	return nil
}
