package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

func (c *ProjectContoller) Create(w http.ResponseWriter, r *http.Request) {
	owenrId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, owenrId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
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
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
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
