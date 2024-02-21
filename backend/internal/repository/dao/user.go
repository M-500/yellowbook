package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 10:49

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, user UserModel) error {
	now := time.Now().UnixMilli() // 毫秒数
	user.Ctime = now
	user.Utime = now
	return dao.db.WithContext(ctx).Create(&user).Error
}

// UserModel
// @Description: 对标到数据库的表结构  有些人叫做entity，有些人叫做model，有些人叫做PO(persistent object)
type UserModel struct {
	Id       int64  `gorm:"primaryKey,autoIncrement;comment:主键ID"`
	Email    string `gorm:"unique;comment:邮箱，唯一索引"`
	Password string
	// 创建时间  存储毫秒数
	Ctime int64
	// 更新时间 存储毫秒数
	Utime int64
}
