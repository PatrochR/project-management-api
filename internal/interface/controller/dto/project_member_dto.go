package dto

type ProjectMemberRequest struct {
	Role   string `json:"role" validate:"oneof=owner member"`
	UserId int    `json:"user_id" validate:"required"`
}
