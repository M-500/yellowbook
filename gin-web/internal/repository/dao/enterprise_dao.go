package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 20:40

type IEnterpriseDao interface {
	Insert(ctx context.Context, c Enterprise) error
}

type EnterpriseDao struct {
	db *gorm.DB
}

func (dao *EnterpriseDao) Insert(ctx context.Context, c Enterprise) error {
	//  业务层保证不会出现任何的重复！ 所以我们不加唯一索引，也不做任何的去重
	err := dao.db.WithContext(ctx).Create(&c).Error
	//if me, ok := err.(*mysql.MySQLError); ok {
	//	const uniqueIndexErrNo uint16 = 1062
	//	if me.Number == uniqueIndexErrNo {
	//		return ErrUserDuplicate
	//	}
	//}
	return err
}

func NewEnterpriseDao(db *gorm.DB) IEnterpriseDao {
	return &EnterpriseDao{
		db: db,
	}
}
