package repository

import (
	"context"
	"gin-web/internal/domain"
	"gin-web/internal/repository/dao"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 20:39

type ICompany interface {
	Create(ctx context.Context, company domain.Company) error
}

type CompanyRepo struct {
	dao dao.IEnterpriseDao
}

func (repo *CompanyRepo) Create(ctx context.Context, c domain.Company) error {
	// 这一层去重嘛？ 还是说 在dao层去重？
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

func NewCompanyRepo(dao dao.IEnterpriseDao) ICompany {
	return &CompanyRepo{
		dao: dao,
	}
}
