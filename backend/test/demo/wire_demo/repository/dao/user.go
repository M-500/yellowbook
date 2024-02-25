//@Author: wulinlin
//@Description:
//@File:  user
//@Version: 1.0.0
//@Date: 2024/02/25 18:55

package dao

import "gorm.io/gorm"

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}
