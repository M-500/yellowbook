package service

import (
	"context"
	"errors"
	"gin-web/internal/domain"
	"gin-web/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:15
type IUserService interface {
	Signup(ctx context.Context, u domain.DMUser) error
	FindOrCreate(ctx context.Context, phone string) (domain.DMUser, error)
	Login(ctx context.Context, email, password string) (domain.DMUser, error)
	Profile(ctx context.Context, id int64) (domain.DMUser, error)
	// UpdateNonSensitiveInfo 更新非敏感数据
	// 你可以在这里进一步补充究竟哪些数据会被更新
	UpdateNonSensitiveInfo(ctx context.Context, user domain.DMUser) error

	// FindOrCreateByWechat 查找或者初始化
	// 随着业务增长，这边可以考虑拆分出去作为一个新的 Service
	FindOrCreateByWechat(ctx context.Context, info domain.WechatInfo) (domain.DMUser, error)
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
func (u2 *UserService) Signup(ctx context.Context, u domain.DMUser) error {
	// 1. 密码加密
	ciphertext, err := bcrypt.GenerateFromPassword([]byte(u.Pwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Pwd = string(ciphertext)
	return u2.userRepo.Create(ctx, u)
}

func (u2 *UserService) FindOrCreate(ctx context.Context, phone string) (domain.DMUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u2 *UserService) Login(ctx context.Context, email, password string) (domain.DMUser, error) {
	// 查找用户
	u, err := u2.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return domain.DMUser{}, errors.New("用户不存在")
	}
	// 校验密码是否正常
	err = bcrypt.CompareHashAndPassword([]byte(u.Pwd), []byte(password))
	if err != nil {
		return domain.DMUser{}, errors.New("密码输入错误")
	}
	// 登录成功 返回用户
	return u, nil
}

func (u2 *UserService) Profile(ctx context.Context, id int64) (domain.DMUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u2 *UserService) UpdateNonSensitiveInfo(ctx context.Context, user domain.DMUser) error {
	//TODO implement me
	panic("implement me")
}

func (u2 *UserService) FindOrCreateByWechat(ctx context.Context, info domain.WechatInfo) (domain.DMUser, error) {
	//TODO implement me
	panic("implement me")
}
