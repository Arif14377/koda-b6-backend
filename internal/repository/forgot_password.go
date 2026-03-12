package repository

import (
	"context"
	"errors"
	"math/big"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ForgotPasswordRepository struct {
	db *pgx.Conn
}

func NewForgotPasswordRepository(db *pgx.Conn) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{
		db: db,
	}
}

func (fp *ForgotPasswordRepository) GenerateOTP(email string, otp *big.Int) {
	cmdTag, _ := fp.db.Exec(context.Background(), "UPDATE forgot_password SET code = $1 WHERE email = $2", otp, email)
	if cmdTag.RowsAffected() > 0 {
		return
	}

	fp.db.Exec(context.Background(), "INSERT INTO forgot_password (email, code) VALUES ($1, $2)", email, otp)
}

func (fp *ForgotPasswordRepository) VerifikasiOTP(email string, otp *big.Int) (bool, error) {
	rows, err := fp.db.Query(context.Background(), "SELECT email, code FROM forgot_password WHERE email = $1", email)
	if err != nil {
		return false, errors.New("Failed to get rows")
	}
	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.VerifOTP])

	if data.Email != email || data.Code != otp {
		return false, errors.New("Incorrect Email or OTP")
	}

	return true, nil
}

// 1. Kirim email
// 2. Get user by email (check if exists)
// 3. Mendapatkan kode (generated code disimpan di table forgot_password bersama email)
// 4. Input kode
// 5. Kirim email dan kode ke server (request forgot password)
// 6. CEK apakah sesuai? (di tabel forgot_password)
// 7. Jika sesuai maka ubah password di table users.
// 8. Hapus code di table forgot_password

// func (fp *ForgotPassword) GetUserByEmailCode(email string, code int) (*models.ForgotPassword, error) {
// 	rows, err := fp.db.Query(context.Background(), "SELECT email, code FROM forgot_password WHERE email = $1 AND code = $2", email, code)
// 	if err != nil {
// 		r := fmt.Errorf("Failed to get rows data: %w\n", err)
// 		return &models.ForgotPassword{}, r
// 	}
// 	defer rows.Close()

// 	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.ForgotPassword])
// 	if err != nil {
// 		r := fmt.Errorf("The Email or Code is wrong: %w\n", err)
// 		return &models.ForgotPassword{}, r
// 	}

// 	return &data, nil //selanjutnya ubah password
// }

// func (fp *ForgotPassword) UpdatePassword(email, hashPassword string) error {
// 	_, err := fp.db.Exec(context.Background(), "UPDATE users SET password = $1 WHERE email = $2", hashPassword, email)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (fp *ForgotPassword) DeleteCode(email string, code int) error {
// 	_, err := fp.db.Exec(context.Background(), "UPDATE forgot_password SET code = NULL WHERE email = $1 AND code = $2", email, code)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
