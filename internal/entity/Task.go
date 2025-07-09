package entity

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          int
	Title       string
	Description string
	Status      string
	ProjectId   int
	OwnerId     int
	AssigneeId  sql.NullInt64
	Deadline    sql.NullTime
	CreatedAt   time.Time
}
