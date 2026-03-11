package service

import (
	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetAllUser() *[]models.UserListRead {
	return u.repo.GetAllUser()
}

func (u *UserService) GetUserByEmail(email string) bool {
	return u.repo.GetUserByEmail(email)
}
