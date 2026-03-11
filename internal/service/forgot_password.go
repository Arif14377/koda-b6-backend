package service

import (
	"github.com/arif14377/koda-b6-backend/internal/repository"
)

type ForgotPasswordService struct {
	fpRepo   *repository.ForgotPassword
	userRepo *repository.UserRepository
}

func NewForgotPasswordService(fpRepo *repository.ForgotPassword, userRepo *repository.UserRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		fpRepo: fpRepo,
		// userRepo: userRepo,
	}
}

func (fp *ForgotPasswordService) ForgotPassword(email, password string, code int) error {
	_, err := fp.fpRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	_, err = fp.fpRepo.GetUserByEmailCode(email, code)
	if err != nil {
		return err
	}

	hashPassword := HashingPassword(password)

	err = fp.fpRepo.UpdatePassword(email, hashPassword)
	if err != nil {
		return err
	}

	err = fp.fpRepo.DeleteCode(email, code)
	if err != nil {
		return err
	}

	return err
}

// TODO: hashing password dengan argon
func HashingPassword(password string) string {
	return ""
}
