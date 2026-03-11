package repository

import (
	"context"
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
