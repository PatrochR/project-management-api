package usecase

import (
	"testing"

	"github.com/patorochr/project-management-api/test"
	"github.com/stretchr/testify/assert"
)

var getByOwnerIdCases = []struct {
	name  string
	input struct {
		ownerId int
	}
	expect  any
	wantErr bool
}{
	{name: "happyPath", input: struct{ ownerId int }{ownerId: 1}, expect: nil, wantErr: false},
}

// FIXME: you should use mock mother father
func TestGetOwnerId(t *testing.T) {
	for _, c := range getByOwnerIdCases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)
			fpr := test.NewFakeProjectRepository()
			fpmr := test.NewFakeProjectMemberRepository()
			uc := NewProjectUseCase(fpr, fpmr)
			projects, err := uc.GetByOwnerId(c.input.ownerId)
			if c.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
				assert.NotNil(projects)
			}

		})
	}

}
