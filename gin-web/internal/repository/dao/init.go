package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 17:16

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(
		&UserModel{},
		&Enterprise{},
		&ParserLogger{},
	)
}
