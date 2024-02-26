package gin_demo

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
	return nil
}
