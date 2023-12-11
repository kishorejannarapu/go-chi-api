package controller

import (
	"encoding/json"
	"net/http"

	models "go-chi-api/model"
	services "go-chi-api/service"
)

type UsersController struct {
	userService services.UserService
}

func NewUsersController(userService services.UserService) *UsersController {
	return &UsersController{userService: userService}
}

func (c *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.userService.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response
}

// Implement other controller methods for user actions
