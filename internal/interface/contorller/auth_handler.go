package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usercase"
)

type AuthController struct {
	usecase *usercase.AuthUseCase
}

func NewAuthController(usecase *usercase.AuthUseCase) *AuthController {
	return &AuthController{
		usecase: usecase,
	}
}

func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.usecase.Register(request.Email, request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helper.WriteJSON(w, http.StatusCreated, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := c.usecase.Login(request.Email, request.Password)
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
