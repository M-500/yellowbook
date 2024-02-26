package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:50

type UserModel struct {
	gorm.Model
	UserName string
	Phone    string
	Password string
	Birthday string
}

func (UserModel) TableName() string {
	return "user"
}
