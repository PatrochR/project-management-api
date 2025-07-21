package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var registerCases = []struct {
	name  string
	input struct {
		email    string
		password string
	}
	expect  any
	wantErr bool
}{

	{name: "nil password", input: struct {
		email    string
		password string
	}{email: "reza@gmail.com", password: ""}, expect: "password is not secure", wantErr: true},
	{name: "short password", input: struct {
		email    string
		password string
	}{email: "reza@gmail.com", password: "123"}, expect: "password is not secure", wantErr: true},
	{name: "happy path", input: struct {
		email    string
		password string
	}{email: "reza@gmail.com", password: "1234567"}, expect: nil, wantErr: false},
}

func TestRegister(t *testing.T) {
	for _, s := range registerCases {
		t.Run(s.name, func(t *testing.T) {
			assert := assert.New(t)

			fakeRepo := test.NewFakeAuthRepository()
			uc := NewAuthUseCase(fakeRepo)
			user, err := uc.Register(s.input.email, s.input.password)
			if s.wantErr {
				assert.EqualError(err, s.expect.(string))
			} else {
				assert.NoError(err)
				assert.Equal(user.Email, s.input.email)

				delta := time.Second
				now := time.Now().UTC()
				assert.WithinDuration(now, user.CreatedAt, delta)

				err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(s.input.password))
				assert.NoError(err)
			}
		})
	}
}

var loginCases = []struct {
	name  string
	input struct {
		email    string
		password string
	}
	expect  any
	wantErr bool
}{
	{name: "happy path", input: struct {
		email    string
		password string
	}{
		email:    "reza@gmail.com",
		password: "1234567",
	}, expect: nil, wantErr: false},
	{name: "invalid email", input: struct {
		email    string
		password string
	}{
		email:    "",
		password: "1234567",
	}, expect: helper.ErrWrongEmailOrPassowrd, wantErr: true},
	{name: "invalid password", input: struct {
		email    string
		password string
	}{
		email:    "reza@gmail.com",
		password: "123456",
	}, expect: helper.ErrWrongEmailOrPassowrd, wantErr: true},
}

func TestLogin(t *testing.T) {
	for _, c := range loginCases {
		t.Run(c.name, func(t *testing.T) {
			assert := assert.New(t)
			fakeRepo := test.NewFakeAuthRepository()
			uc := NewAuthUseCase(fakeRepo)
			token, err := uc.Login(c.input.email, c.input.password)
			if c.wantErr {
				assert.Error(err, c.expect)
			} else {
				assert.NoError(err)
				assert.NotNil(token)
				parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
					return secretKey, nil
				})
				assert.NoError(err)
				claims, ok := parsed.Claims.(jwt.MapClaims)
				assert.True(ok)
				assert.Equal("1", fmt.Sprintf("%v", claims["id"]))
			}

		})
	}

}
