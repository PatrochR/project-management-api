package usecase

import (
	"fmt"
	"testing"

	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/test"
	"github.com/stretchr/testify/assert"
)

var registerStruct = []struct {
	name        string
	email       string
	password    string
	expect      *entity.User
	errorExpect error
}{

	{name: "register nil password", email: "reza@gmail.com", password: "", expect: nil, errorExpect: fmt.Errorf("password is not secure")},
}

func TestRegister(t *testing.T) {
	for _, s := range registerStruct {
		t.Run(s.name, func(t *testing.T) {
			assert := assert.New(t)

			fakeRepo := test.NewFakeAuthRepository()
			uc := NewAuthUseCase(fakeRepo)
			_, err := uc.Register(s.email, s.password)
			assert.Error(err, s.errorExpect)
		})

	}

}
