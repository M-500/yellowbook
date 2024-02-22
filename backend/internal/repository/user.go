package repository

import (
	"backend/internal/domain"
	"backend/internal/repository/cache"
	"backend/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 10:45

var RepoErrUserDuplicateEmail = dao.ErrUserDuplicateEmail

type UserRepository struct {
	userDao *dao.UserDAO
	cache   *cache.UserCache
}

func NewUserRepository(userDao *dao.UserDAO, c *cache.UserCache) *UserRepository {
	return &UserRepository{
		userDao: userDao,
		cache:   c,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.userDao.Insert(ctx, dao.UserModel{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, u domain.User) (domain.User, error) {
	email, err := r.userDao.FindByEmail(ctx, u.Email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       email.Id,
		Email:    email.Email,
		Password: email.Password,
	}, err
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	// 先从cache中查找
	user, err := r.cache.GetUser(ctx, id)
	if err == nil {
		// 必然有数据
		return user, err
	}
	if err == cache.ErrKeyNotExist {
		// 数据必然不存在
		// 去数据库加载
	}
	// 如果这里 err = io.EOF 怎么办？要不要去数据库加载？缓存击穿
	// 有没有可能 redis整个都崩了？

	// 选择加载 ==> 做好兜底，万一Redis真崩了，你要保护好数据库 限流（内存限流） 用ORM的middleware

	// 选择不加载 => 用户体验不好
	byId, err := r.userDao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	userEntiry := domain.User{
		Id:       id,
		Email:    byId.Email,
		Password: byId.Password,
	}
	go func() {
		// 载入缓存
		err = r.cache.SetUser(ctx, userEntiry)
		if err != nil {
			// 这里我怎么办？  其实打日志就好了做监控
			//return userEntiry, err
			//return userEntiry, nil
		}
	}()
	return userEntiry, nil
}

//  用缓存的时候 一定要注意的两个问题 1. 数据一致性问题 2. 缓存崩溃的问题

// 1. 两个集群 二级缓存
// 2. redis + 本地缓存
// 3. 高配redis+ 低配redis切换

// 不管用几级缓存，一定保护好你的数据库
