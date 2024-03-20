package repository

import (
	"context"
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

func (repo *CompanyRepo) Create(ctx context.Context, c domain.Company) error {
	// 这一层去重嘛？ 还是说 在dao层去重？
	exist, err := repo.cache.Exist(ctx, c.UnifiedSocialCreditCode)
	if err != nil {
		// TODO 报错了怎么办？
	}
	if !exist {
		return repo.dao.Insert(ctx, dao.Enterprise{
			EnterpriseName:      c.EnterpriseName,
			RegNumber:           c.RegistrationNumber,
			LegalRepresentative: c.LegalRepresentative,
			//RegDate: c.EstablishmentDate,
			SocialCode: c.UnifiedSocialCreditCode,
			Province:   c.Province,
			City:       c.City,
		})
	}
	// 存在的话 要不要更新？？？？
	return repo.dao.Insert(ctx, dao.Enterprise{
		EnterpriseName:      c.EnterpriseName,
		RegNumber:           c.RegistrationNumber,
		LegalRepresentative: c.LegalRepresentative,
		//RegDate: c.EstablishmentDate,
		SocialCode: c.UnifiedSocialCreditCode,
		Province:   c.Province,
		City:       c.City,
	})
}

func NewCompanyRepo(dao dao.IEnterpriseDao, cache cache.CompanyCache) ICompany {
	return &CompanyRepo{
		dao:   dao,
		cache: cache,
	}
}
