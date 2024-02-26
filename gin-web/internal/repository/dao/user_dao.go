package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:50

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
	//TODO implement me
	panic("implement me")
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
