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
}

func NewAPIServier(
	address string,
	authHandler controller.AuthController,
	projectHandler controller.ProjectContoller,
	projectMemberHandler controller.ProjectMemberController) *APIServier {
	return &APIServier{
		Address:              address,
		AuthHandler:          &authHandler,
		ProjectHandler:       &projectHandler,
		ProjectMemberHandler: &projectMemberHandler,
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

	log.Println("listen and serve : ", s.Address)
	http.ListenAndServe(s.Address, r)

}
