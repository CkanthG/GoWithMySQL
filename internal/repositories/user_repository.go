package repositories

import (
	"GoWithMySQL/internal/models"
	"context"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	// Database query to get all users
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, password FROM users")
	log.Default().Println("rows:", rows)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Default().Fatal(err)
		}
	}(rows)

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	// Database query to get user by ID
	var user models.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Implement other CRUD operations here
