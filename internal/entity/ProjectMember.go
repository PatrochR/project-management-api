package entity

import (
	"time"
)

type ProjectMember struct {
	Id        int
	ProjectId int
	UserId    int
	Role      string
	CreatedAt time.Time
}
