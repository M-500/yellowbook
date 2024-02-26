package main

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 16:53

type User struct {
	gorm.Model
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}
