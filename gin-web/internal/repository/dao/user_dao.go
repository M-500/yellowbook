package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:50

// ErrDataNotFound 通用的数据没找到
var ErrDataNotFound = gorm.ErrRecordNotFound

// ErrUserDuplicate 这个算是 user 专属的
var ErrUserDuplicate = errors.New("用户邮箱或者手机号冲突")

type UserDaoInterface interface {
	Insert(ctx context.Context, u UserModel) error
	UpdateNonZeroFields(ctx context.Context, u UserModel) error
	FindByPhone(ctx context.Context, phone string) (UserModel, error)
	FindByEmail(ctx context.Context, email string) (UserModel, error)
	FindByWechat(ctx context.Context, openId string) (UserModel, error)
	FindById(ctx context.Context, id int64) (UserModel, error)
}

type UserDAO struct {
	db *gorm.DB
}

func (u2 UserDAO) Insert(ctx context.Context, u UserModel) error {
	err := u2.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const uniqueIndexErrNo uint16 = 1062
		if me.Number == uniqueIndexErrNo {
			return ErrUserDuplicate
		}
	}
	return err
}

func (u2 UserDAO) UpdateNonZeroFields(ctx context.Context, u UserModel) error {
	//TODO implement me
	panic("implement me")
}

func (u2 UserDAO) FindByPhone(ctx context.Context, phone string) (UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u2 UserDAO) FindByEmail(ctx context.Context, email string) (UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u2 UserDAO) FindByWechat(ctx context.Context, openId string) (UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u2 UserDAO) FindById(ctx context.Context, id int64) (UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}
