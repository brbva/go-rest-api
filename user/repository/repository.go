package repository

import (
	"database/sql"
	"log"
	"gorest/internal/user/model"
)

// UserRepository defines methods for interacting with the user table
type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	CreateUser(user model.User) (int, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, email, phone_number FROM users")
	if err != nil {
		log.Println("Error querying users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
		if err != nil {
			log.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(id int) (model.User, error) {
	row := r.db.QueryRow("SELECT id, first_name, last_name, email, phone_number FROM users WHERE id=$1", id)
	var user model.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
	if err != nil {
		log.Println("Error fetching user:", err)
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO users (first_name, last_name, email, phone_number) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Email, user.PhoneNumber).Scan(&id)
	if err != nil {
		log.Println("Error creating user:", err)
		return 0, err
	}
	return id, nil
}

func (r *userRepository) UpdateUser(user model.User) error {
	_, err := r.db.Exec("UPDATE users SET first_name=$1, last_name=$2, email=$3, phone_number=$4 WHERE id=$5",
		user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.ID)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	return nil
}
