package dto

import "time"

type TaskRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Status      string    `json:"status" validate:"oneof=todo in_progress done"`
	AssigneeId  int       `json:"assignee_id"`
	Deadline    time.Time `json:"deadline" validate:"datetime"`
}
