package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
)

type AuthService struct {
	authRepo *repository.AuthRepository
	userRepo *repository.UserRepository
}

func NewAuthService(authRepo *repository.AuthRepository, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (a *AuthService) Register(data *models.UserRegister) error {
	if !strings.Contains(data.Email, "@") {
		return errors.New("Email tidak valid.")
	}

	if data.FullName == "" || data.Email == "" || data.Password == "" {
		return errors.New("Data tidak boleh kosong.")
	}

	isExist := a.userRepo.GetUserByEmail(data.Email)
	if isExist {
		return errors.New("Email sudah terdaftar.")
	}

	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(data.Password))
	if err != nil {
		return fmt.Errorf("Failed to hashing password: %w", err)
	}

	data.Password = string(encoded)
	userID := uuid.NewString()

	err = a.authRepo.Register(userID, data)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Login(email, password string) (*models.UserLogin, error) {
	if !strings.Contains(email, "@") {
		err := errors.New("Email wrong.")
		return &models.UserLogin{}, err
	}

	if email == "" || password == "" {
		err := errors.New("Email or Password is empty.")
		return &models.UserLogin{}, err
	}

	user, err := a.authRepo.Login(email, password)

	return user, err
}
