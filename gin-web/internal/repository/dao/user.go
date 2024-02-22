package dao

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:16

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	NickName string `json:"nickName"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}
