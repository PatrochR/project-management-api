package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/patorochr/project-management-api/internal/interface/helper"
	"github.com/patorochr/project-management-api/internal/usecase"
)

type TaskController struct {
	uc        *usecase.TaskUseCase
	validator *validator.Validate
}

func NewTaskController(uc *usecase.TaskUseCase, validator *validator.Validate) *TaskController {
	return &TaskController{
		uc:        uc,
		validator: validator,
	}
}

func (c *TaskController) GetBYProjectId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectIdStr, ok := params["projectId"]
	if !ok {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		http.Error(w, "invalid projectId", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, userId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tasks, err := c.uc.GetTaskByProjectId(projectId, int(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := helper.WriteJSON(w, http.StatusOK, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *TaskController) GetBYId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskIdStr, ok := params["taskId"]
	if !ok {
		http.Error(w, "taskId is required", http.StatusBadRequest)
		return
	}
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid taskId", http.StatusBadRequest)
		return
	}

	userIdFloat, ok := r.Context().Value("userID").(float64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userId := int(userIdFloat)

	task, err := c.uc.GetTaskById(taskId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := helper.WriteJSON(w, http.StatusOK, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string    `json:"title" vaildate:"required"`
		Description string    `json:"description" vaildate:"required"`
		Status      string    `json:"status" vaildate:"oneof=todo in_progress done"`
		AssigneeId  int       `json:"assignee_id"`
		Deadline    time.Time `json:"deadline" vaildate:"datetime"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input value", http.StatusBadRequest)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	projectIdStr, ok := params["projectId"]
	if !ok {
		http.Error(w, "projectId is required", http.StatusBadRequest)
		return
	}
	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		http.Error(w, "invalid projectId", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, userId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err := c.uc.Create(input.Title, input.Description, input.Status, projectId, int(userId), &input.AssigneeId, &input.Deadline); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := helper.WriteJSON(w, http.StatusNoContent, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string    `json:"title" vaildate:"required"`
		Description string    `json:"description" vaildate:"required"`
		Status      string    `json:"status" vaildate:"oneof=todo in_progress done"`
		AssigneeId  int       `json:"assignee_id"`
		Deadline    time.Time `json:"deadline" vaildate:"datetime"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input value", http.StatusBadRequest)
		return
	}
	if err := c.validator.Struct(input); err != nil {
		http.Error(w, "validation error", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	taskIdStr, ok := params["taskId"]
	if !ok {
		http.Error(w, "taskId is required", http.StatusBadRequest)
		return
	}
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid taskId", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, userId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = c.uc.Update(&input.Title, &input.Description, &input.Status, &input.AssigneeId, &input.Deadline, taskId, int(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := helper.WriteJSON(w, http.StatusNoContent, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskIdStr, ok := params["taskId"]
	if !ok {
		http.Error(w, "taskId is required", http.StatusBadRequest)
		return
	}
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid taskId", http.StatusBadRequest)
		return
	}
	userId, ok := r.Context().Value("userID").(float64)
	if !ok {
		log.Println(ok, userId)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := c.uc.Delete(taskId, int(userId)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := helper.WriteJSON(w, http.StatusNoContent, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
