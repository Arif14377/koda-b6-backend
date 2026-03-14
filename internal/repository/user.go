package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAllUser() *[]models.UserListRead {
	rows, err := u.db.Query(context.Background(), "SELECT id, full_name, email, phone, role_id as role FROM users")
	if err != nil {
		log.Fatalf("Gagal get Query Rows: %v", err)
		return &[]models.UserListRead{}
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.UserListRead])
	if err != nil {
		log.Fatalf("Gagal Collect Rows: %v", err)
		return &[]models.UserListRead{}
	}

	return &users

}

func (u *UserRepository) GetUserByID() {

}

func (u *UserRepository) GetUserByEmail(email string) bool {
	rows, err := u.db.Query(context.Background(), "SELECT email FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Printf("Failed to get rows data: %v\n", err)
		return false
	}
	defer rows.Close()

	_, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.UserEmail])
	if err != nil {
		fmt.Printf("User not found: %v\n", err)
		return false
	}

	return true
}
