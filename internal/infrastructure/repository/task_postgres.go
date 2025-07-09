package db

import (
	"database/sql"

	"github.com/patorochr/project-management-api/internal/entity"
)

type PostgresTaskRepo struct {
	db *sql.DB
}

func NewPostgresTaskRepo(db *sql.DB) *PostgresTaskRepo {
	return &PostgresTaskRepo{
		db: db,
	}
}

func (r *PostgresTaskRepo) GetTaskByProjectId(projectId int) (*[]entity.Task, error) {
	query := `
		select * from task where Project_Id = $1
	`
	rows, err := r.db.Query(query, projectId)
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.ProjectId, &task.OwnerId, &task.AssigneeId, &task.Deadline, &task.CreatedAt); err != nil {
			return nil, err
		}
	}
	return &tasks, nil
}

func (r *PostgresTaskRepo) GetTaskById(taskId int) (*entity.Task, error) {
	query := `
		select * from task where Id = $1
	`
	row := r.db.QueryRow(query, taskId)
	var task entity.Task
	if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.ProjectId, &task.OwnerId, &task.AssigneeId, &task.Deadline, &task.CreatedAt); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *PostgresTaskRepo) Create(task *entity.Task) error {
	query := `
		insert into task (Title , Description , Status , Project_Id , Owner_Id , Assignee_Id , Deadline , Created_At)
		valuse ($1,$2,$3,$4,$5,$6,$7,$8)
	`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ProjectId, task.OwnerId, task.AssigneeId, task.Deadline, task.CreatedAt)
	return err
}

func (r *PostgresTaskRepo) IsItOwner(ownerId, taskId int) error {
	query := `
		select Owner_Id from tasks where Id = $1
	`
	row := r.db.QueryRow(query, taskId)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	if id != ownerId {
		return err
	}
	return nil
}

func (r *PostgresTaskRepo) Update(task *entity.Task, id int) error {
	query := `
		update tasks set Title=$1, Description=$2, Status=$3 , Assignee_Id=$4 , Deadline=$5 where Id = $6
	`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.AssigneeId, task.Deadline, id)
	return err
}

func (r *PostgresTaskRepo) Delete(id int) error {
	query := `
		delete from tasks where Id = $1
	`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresTaskRepo) Init() error {
	return r.createTable()
}

func (r *PostgresTaskRepo) createTable() error {
	query := `
		create table if not exists task(
			Id serial primary key,
			Title varchar(150),
			Description text ,
			status varchar(20),
			Project_Id integer not null references projects(Id),
			Owner_Id integer not null references users(Id),
			Assignee_Id integer not null references users(Id),
			Deadline timestamp,
			Created_At timestamp,
		) 
	`
	_, err := r.db.Exec(query)
	return err
}
