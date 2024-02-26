package service

import (
	"gin-web/internal/repository"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:15

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
