package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 10:05

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	// 1. 加密要放在哪里？

	// 2. 查询数据库 并存储
	return svc.repo.Create(ctx, user)
}
