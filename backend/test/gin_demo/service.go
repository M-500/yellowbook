package main

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:18

type UserService struct {
	userDao *UserDAO
}

func NewUserService(userDao *UserDAO) *UserService {
	return &UserService{
		userDao: userDao,
	}
}

func (svc *UserService) FindByUserName(u string) interface{} {
	return svc.userDao.FindByUserName(u)
}

func (svc *UserService) CreateUser(u User) error {
	return svc.userDao.CreateUser(u)
}
