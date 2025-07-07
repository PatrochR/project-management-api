package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/patorochr/project-management-api/internal/entity"
)

type PostgresAuthRepo struct {
	db *sql.DB
}

func NewPostgresAuthRepo(db *sql.DB) *PostgresAuthRepo {
	return &PostgresAuthRepo{
		db: db,
	}
}

func (s *PostgresAuthRepo) Init() error {
	return s.createUsersTable()
}
func (s *PostgresAuthRepo) createUsersTable() error {
	_, err := s.db.Exec(`create table if not exists users(
		Id serial primary key,
		Email varchar(50),
		HashPassword varchar(200),
		IsAdmin boolean,
		CreatedAt timestamp
	)`)

	return err
}

func (s *PostgresAuthRepo) GetUserByEmail(email string) (*entity.User, error) {
	query := `select * from users where Email = $1`
	rows := s.db.QueryRow(query, email)
	var user entity.User
	if err := rows.Scan(&user.Id, &user.Email, &user.HashPassword, &user.IsAdmin, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresAuthRepo) CreateUser(user *entity.User) error {
	query := `insert into users (Email, HashPassword , IsAdmin , CreatedAt)
		values ($1,$2,$3,$4)`
	_, err := s.db.Exec(query, user.Email, user.HashPassword, user.IsAdmin, user.CreatedAt)
	return err
}
