package repository

import "gin-web/internal/repository/dao"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:53

type UserRepository struct {
	userDao *dao.UserDAO
}

func NewUserRepository(userDao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		userDao: userDao,
	}
}
