package repository

import (
	"backend/internal/domain"
	"backend/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 10:45

type UserRepository struct {
	userDao *dao.UserDAO
}

func NewUserRepository(userDao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		userDao: userDao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.userDao.Insert(ctx, dao.UserModel{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(int642 int64) {
	// 先从cache中查找
	// 再从Dao中查询
	// 找到结果返回
}
