package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/patorochr/project-management-api/internal/interface/contorller"
	"github.com/patorochr/project-management-api/internal/interface/middleware"
)

type APIServier struct {
	Address              string
	AuthHandler          *controller.AuthController
	ProjectHandler       *controller.ProjectContoller
	ProjectMemberHandler *controller.ProjectMemberController
	TaskHandler          *controller.TaskController
}

func NewAPIServier(
	address string,
	authHandler *controller.AuthController,
	projectHandler *controller.ProjectContoller,
	projectMemberHandler *controller.ProjectMemberController,
	TaskHandler *controller.TaskController) *APIServier {
	return &APIServier{
		Address:              address,
		AuthHandler:          authHandler,
		ProjectHandler:       projectHandler,
		ProjectMemberHandler: projectMemberHandler,
		TaskHandler:          TaskHandler,
	}
}

func (s *APIServier) Run() {
	r := mux.NewRouter()

	aRouter := r.PathPrefix("/auth").Subrouter()

	aRouter.HandleFunc("/register", s.AuthHandler.RegisterHandler).Methods("POST")
	aRouter.HandleFunc("/login", s.AuthHandler.LoginHandler).Methods("POST")

	pRouter := r.PathPrefix("/project").Subrouter()
	pRouter.Use(middleware.JWTMiddleware)
	pRouter.HandleFunc("{id}", s.ProjectHandler.GetById).Methods("GET")
	pRouter.HandleFunc("", s.ProjectHandler.GetByOwnerId).Methods("GET")
	pRouter.HandleFunc("", s.ProjectHandler.Create).Methods("POST")
	pRouter.HandleFunc("{id}", s.ProjectHandler.Update).Methods("PUT")
	pRouter.HandleFunc("{id}", s.ProjectHandler.Delete).Methods("DELETE")

	pmRouter := r.PathPrefix("/project/{id}/member").Subrouter()
	pmRouter.Use(middleware.JWTMiddleware)
	pmRouter.HandleFunc("", s.ProjectMemberHandler.GetByProjectId).Methods("GET")
	pmRouter.HandleFunc("", s.ProjectMemberHandler.AddMemberToProject).Methods("POST")
	pmRouter.HandleFunc("/{userId}", s.ProjectMemberHandler.DeleteMemberToProject).Methods("DELETE")

	tRouter := r.PathPrefix("/project/{projectId}/tasks").Subrouter()
	tRouter.Use(middleware.JWTMiddleware)
	tRouter.HandleFunc("", s.TaskHandler.GetBYProjectId).Methods("GET")
	tRouter.HandleFunc("", s.TaskHandler.Create).Methods("POST")

	taskRouter := r.PathPrefix("/tasks/{taskId}").Subrouter()
	taskRouter.Use(middleware.JWTMiddleware)
	taskRouter.HandleFunc("", s.TaskHandler.GetBYId).Methods("GET")
	taskRouter.HandleFunc("", s.TaskHandler.Update).Methods("PUT")
	taskRouter.HandleFunc("", s.TaskHandler.Delete).Methods("DELETE")

	log.Println("listen and serve : ", s.Address)
	if err := http.ListenAndServe(s.Address, r); err != nil {
		log.Fatal(err.Error())
	}

}
