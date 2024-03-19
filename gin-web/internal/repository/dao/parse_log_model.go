package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 21:22

type ParserLogger struct {
	gorm.Model
	FileName  string  `gorm:"comment:文件名称"`
	FileSize  float64 `gorm:"comment:文件大小，单位MB"`
	Rows      uint32  `gorm:"comment:文件有多少行"`
	IOTime    float32 `gorm:"comment:IO打开文件耗时"`
	ParseTime float32 `gorm:"comment:解析文章耗时"`
	WaitTime  float32 `gorm:"comment:总耗时"`
}
