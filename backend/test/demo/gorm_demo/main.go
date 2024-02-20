//@Author: wulinlin
//@Description:
//@File:  main
//@Version: 1.0.0
//@Date: 2024/02/20 21:38

package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Pirce int
}

func main() {
	//db, err := gorm.Open(mysql.Open("admin:123456@tcp(127.0.0.1:3306)/test_db"),&gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DryRun: true,
	})

	db = db.Debug()
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移schema

	//建表
	db.AutoMigrate(&Product{})

	// 插入数据
	db.Create(&Product{Code: "A11", Pirce: 43})

	// 查询数据
	var p Product
	db.First(&p, 1)
	db.First(&p, "code = ?", "A11")

	// 更新
	db.Model(&p).Update("Price", 210)

	db.Model(&p).Updates(Product{Pirce: 210, Code: "B12"})                    // 仅仅更新非零字段
	db.Model(&p).Updates(map[string]interface{}{"Price": 210, "Code": "B11"}) // 仅仅更新非零字段

	// 删除
	db.Delete(&p, 1)
}
