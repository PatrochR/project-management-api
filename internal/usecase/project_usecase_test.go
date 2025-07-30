package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/patorochr/project-management-api/internal/entity"
<<<<<<< HEAD
	"github.com/patorochr/project-management-api/internal/interface/helper"
=======
>>>>>>> f0183d44b78e7d868d2c741e9877471bf97bccda
	"github.com/patorochr/project-management-api/mocks"
	"github.com/stretchr/testify/assert"
)

type getByOwnerIdInput struct {
	ownerId int
<<<<<<< HEAD
}

var getByOwnerIdCases = []struct {
	name        string
	input       getByOwnerIdInput
	mockSetup   func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy)
	expectErr   error
	expectCalls bool
}{
	{name: "success", input: getByOwnerIdInput{ownerId: 1}, expectErr: nil, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetByOwnerId", 1).Return(&[]entity.Project{
			{
				Id:          1,
				Name:        "project",
				Description: "project's description",
				Owner:       1,
				CreatedAt:   time.Now().UTC(),
			},
		}, nil)
	}},
	{name: "db error", input: getByOwnerIdInput{ownerId: 1}, expectErr: helper.ErrDb, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetByOwnerId", 1).Return(nil, helper.ErrDb)
	}},
}

func TestGetByOwnerId(t *testing.T) {
=======
}

var getByOwnerIdCases = []struct {
	name        string
	input       getByOwnerIdInput
	mockSetup   func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy)
	expectErr   error
	expectCalls bool
}{
	{name: "success", input: getByOwnerIdInput{ownerId: 1}, expectErr: nil, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetByOwnerId", 1).Return(&[]entity.Project{
			{
				Id:          1,
				Name:        "project",
				Description: "project's description",
				Owner:       1,
				CreatedAt:   time.Now().UTC(),
			},
		}, nil)
	}},
	{name: "db error", input: getByOwnerIdInput{ownerId: 1}, expectErr: errors.New("db error"), expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetByOwnerId", 1).Return(nil, errors.New("db error"))
	}},
}

func TestGetOwnerId(t *testing.T) {
>>>>>>> f0183d44b78e7d868d2c741e9877471bf97bccda
	for _, tc := range getByOwnerIdCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			projectMock := mocks.NewProjectRepository(t)
			projectMemberMock := mocks.NewProjectMemberRepostiroy(t)
			tc.mockSetup(projectMock, projectMemberMock)
			uc := NewProjectUseCase(projectMock, projectMemberMock)
			projects, err := uc.GetByOwnerId(tc.input.ownerId)
			if tc.expectErr != nil {
				assert.Error(err)
<<<<<<< HEAD
				assert.ErrorIs(err, tc.expectErr)
=======
>>>>>>> f0183d44b78e7d868d2c741e9877471bf97bccda
				assert.Nil(projects)
			} else {
				assert.NoError(err)
				assert.NotNil(projects)
			}

			if tc.expectCalls {
				projectMock.AssertExpectations(t)
			} else {
				projectMock.AssertNotCalled(t, "GetByOwnerId")
			}

<<<<<<< HEAD
		})
	}

}

type getByIdInput struct {
	ownerId int
	id      int
}

var getByIdCases = []struct {
	name        string
	input       getByIdInput
	mockSetup   func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy)
	expectErr   error
	expectCalls bool
}{
	{name: "success", input: getByIdInput{ownerId: 1, id: 1}, expectErr: nil, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetById", 1).Return(&entity.Project{
			Id:          1,
			Name:        "project",
			Description: "project's description",
			Owner:       1,
			CreatedAt:   time.Now().UTC(),
		}, nil)
	}},
	{name: "db error", input: getByIdInput{ownerId: 1, id: 1}, expectErr: helper.ErrNotFound, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetById", 1).Return(nil, helper.ErrDb)
	}},
	{name: "no access", input: getByIdInput{ownerId: 1, id: 1}, expectErr: helper.ErrNoAccess, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetById", 1).Return(&entity.Project{
			Id:          1,
			Name:        "project",
			Description: "project's description",
			Owner:       2,
			CreatedAt:   time.Now().UTC(),
		}, nil)
	}},
}

func TestGetById(t *testing.T) {
	for _, tc := range getByIdCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			projectMock := mocks.NewProjectRepository(t)
			projectMemberMock := mocks.NewProjectMemberRepostiroy(t)
			tc.mockSetup(projectMock, projectMemberMock)
			uc := NewProjectUseCase(projectMock, projectMemberMock)
			projects, err := uc.GetById(tc.input.ownerId, tc.input.id)
			if tc.expectErr != nil {
				assert.Error(err)
				assert.ErrorIs(err, tc.expectErr)
				assert.Nil(projects)
			} else {
				assert.NoError(err)
				assert.NotNil(projects)
			}

			if tc.expectCalls {
				projectMock.AssertExpectations(t)
			} else {
				projectMock.AssertNotCalled(t, "GetById")
			}

		})
	}

}

type createInput struct {
	name        string
	description string
	ownerId     int
}

var createCases = []struct {
	name        string
	input       createInput
	mockSetup   func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy)
	expectErr   error
	expectCalls bool
}{
	{name: "success", input: createInput{name: "project1", description: "description of project1", ownerId: 1}, expectErr: nil, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		var project = &entity.Project{
			Id:          0,
			Name:        "project1",
			Description: "description of project1",
			Owner:       1,
			CreatedAt:   time.Now().UTC(),
		}
		projectMock.On("Create", project).Return(&project.Id, nil)

		projectMemberMock.On("Create", &entity.ProjectMember{
			UserId:    1,
			Role:      "owner",
			CreatedAt: time.Now().UTC(),
			ProjectId: project.Id,
		}).Return(nil)
	}},
}

func TestCreate(t *testing.T) {
	for _, tc := range createCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			projectMock := mocks.NewProjectRepository(t)
			projectMemberMock := mocks.NewProjectMemberRepostiroy(t)
			tc.mockSetup(projectMock, projectMemberMock)
			uc := NewProjectUseCase(projectMock, projectMemberMock)
			err := uc.Create(tc.input.name, tc.input.description, tc.input.ownerId)
			if tc.expectErr != nil {
				assert.Error(err)
			} else {
				assert.NoError(err)

			}
			if tc.expectCalls {
				projectMock.AssertExpectations(t)
				projectMemberMock.AssertExpectations(t)
			} else {
				projectMock.AssertNotCalled(t, "Create")
				projectMemberMock.AssertNotCalled(t, "Create")
			}

		})
	}

}

type updateInput struct {
	name        string
	description string
	ownerId     int
	id          int
}

var updateCases = []struct {
	name        string
	input       updateInput
	mockSetup   func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy)
	expectErr   error
	expectCalls bool
}{
	{name: "success", input: updateInput{name: "project1", description: "description of project1", ownerId: 1, id: 1}, expectErr: nil, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetById", 1).Return(&entity.Project{
			Id:          1,
			Name:        "project",
			Description: "project's description",
			Owner:       1,
			CreatedAt:   time.Now().UTC(),
		}, nil)
		var project = &entity.Project{
			Name:        "project1",
			Description: "description of project1",
		}
		projectMock.On("Update", project, 1).Return(nil)
	}},
}

func TestUpdate(t *testing.T) {
	for _, tc := range updateCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			projectMock := mocks.NewProjectRepository(t)
			projectMemberMock := mocks.NewProjectMemberRepostiroy(t)
			tc.mockSetup(projectMock, projectMemberMock)
			uc := NewProjectUseCase(projectMock, projectMemberMock)
			err := uc.Update(tc.input.name, tc.input.description, tc.input.ownerId, tc.input.id)
			if tc.expectErr != nil {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
			if tc.expectCalls {
				projectMock.AssertExpectations(t)
			} else {
				projectMock.AssertNotCalled(t, "Update")
			}

		})
	}

}

type deleteInput struct {
	ownerId int
	id      int
}

var deleteCases = []struct {
	name        string
	input       deleteInput
	mockSetup   func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy)
	expectErr   error
	expectCalls bool
}{
	{name: "success", input: deleteInput{ownerId: 1, id: 1}, expectErr: nil, expectCalls: true, mockSetup: func(projectMock *mocks.ProjectRepository, projectMemberMock *mocks.ProjectMemberRepostiroy) {
		projectMock.On("GetById", 1).Return(&entity.Project{
			Id:          1,
			Name:        "project",
			Description: "project's description",
			Owner:       1,
			CreatedAt:   time.Now().UTC(),
		}, nil)
		projectMock.On("Delete", 1).Return(nil)
	}},
}

func TestDelete(t *testing.T) {
	for _, tc := range deleteCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			projectMock := mocks.NewProjectRepository(t)
			projectMemberMock := mocks.NewProjectMemberRepostiroy(t)
			tc.mockSetup(projectMock, projectMemberMock)
			uc := NewProjectUseCase(projectMock, projectMemberMock)
			err := uc.Delete(tc.input.ownerId, tc.input.id)
			if tc.expectErr != nil {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
			if tc.expectCalls {
				projectMock.AssertExpectations(t)
			} else {
				projectMock.AssertNotCalled(t, "Delete")
			}

=======
>>>>>>> f0183d44b78e7d868d2c741e9877471bf97bccda
		})
	}

}
