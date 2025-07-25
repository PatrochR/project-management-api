package usecase

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type registerInput struct {
	email    string
	password string
}

var registerCases = []struct {
	name        string
	input       registerInput
	expectErr   error
	mockSetup   func(mockRepo *mocks.AuthRepository)
	expectCalls bool
}{

	{name: "nil password", input: registerInput{email: "reza@gmail.com", password: ""}, expectErr: errors.New("password is not secure"), mockSetup: func(mockRepo *mocks.AuthRepository) {}, expectCalls: false},
	{name: "short password", input: registerInput{email: "reza@gmail.com", password: "123"}, expectErr: errors.New("password is not secure"), mockSetup: func(mockRepo *mocks.AuthRepository) {}, expectCalls: false},
	{name: "success", input: registerInput{email: "reza@gmail.com", password: "1234567"}, expectErr: nil, mockSetup: func(mockRepo *mocks.AuthRepository) {
		mockRepo.On("CreateUser", mock.Anything).Return(nil)
	}, expectCalls: true},
	{name: "db error", input: registerInput{email: "reza@gmail.com", password: "1234567"}, expectErr: errors.New("db error"), mockSetup: func(mockRepo *mocks.AuthRepository) {
		mockRepo.On("CreateUser", mock.Anything).Return(errors.New("db error"))
	}, expectCalls: true},
}

func TestRegister(t *testing.T) {
	for _, tc := range registerCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			mockRepo := mocks.NewAuthRepository(t)
			tc.mockSetup(mockRepo)
			uc := NewAuthUseCase(mockRepo)
			user, err := uc.Register(tc.input.email, tc.input.password)
			if tc.expectErr != nil {
				assert.EqualError(err, tc.expectErr.Error())
				assert.Nil(user)
			} else {
				assert.NoError(err)
				assert.Equal(user.Email, tc.input.email)

				delta := time.Second
				now := time.Now().UTC()
				assert.WithinDuration(now, user.CreatedAt, delta)

				err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(tc.input.password))
				assert.NoError(err)
			}

			if tc.expectCalls {
				mockRepo.AssertExpectations(t)
			} else {
				mockRepo.AssertNotCalled(t, "CreateUser")

			}
		})
	}
}

type loginInput struct {
	email    string
	password string
}

var loginCases = []struct {
	name        string
	input       loginInput
	expectErr   error
	mockSetup   func(mockRepo *mocks.AuthRepository)
	expectCalls bool
}{
	{name: "success", input: loginInput{
		email:    "reza@gmail.com",
		password: "1234567",
	}, expectErr: nil, expectCalls: true, mockSetup: func(mockRepo *mocks.AuthRepository) {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
		mockRepo.On("GetUserByEmail", "reza@gmail.com").Return(&entity.User{
			Id:           1,
			Email:        "reza@gmail.com",
			HashPassword: string(hashed),
			IsAdmin:      false,
			CreatedAt:    time.Now().UTC(),
		}, nil)

	}},
	{name: "invalid email", input: loginInput{
		email:    "",
		password: "1234567",
	}, expectErr: helper.ErrWrongEmailOrPassowrd, mockSetup: func(mockRepo *mocks.AuthRepository) {
		mockRepo.On("GetUserByEmail", "").Return(nil, errors.New(""))
	}, expectCalls: true},
	{name: "invalid password", input: loginInput{
		email:    "reza@gmail.com",
		password: "123456",
	}, expectErr: helper.ErrWrongEmailOrPassowrd, mockSetup: func(mockRepo *mocks.AuthRepository) {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
		mockRepo.On("GetUserByEmail", "reza@gmail.com").Return(&entity.User{
			Id:           1,
			Email:        "reza@gmail.com",
			HashPassword: string(hashed),
			IsAdmin:      false,
			CreatedAt:    time.Now().UTC(),
		}, nil)

	}, expectCalls: true},
}

func TestLogin(t *testing.T) {
	for _, tc := range loginCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			mockRepo := mocks.NewAuthRepository(t)
			tc.mockSetup(mockRepo)
			uc := NewAuthUseCase(mockRepo)
			token, err := uc.Login(tc.input.email, tc.input.password)
			if tc.expectErr != nil {
				assert.Error(err, tc.expectErr)
				assert.Equal(token, "")
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

			if tc.expectCalls {
				mockRepo.AssertExpectations(t)
			} else {
				mockRepo.AssertNotCalled(t, "GetUserByEmail")
			}

		})
	}

}
