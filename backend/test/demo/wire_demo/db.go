//@Author: wulinlin
//@Description:
//@File:  db
//@Version: 1.0.0
//@Date: 2024/02/25 19:12

package wire_demo

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("dsn"))
	if err != nil {
		panic(err)
	}
	return db
}
