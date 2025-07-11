package main

import (
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
	database "github.com/patorochr/project-management-api/internal/infrastructure/repository"
	"github.com/patorochr/project-management-api/internal/infrastructure/router"
	controller "github.com/patorochr/project-management-api/internal/interface/contorller"
	"github.com/patorochr/project-management-api/internal/usercase"
)

func main() {

	validator := validator.New()

	connStr := "user=postgres dbname=postgres port=5431 password=thisworldshallknowpain720 sslmode=disable"
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

	projectUsecase := usercase.NewProjectUseCase(projectPostgres, projectMemberPostgres)
	authUsecase := usercase.NewAuthUseCase(authPostgres)
	projectMemberUsecase := usercase.NewProjectMemberUseCase(projectMemberPostgres, projectPostgres)
	taskUsecase := usercase.NewTaskUseCase(taskPostgres, projectMemberPostgres)

	projectHandler := controller.NewProjectContoller(projectUsecase)
	authHandler := controller.NewAuthController(authUsecase)
	projectMemberHandler := controller.NewProjectMemberController(projectMemberUsecase)
	taskHandler := controller.NewTaskController(taskUsecase, validator)

	router := router.NewAPIServier(":8888", authHandler, projectHandler, projectMemberHandler, taskHandler)

	router.Run()
}
