package db

import (
	"database/sql"

	"github.com/patorochr/project-management-api/internal/entity"
)

type PostgresProjectMemberRepo struct {
	db *sql.DB
}

func NewPostgresProjectMemberRepo(db *sql.DB) *PostgresProjectMemberRepo {
	return &PostgresProjectMemberRepo{
		db: db,
	}
}

func (pm *PostgresProjectMemberRepo) Init() error {
	return pm.createProjectMemberTable()

}

func (pm *PostgresProjectMemberRepo) createProjectMemberTable() error {
	query := `
		create table if not exists project_members(
			Id serial primary key,
			Project_Id integer not null references projects(Id),
			User_Id integer not null references users(Id),
			Role text default 'member',
			CreatedAt timestamp,
			unique(Project_Id , User_Id)
		)
	`
	_, err := pm.db.Exec(query)
	return err
}
func (pm *PostgresProjectMemberRepo) GetByProjectId(projectId int) (*[]entity.ProjectMember, error) {
	query := `
		select * from project_members where Project_Id = $1
	`
	rows, err := pm.db.Query(query, projectId)
	if err != nil {
		return nil, err
	}

	var projectMembers []entity.ProjectMember
	for rows.Next() {
		var projectMember entity.ProjectMember
		if err := rows.Scan(&projectMember.Id, &projectMember.ProjectId, &projectMember.UserId, &projectMember.Role, &projectMember.CreatedAt); err != nil {
			return nil, err
		}
		projectMembers = append(projectMembers, projectMember)
	}

	return &projectMembers, nil

}
func (pm *PostgresProjectMemberRepo) Create(projectMember *entity.ProjectMember) error {
	query := `
		insert into project_members (Project_Id , User_Id , Role , CreatedAt) values ($1,$2,$3,$4)
	`
	_, err := pm.db.Exec(query, projectMember.ProjectId, projectMember.UserId, projectMember.Role, projectMember.CreatedAt)
	return err
}
func (r *PostgresProjectMemberRepo) ChangeRole(role string, id int) error {
	query := `
		update project_members set Role=$1 where Id = $2
	`
	_, err := r.db.Exec(query, role, id)
	return err
}

func (r *PostgresProjectMemberRepo) Delete(userId int) error {
	query := `
		delete from project_members where User_Id = $1
	`
	_, err := r.db.Exec(query, userId)
	return err
}
