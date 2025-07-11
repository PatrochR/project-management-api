package db

import (
	"database/sql"

	"github.com/patorochr/project-management-api/internal/entity"
)

type PostgresProjectRepo struct {
	db *sql.DB
}

func NewPostgresProjectRepo(db *sql.DB) *PostgresProjectRepo {
	return &PostgresProjectRepo{
		db: db,
	}
}
func (r *PostgresProjectRepo) Init() error {
	return r.createTable()
}

func (r *PostgresProjectRepo) GetByOwnerId(ownerId int) (*[]entity.Project, error) {
	query := `
		select * from projects where Owner = $1
	`
	rows, err := r.db.Query(query, ownerId)
	if err != nil {
		return nil, err
	}
	var projects []entity.Project
	for rows.Next() {
		var project entity.Project
		err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.Owner, &project.CreatedAt)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}
	return &projects, nil
}

func (r *PostgresProjectRepo) GetById(Id int) (*entity.Project, error) {
	query := `
		select * from projects where Id = $1
	`
	row := r.db.QueryRow(query, Id)
	var project entity.Project
	err := row.Scan(&project.Id, &project.Name, &project.Description, &project.Owner, &project.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *PostgresProjectRepo) Create(project *entity.Project) (*int, error) {
	query := `
		INSERT INTO projects (Name , Description , Owner , CreatedAt) VALUES($1,$2,$3,$4) RETURNING Id
	`

	var id int
	err := r.db.QueryRow(query, project.Name, project.Description, project.Owner, project.CreatedAt).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *PostgresProjectRepo) Update(project *entity.Project, id int) error {
	query := `
		update projects set Name=$1, Description=$2 where Id = $3
	`
	_, err := r.db.Exec(query, project.Name, project.Description, id)
	return err
}

func (r *PostgresProjectRepo) Delete(id int) error {
	query := `
		delete from projects where Id = $1
	`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresProjectRepo) IsItOwner(ownerId, projectId int) error {
	query := `
		select Owner from projects where Id = $1
	`
	row := r.db.QueryRow(query, projectId)
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

func (r *PostgresProjectRepo) createTable() error {
	query := `
		create table if not exists projects(
			Id serial primary key,
			Name varchar(200),
			Description varchar(500),
			Owner integer references users (id),
			CreatedAt timestamp
		)
	`
	_, err := r.db.Exec(query)
	return err
}
