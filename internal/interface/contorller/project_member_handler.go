package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usercase"
)

type ProjectMemberController struct {
	uc *usercase.ProjectMemberUseCase
}

func NewProjectMemberController(uc *usercase.ProjectMemberUseCase) *ProjectMemberController {
	return &ProjectMemberController{
		uc: uc,
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
		Role   string `json:"role"`
		UserId int    `json:"user_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
