package main

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:18

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) FindByUserName(u string) interface{} {
	var user User
	err := dao.db.Where("username = ?", u).First(&user).Error
	if err != nil {
		return err
	}
	return user
}

func (dao *UserDAO) CreateUser(u User) error {
	err := dao.db.Create(User{
		Username: u.Username,
		Pwd:      u.Pwd,
	}).Error
	return err
}
