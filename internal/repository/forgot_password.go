package repository

import (
	"context"
	"fmt"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ForgotPassword struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository(db *pgx.Conn) *ForgotPassword {
	return &ForgotPassword{
		db: db,
	}
}

// 1. Kirim email
// 2. Get user by email (check if exists)
// 3. Mendapatkan kode (generated code disimpan di table forgot_password bersama email)
// 4. Input kode
// 5. Kirim email dan kode ke server (request forgot password)
// 6. CEK apakah sesuai? (di tabel forgot_password)
// 7. Jika sesuai maka ubah password di table users.
// 8. Hapus code di table forgot_password

// TODO: add method get user by email
func (fp *ForgotPassword) GetUserByEmail(email string) (*models.ForgotPassword, error) {
	rows, err := fp.db.Query(context.Background(), "SELECT email FROM users WHERE email = $1", email)
	if err != nil {
		r := fmt.Errorf("Failed to get rows data: %w\n", err)
		return &models.ForgotPassword{}, r
	}
	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.ForgotPassword])
	if err != nil {
		r := fmt.Errorf("User not found: %w\n", err)
		return &models.ForgotPassword{}, r
	}

	return &data, err //selenjutnya generate code dan disimpan ke db table forgot_password
}

// TODO: get user by code
func (fp *ForgotPassword) GetUserByEmailCode(email string, code int) (*models.ForgotPassword, error) {
	rows, err := fp.db.Query(context.Background(), "SELECT email, code FROM forgot_password WHERE email = $1 AND code = $2", email, code)
	if err != nil {
		r := fmt.Errorf("Failed to get rows data: %w\n", err)
		return &models.ForgotPassword{}, r
	}
	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.ForgotPassword])
	if err != nil {
		r := fmt.Errorf("The Email or Code is wrong: %w\n", err)
		return &models.ForgotPassword{}, r
	}

	return &data, err //selanjutnya ubah password
}

// TODO: add method delete user by code
