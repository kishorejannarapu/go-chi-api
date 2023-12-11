package routes

import (
	"github.com/go-chi/chi"
)

func NewRouter(controllers map[string]interface{}) *chi.Mux {
	router := chi.NewMux()

	// Mount sub-routers
	for name, controller := range controllers {
		switch name {
		case "users":
			router.Mount("/users", NewUsersRouter(controller.(usersController)))
			// Add sub-routers for other resources
		}
	}

	return router
}
