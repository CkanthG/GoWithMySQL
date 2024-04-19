package services

import (
	"GoWithMySQL/internal/models"
	"GoWithMySQL/internal/repositories"
	"context"
	"log"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	// Call UserRepository to fetch all users
	users, err := s.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	// Call UserRepository to fetch user by ID
	log.Default().Println("id:", id)
	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Implement other business logic here
