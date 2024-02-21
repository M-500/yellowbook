package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 12:15

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&UserModel{})
}
