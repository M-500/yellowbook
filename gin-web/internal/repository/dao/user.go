package dao

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:16

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UserName string
	Phone    string
	Email    string
	Pwd      string
}
