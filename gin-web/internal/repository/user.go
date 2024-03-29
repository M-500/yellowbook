package repository

import (
	"context"
	"gin-web/internal/domain"
	"gin-web/internal/repository/dao"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:53

type IUserRepository interface {
	Create(ctx context.Context, u domain.DMUser) error
	// Update 更新数据，只有非 0 值才会更新
	Update(ctx context.Context, u domain.DMUser) error
	FindByPhone(ctx context.Context, phone string) (domain.DMUser, error)
	FindByUserName(ctx context.Context, username string) (domain.DMUser, error)
	FindByEmail(ctx context.Context, email string) (domain.DMUser, error)
	FindById(ctx context.Context, id int64) (domain.DMUser, error)
	// FindByWechat 暂时可以认为按照 openId来查询
	// 将来可能需要按照 unionId 来查询
	FindByWechat(ctx context.Context, openId string) (domain.DMUser, error)
}

type UserRepository struct {
	userDao dao.IUserDao
}

func NewUserRepository(userDao dao.IUserDao) IUserRepository {
	return &UserRepository{
		userDao: userDao,
	}
}

func (u2 *UserRepository) Create(ctx context.Context, u domain.DMUser) error {
	return u2.userDao.Insert(ctx, dao.UserModel{
		UserName: u.Username,
		Phone:    u.Phone,
		Password: u.Pwd,
	})
}

func (u2 *UserRepository) Update(ctx context.Context, u domain.DMUser) error {
	//TODO implement me
	panic("implement me")
}

func (u2 *UserRepository) FindByUserName(ctx context.Context, username string) (domain.DMUser, error) {
	byEmail, err := u2.userDao.FindByUsername(ctx, username)
	if err != nil {
		return domain.DMUser{}, err
	}
	return domain.DMUser{
		Id:       byEmail.ID,
		Pwd:      byEmail.Password,
		Username: byEmail.UserName,
		Phone:    byEmail.Phone,
	}, nil
}

func (u2 *UserRepository) FindByPhone(ctx context.Context, phone string) (domain.DMUser, error) {
	byEmail, err := u2.userDao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.DMUser{}, err
	}
	return domain.DMUser{
		Id:       byEmail.ID,
		Pwd:      byEmail.Password,
		Username: byEmail.UserName,
		Phone:    byEmail.Phone,
	}, nil
}

func (u2 *UserRepository) FindByEmail(ctx context.Context, email string) (domain.DMUser, error) {
	byEmail, err := u2.userDao.FindByEmail(ctx, email)
	if err != nil {
		return domain.DMUser{}, err
	}
	return domain.DMUser{
		Id:       byEmail.ID,
		Pwd:      byEmail.Password,
		Username: byEmail.UserName,
		Phone:    byEmail.Phone,
	}, nil
}

func (u2 *UserRepository) FindById(ctx context.Context, id int64) (domain.DMUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u2 *UserRepository) FindByWechat(ctx context.Context, openId string) (domain.DMUser, error) {
	//TODO implement me
	panic("implement me")
}
