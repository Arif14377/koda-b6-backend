package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
	"github.com/matthewhartstonge/argon2"
)

type ForgotPasswordService struct {
	fpRepo   *repository.ForgotPasswordRepository
	userRepo *repository.UserRepository
}

func NewForgotPasswordService(fpRepo *repository.ForgotPasswordRepository, userRepo *repository.UserRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		fpRepo:   fpRepo,
		userRepo: userRepo,
	}
}

func (fp *ForgotPasswordService) GenerateOTP(email string) {

	otp, _ := rand.Int(rand.Reader, big.NewInt(1000000))

	otpValue := int(otp.Int64())
	fp.fpRepo.GenerateOTP(email, otpValue)
	fmt.Printf("Kode OTP Anda: %v\n", otpValue)
}

func (fp *ForgotPasswordService) VerifikasiOTP(email string, otp int) error {
	isTrue, err := fp.fpRepo.VerifikasiOTP(email, otp)

	if !isTrue {
		return err
	}

	return nil
}

func (fp *ForgotPasswordService) ChangePassword(data *models.ForgotPassword) error {

	if data.NewPassword != data.ConfirmPassword {
		err := errors.New("Password tidak sama.")
		return err
	}

	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(data.NewPassword))
	if err != nil {
		r := fmt.Errorf("Failed to hashing password: %w", err)
		return r
	}

	err = fp.fpRepo.ChangePassword(data.Email, string(encoded))
	if err != nil {
		return err
	}
	return nil
}

// func (fp *ForgotPasswordService) ForgotPassword(email, password string, code int) error {
// 	userExist := fp.userRepo.GetUserByEmail(email)
// 	if userExist {
// 		otp, _ := rand.Int(rand.Reader, big.NewInt(1000000))

// 	}

// 	_, err := fp.fpRepo.GetUserByEmailCode(email, code)
// 	if err != nil {
// 		return err
// 	}

// 	var argon argon2.Config
// 	hashPassword, err := HashingPassword(password, argon)
// 	if err != nil {
// 		return err
// 	}

// 	err = fp.fpRepo.UpdatePassword(email, hashPassword)
// 	if err != nil {
// 		return err
// 	}

// 	err = fp.fpRepo.DeleteCode(email, code)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

// // TODO: hashing password dengan argon
// func HashingPassword(password string, argon argon2.Config) (string, error) {
// 	encoded, err := argon.HashEncoded([]byte(password))
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(encoded), nil
// }
