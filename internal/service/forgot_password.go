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

// TODO:
func (fp *ForgotPasswordService) ForgotPassword(email string, code int) error {
	_, err := fp.fpRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	_, err = fp.fpRepo.GetUserByEmailCode(email, code)
	if err != nil {
		return err
	}

	// TODO: panggil method repo ganti password

	// TODO: panggil method repo hapus

	return err
}
