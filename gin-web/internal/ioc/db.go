package ioc

import (
	"gin-web/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 13:46

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("admin:123456@tcp(192.168.1.52:3306)/xhs?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 初始化数据库
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
