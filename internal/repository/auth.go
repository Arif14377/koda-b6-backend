package repository

import (
	"context"
	"fmt"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/matthewhartstonge/argon2"
)

type AuthRepository struct {
	db *pgx.Conn
}

func NewAuthRepository(db *pgx.Conn) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) Register(userID string, user *models.UserRegister) error {
	_, err := a.db.Exec(
		context.Background(),
		"INSERT INTO users (id, full_name, email, password, role_id) VALUES ($1, $2, $3, $4, '2')", userID, user.FullName, user.Email, user.Password)
	if err != nil {
		fmt.Printf("Failed to register user: %v\n", err)
		return err
	}

	return nil
}

func (a *AuthRepository) Login(email, password string) (*models.UserLogin, error) {
	rows, err := a.db.Query(context.Background(), "SELECT email, password FROM users WHERE email = $1", email)
	if err != nil {
		r := fmt.Errorf("Failed to get rows data: %w\n", err)
		return &models.UserLogin{}, r
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.UserLogin])
	if err != nil {
		r := fmt.Errorf("User not found: %w\n", err)
		return &models.UserLogin{}, r
	}

	// fmt.Println("password user from DB: ", user.Password)
	// fmt.Println("password user from parameter: ", password)

	pwdOk, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if !pwdOk {
		r := fmt.Errorf("Password not match, %w", err)
		return &models.UserLogin{}, r
	}

	return &user, nil
}
