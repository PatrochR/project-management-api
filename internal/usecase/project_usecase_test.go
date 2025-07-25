package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/mocks"
	"github.com/stretchr/testify/assert"
)

type getByOwnerIdInput struct {
	ownerId int
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

		})
	}

}
