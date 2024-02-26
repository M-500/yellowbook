package domain

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 18:16

import "github.com/golang-jwt/jwt/v5"

type DMUser struct {
	Id       uint
	Email    string
	Pwd      string
	Username string
	Phone    string
	BirthDay string
}

type WechatInfo struct {
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明自己需要放入token里的数据
	UserId int64
}
