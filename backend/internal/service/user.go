package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	// 1. 将密码加密为密文
	ciphertext, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(ciphertext)
	// 2. 保存数据库
	return svc.repo.Create(ctx, user)
}

func (svc *UserService) Login(ctx context.Context, user domain.User) (domain.User, error) {
	// 1. 查找用户
	u, err := svc.repo.FindByEmail(ctx, user)
	if err != nil {
		return domain.User{}, errors.New("用户不存在")
	}
	// 2. 校验密码是否正常
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)) // 返回nil则通过，非nil则失败
	if err != nil {
		return domain.User{}, errors.New("密码输入错误")
	}
	// 3. 返回登录成功
	return u, nil
}
