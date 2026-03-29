package repository

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ForgotPasswordRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewForgotPasswordRepository(db *pgxpool.Pool, rdb *redis.Client) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{
		db:    db,
		redis: rdb,
	}
}

func (fp *ForgotPasswordRepository) GenerateOTP(email string, otp int) {
	ctx := context.Background()
	key := fmt.Sprintf("otp:%s", email)

	// Simpan OTP di Redis dengan TTL 5 menit
	err := fp.redis.Set(ctx, key, otp, 5*time.Minute).Err()
	if err != nil {
		fmt.Printf("Failed to save OTP to Redis: %v\n", err)
		// Fallback ke DB jika Redis gagal (opsional, tapi di sini kita ikuti instruksi implementasi Redis)
	}

	// Tetap simpan di DB sebagai backup atau jika sistem lain membutuhkannya (opsional)
	cmdTag, _ := fp.db.Exec(ctx, "UPDATE forgot_password SET code = $1 WHERE email = $2", otp, email)
	if cmdTag.RowsAffected() == 0 {
		fp.db.Exec(ctx, "INSERT INTO forgot_password (email, code) VALUES ($1, $2)", email, otp)
	}
}

func (fp *ForgotPasswordRepository) VerifikasiOTP(email string, otp int) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("otp:%s", email)

	// Cek di Redis dulu
	val, err := fp.redis.Get(ctx, key).Result()
	if err == nil {
		if val == fmt.Sprintf("%d", otp) {
			// Hapus OTP setelah diverifikasi agar tidak bisa dipakai lagi
			fp.redis.Del(ctx, key)
			return true, nil
		}
		return false, errors.New("Incorrect OTP")
	}

	// Jika di Redis tidak ada, cek di DB (fallback atau jika expired di Redis tapi belum di DB)
	rows, err := fp.db.Query(ctx, "SELECT email, code FROM forgot_password WHERE email = $1", email)
	if err != nil {
		return false, errors.New("Failed to get rows")
	}
	// fmt.Println("rows dari tabel forgot_password: ", rows)
	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[models.VerifOTP])
	if err != nil {
		return false, errors.New("Failed to collect row")
	}

	// fmt.Println("data dari tabel forgot_password: ", data)

	// fmt.Printf("data.Email: %v\n", data.Email)
	// fmt.Printf("Email Parameter: %v\n", email)
	// fmt.Printf("data.Code: %v\n", data.Code)
	// fmt.Printf("otp Parameter: %v\n", otp)
	if data.Email != email || data.Code != otp {
		return false, errors.New("Incorrect Email or OTP")
	}

	return true, nil
}

func (fp *ForgotPasswordRepository) ChangePassword(email string, password string) error {
	cmdTag, err := fp.db.Exec(context.Background(), "UPDATE users SET password=$1 WHERE email=$2", password, email)
	if err != nil {
		fmt.Printf("Failed to change password: %v\n", err)
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("Email not found")
	}
	return nil
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
