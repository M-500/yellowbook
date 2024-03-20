package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 21:22

type ParserLogger struct {
	gorm.Model
	FileName  string  `gorm:"comment:文件名称"`
	FileSize  float64 `gorm:"comment:文件大小，单位MB"`
	Rows      uint32  `gorm:"comment:文件有多少行"`
	IOTime    int64   `gorm:"comment:IO打开文件耗时"`
	ParseTime int64   `gorm:"comment:解析文章耗时"`
	WaitTime  int64   `gorm:"comment:总耗时"`
	Mark      string  `gorm:"comment:标签"`
}

type IParseLog interface {
	Insert(ctx context.Context, l ParserLogger) error
}

type ParseLogDao struct {
	db *gorm.DB
}

func (dao *ParseLogDao) Insert(ctx context.Context, l ParserLogger) error {
	return dao.db.WithContext(ctx).Create(&l).Error
}

func NewParseLogDao(db *gorm.DB) IParseLog {
	return &ParseLogDao{
		db: db,
	}
}
