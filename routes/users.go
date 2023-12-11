package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type usersController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	// GetUserByID(w http.ResponseWriter, r *http.Request)
	// GetAllUsers(w http.ResponseWriter, r *http.Request)
	// UpdateUser(w http.ResponseWriter, r *http.Request)
	// DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewUsersRouter(usersController usersController) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", usersController.CreateUser)
	// router.Get("/{id}", usersController.GetUserByID)
	// router.Get("/", usersController.GetAllUsers)
	// router.Put("/{id}", usersController.UpdateUser)
	// router.Delete("/{id}", usersController.DeleteUser)

	return router
}
