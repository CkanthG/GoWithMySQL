package handlers

import (
	"GoWithMySQL/internal/services"
	_ "context"
	"encoding/json"
	_ "encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("request received")
	ctx := r.Context()
	// Call UserService to get all users
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	// Convert users slice to JSON and write response
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to marshal users to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("request received")
	// Parse ID from request URL or request body, then call UserService to get user by ID
	// Extract the path parameter "id"
	vars := mux.Vars(r)
	log.Default().Println("vars:", vars)
	id := vars["id"]
	log.Default().Println("id:", id)
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	//Convert ID to integer (assuming it's an integer)
	//If it's not an integer, handle the error
	//Then call GetUserByID method
	//Example:
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Call UserService to get user by ID
	user, err := h.userService.GetUserByID(ctx, userID)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// Convert user object to JSON and write response
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// Implement other HTTP handlers here
