package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/patorochr/project-management-api/internal/interface/controller/dto"
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

// GetByProjectId godoc
// @Summary Get all members of a project
// @Description Returns a list of users assigned to a given project
// @Tags project-members
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {array} entity.ProjectMember
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects/{id}/members [get]
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

// AddMemberToProject godoc
// @Summary Add a member to a project
// @Description Adds a user to the project with a given role
// @Tags project-members
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 201 {string} string "Member added"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects/{id}/members [post]
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
	input := dto.ProjectMemberRequest{}

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

// DeleteMemberToProject godoc
// @Summary Remove a user from a project
// @Description Deletes a user from a project if requester has access
// @Tags project-members
// @Produce json
// @Param id path int true "Project ID"
// @Param userId path int true "User ID to remove"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects/{id}/members/{userId} [delete]
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
