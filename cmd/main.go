package main

import (
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
	database "github.com/patorochr/project-management-api/internal/infrastructure/repository"
	"github.com/patorochr/project-management-api/internal/infrastructure/router"
	controller "github.com/patorochr/project-management-api/internal/interface/controller"
	"github.com/patorochr/project-management-api/internal/usecase"

	_ "github.com/patorochr/project-management-api/docs"
)

// @title managment api
// @version 1.0
// @description simple api project
// @host localhost:8888
// @BasePath /
func main() {

	validator := validator.New()

	connStr := "user=postgres dbname=postgres port=5434 password=yourpassword sslmode=disable host=host.docker.internal"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	authPostgres := database.NewPostgresAuthRepo(db)
	if err := authPostgres.Init(); err != nil {
		log.Fatal(err)
	}
	projectPostgres := database.NewPostgresProjectRepo(db)
	if err := projectPostgres.Init(); err != nil {
		log.Fatal(err)
	}
	projectMemberPostgres := database.NewPostgresProjectMemberRepo(db)
	if err := projectMemberPostgres.Init(); err != nil {
		log.Fatal(err)
	}
	taskPostgres := database.NewPostgresTaskRepo(db)
	if err := taskPostgres.Init(); err != nil {
		log.Fatal(err)
	}

	projectUsecase := usecase.NewProjectUseCase(projectPostgres, projectMemberPostgres)
	authUsecase := usecase.NewAuthUseCase(authPostgres)
	projectMemberUsecase := usecase.NewProjectMemberUseCase(projectMemberPostgres, projectPostgres)
	taskUsecase := usecase.NewTaskUseCase(taskPostgres, projectMemberPostgres)

	projectHandler := controller.NewProjectContoller(projectUsecase, validator)
	authHandler := controller.NewAuthController(authUsecase, validator)
	projectMemberHandler := controller.NewProjectMemberController(projectMemberUsecase, validator)
	taskHandler := controller.NewTaskController(taskUsecase, validator)

	router := router.NewAPIServier(":8888", authHandler, projectHandler, projectMemberHandler, taskHandler)

	router.Run()
}
