//@Author: wulinlin
//@Description:
//@File:  user
//@Version: 1.0.0
//@Date: 2024/02/25 18:57

package repository

import "backend/test/demo/wire_demo/repository/dao"

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(d *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: d,
	}
}
