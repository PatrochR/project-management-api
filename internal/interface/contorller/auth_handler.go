package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usercase"
)

type AuthController struct {
	usecase   *usercase.AuthUseCase
	validator *validator.Validate
}

func NewAuthController(usecase *usercase.AuthUseCase, validator *validator.Validate) *AuthController {
	return &AuthController{
		usecase:   usecase,
		validator: validator,
	}
}

func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email" vaildate:"email , required"`
		Password string `json:"password" vaildate:"min=6 , required"`
	}
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

func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email" vaildate:"email , required"`
		Password string `json:"password" vaildate:"min=6 , required"`
	}
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
