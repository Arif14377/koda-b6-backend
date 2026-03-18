package service

import (
	"errors"
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
		err := errors.New("Email tidak valid.")
		return err
	}

	if data.FullName == "" || data.Email == "" || data.Password == "" {
		err := errors.New("Data tidak boleh kosong.")
		return err
	}

	isExist := a.userRepo.GetUserByEmail(data.Email)
	if isExist {
		err := errors.New("Email sudah terdaftar.")
		return err
	}

	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(data.Password))
	if err != nil {
		err2 := errors.New("Failed to hashing password")
		return err2
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
