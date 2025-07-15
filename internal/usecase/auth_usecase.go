package usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/patorochr/project-management-api/internal/entity"
	"github.com/patorochr/project-management-api/internal/interface/repository"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("this world shall know pain 720")

type AuthUseCase struct {
	repo repository.AuthRepository
}

func NewAuthUseCase(r repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		repo: r,
	}
}

func (uc *AuthUseCase) Register(email, password string) (*entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Email:        email,
		HashPassword: string(hash),
		IsAdmin:      false,
		CreatedAt:    time.Now().UTC(),
	}
	return user, uc.repo.CreateUser(user)
}

func (uc *AuthUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  user.Id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
