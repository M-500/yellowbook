package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 10:49

var (
	ErrUserDuplicateEmail = errors.New("邮箱已经存在")
)

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
	err := dao.db.WithContext(ctx).Create(&user).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (UserModel, error) {
	u := UserModel{}
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	return u, err
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

//func (UserModel) TableName() string {
//	return "user_info"
//}
