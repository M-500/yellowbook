//@Author: wulinlin
//@Description:
//@File:  wire
//@Version: 1.0.0
//@Date: 2024/02/25 19:04

//go:build wireinject

// 让wire来注入这里的代码
package wire_demo

import (
	"backend/test/demo/wire_demo/repository"
	"backend/test/demo/wire_demo/repository/dao"
	"github.com/google/wire"
)

func InitRepo() *repository.UserRepository {
	// 这里面传入各个组件的初始化方法 顺序无所谓
	wire.Build(dao.NewUserDao, repository.NewUserRepository, InitDb)
	return new(repository.UserRepository)
}
