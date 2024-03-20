package repository

import (
	"context"
	"errors"
	"fmt"
	"gin-web/internal/domain"
	"gin-web/internal/repository/cache"
	"gin-web/internal/repository/dao"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 20:39

type ICompany interface {
	Create(ctx context.Context, company domain.Company) error
}

type CompanyRepo struct {
	cache cache.CompanyCache
	dao   dao.IEnterpriseDao
}

func (repo *CompanyRepo) key(taxCode string) string {
	return fmt.Sprintf("company:uni:%s", taxCode[2:8])
}
func (repo *CompanyRepo) Create(ctx context.Context, c domain.Company) error {
	// 这一层去重嘛？ 还是说 在dao层去重？
	key := repo.key(c.UnifiedSocialCreditCode)
	exist, err := repo.cache.Exist(ctx, key, c.UnifiedSocialCreditCode)
	if err != nil {
		// TODO 报错了怎么办？ 基本上是网络错误之类的吧 只能做监控啦还能怎么办
	}
	if !exist {
		err = repo.cache.Add(ctx, key, c.UnifiedSocialCreditCode)
		if err != nil {
			return err
		}
		return repo.dao.Insert(ctx, dao.Enterprise{
			CompanyName: c.EnterpriseName,
			RegNumber:   c.RegistrationNumber,
			LegalBoss:   c.LegalRepresentative,
			//RegDate: c.EstablishmentDate,
			SocialCode: c.UnifiedSocialCreditCode,
			Province:   c.Province,
			City:       c.City,
		})
	}
	fmt.Println(exist, c.UnifiedSocialCreditCode)
	return errors.New("税号已经存在")
	//// 存在的话 要不要更新？？？？
	//return repo.dao.Insert(ctx, dao.Enterprise{
	//	CompanyName: c.EnterpriseName,
	//	RegNumber:   c.RegistrationNumber,
	//	LegalBoss:   c.LegalRepresentative,
	//	//RegDate: c.EstablishmentDate,
	//	SocialCode: c.UnifiedSocialCreditCode,
	//	Province:   c.Province,
	//	City:       c.City,
	//})
}

func NewCompanyRepo(dao dao.IEnterpriseDao, cache cache.CompanyCache) ICompany {
	return &CompanyRepo{
		dao:   dao,
		cache: cache,
	}
}
