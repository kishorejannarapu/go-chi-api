package main

import (
	"fmt"
	"log"
	"net/http"

	"go-chi-api/controller"
	"go-chi-api/database"
	"go-chi-api/repository"
	"go-chi-api/routes"
	services "go-chi-api/service"
)

func main() {

	// configs.StartDB()
	// routes.Routes()
	// Configure database
	db, err := database.ConnectToPostgres("postgres://localhost:postgres:@localhost/hrms?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Configure LDAP
	// ldapConn, err := ldap.ConnectToLDAP(ldapServerURL, ldapBindDN, ldapBindPassword)
	// if err != nil {
	// 	panic(err)
	// }

	// Define user repository
	//userRepository := repository.NewUserRepositoryImpl(db, ldapConn)

	// Define user service
	//
	userRepository := repository.NewUserRepository(db)

	userService := services.NewUserService(userRepository)

	// Define users controller
	usersController := controller.NewUsersController(userService)

	// Define JWT secret
	//jwtSecret := "your_secret_key"

	// Initialize router
	router := routes.NewRouter(map[string]interface{}{
		"users": usersController,
	})

	// Add middleware
	//router.Use(middleware.Logger)
	//router.Use(middleware.AuthMiddleware(jwtSecret))

	// Start server
	fmt.Println("Server listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
