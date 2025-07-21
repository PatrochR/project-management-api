package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/patorochr/project-management-api/internal/interface/controller/dto"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usecase"
)

type AuthController struct {
	usecase   *usecase.AuthUseCase
	validator *validator.Validate
}

func NewAuthController(usecase *usecase.AuthUseCase, validator *validator.Validate) *AuthController {
	return &AuthController{
		usecase:   usecase,
		validator: validator,
	}
}

// RegisterHandler godoc
// @Summary User Register
// @Description Register a user
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.RegisterRequest true "register credentials"
// @Success 201 {string} string "Created - register successful"
// @Failure 400 {string} string "Bad Request - validation error"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/register [post]
func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	input := dto.RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}

	user, err := c.usecase.Register(input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helper.WriteJSON(w, http.StatusCreated, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// LoginHandler godoc
// @Summary User login
// @Description Authenticates a user and returns a JWT in a cookie and header
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "Login credentials"
// @Success 204 {string} string "No Content - login successful"
// @Failure 400 {string} string "Bad Request - validation error"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login [post]
func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	input := dto.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}

	token, err := c.usecase.Login(input.Email, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + token,
		Expires:  time.Now().Add(24 * time.Hour).UTC(),
		HttpOnly: true,
	}

	w.Header().Add("Authorization", "Bearer "+token)
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusNoContent)
}
