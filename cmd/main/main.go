package main

import (
	"GoWithMySQL/internal/db"
	"GoWithMySQL/internal/handlers"
	"GoWithMySQL/internal/repositories"
	"GoWithMySQL/internal/services"
	"net/http"
)

func main() {
	// Initialize database connection
	db, err := db.NewDB()
	if err != nil {
		// Handle error
	}
	defer db.Close()

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Define routes
	http.HandleFunc("/users", userHandler.GetAllUsers)
	http.HandleFunc("/users/{id}", userHandler.GetUserByID)

	// Start HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
