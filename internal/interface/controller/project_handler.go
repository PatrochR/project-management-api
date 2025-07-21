package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/patorochr/project-management-api/internal/interface/controller/dto"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usecase"
)

type ProjectContoller struct {
	usecase   *usecase.ProjectUseCase
	validator *validator.Validate
}

func NewProjectContoller(usecase *usecase.ProjectUseCase, validator *validator.Validate) *ProjectContoller {
	return &ProjectContoller{
		usecase:   usecase,
		validator: validator,
	}
}

// GetByOwnerId godoc
// @Summary Get all projects for the authenticated user
// @Description Returns a list of projects owned by the current user
// @Tags projects
// @Produce json
// @Success 200 {array} entity.Project
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects [get]
func (c *ProjectContoller) GetByOwnerId(w http.ResponseWriter, r *http.Request) {
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	projects, err := c.usecase.GetByOwnerId(int(owenrId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.WriteJSON(w, http.StatusOK, projects)
}

// Create godoc
// @Summary Create a new project
// @Description Create a new project for the authenticated user
// @Tags projects
// @Accept json
// @Produce json
// @Success 201 {object} map[string]string "Project created"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects [post]
func (c *ProjectContoller) Create(w http.ResponseWriter, r *http.Request) {
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	input := dto.ProjectRequest{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	if err := c.usecase.Create(input.Name, input.Description, int(owenrId)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helper.WriteJSON(w, http.StatusCreated, map[string]string{"message": "hooora"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetById godoc
// @Summary Get a project by ID
// @Description Retrieves a project by its ID if the user is the owner
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} entity.Project
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects/{id} [get]
func (c *ProjectContoller) GetById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
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
	project, err := c.usecase.GetById(int(owenrId), id)

	if err != nil {
		if errors.Is(err, helper.ErrNoAccess) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if errors.Is(err, helper.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helper.WriteJSON(w, http.StatusOK, project); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Update godoc
// @Summary Update a project
// @Description Update a project owned by the user
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects/{id} [put]
func (c *ProjectContoller) Update(w http.ResponseWriter, r *http.Request) {
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input := dto.ProjectRequest{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	if err := c.usecase.Update(input.Name, input.Description, int(owenrId), id); err != nil {
		if errors.Is(err, helper.ErrNoAccess) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if errors.Is(err, helper.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := helper.WriteJSON(w, http.StatusNoContent, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Delete godoc
// @Summary Delete a project
// @Description Delete a project owned by the user
// @Tags projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /projects/{id} [delete]
func (c *ProjectContoller) Delete(w http.ResponseWriter, r *http.Request) {
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.usecase.Delete(int(owenrId), id); err != nil {
		if errors.Is(err, helper.ErrNoAccess) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if errors.Is(err, helper.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := helper.WriteJSON(w, http.StatusNoContent, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
