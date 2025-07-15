package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usecase"
)

type ProjectMemberController struct {
	uc        *usecase.ProjectMemberUseCase
	validator *validator.Validate
}

func NewProjectMemberController(uc *usecase.ProjectMemberUseCase, validator *validator.Validate) *ProjectMemberController {
	return &ProjectMemberController{
		uc:        uc,
		validator: validator,
	}
}

func (c *ProjectMemberController) GetByProjectId(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	members, err := c.uc.GetByProjectId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	helper.WriteJSON(w, http.StatusOK, members)
}

func (c *ProjectMemberController) AddMemberToProject(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	projectId, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var input struct {
		Role   string `json:"role" validate:"oneof=owner member"`
		UserId int    `json:"user_id" validate:"required"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	err = c.uc.AddMemberToProject(input.Role, int(owenrId), input.UserId, projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helper.WriteJSON(w, http.StatusCreated, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *ProjectMemberController) DeleteMemberToProject(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	userIdParam := mux.Vars(r)["userId"]
	projectId, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = c.uc.DeleteMemberFromProject(int(owenrId), userId, projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = helper.WriteJSON(w, http.StatusNoContent, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
