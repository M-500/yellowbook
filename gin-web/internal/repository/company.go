package repository

import (
	"context"
	"errors"
	"fmt"
	"gin-web/internal/domain"
	"gin-web/internal/repository/cache"
	"gin-web/internal/repository/dao"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-19 20:39

type ICompany interface {
	Create(ctx context.Context, company domain.Company) error
	CreateLog(ctx context.Context, log domain.ParserLogger) error
}

type CompanyRepo struct {
	cache     cache.CompanyCache
	dao       dao.IEnterpriseDao
	daoParser dao.IParseLog
}

func (repo *CompanyRepo) key(taxCode string) string {
	return fmt.Sprintf("company:uni:%s", taxCode[2:8])
}

func (repo *CompanyRepo) DoToModel(c domain.Company) dao.Enterprise {
	enterprise := dao.Enterprise{
		CompanyName:    c.EnterpriseName,
		RegNumber:      c.RegistrationNumber,
		LegalBoss:      c.LegalRepresentative,
		RegStatus:      c.RegistrationStatus,
		RegCapital:     c.RegisteredCapital,
		CompanyRegAddr: c.EnterpriseRegistrationAddress,
		Phone:          c.Phone,
		MorePhone:      c.MorePhone,
		Email:          c.Email,
		MoreEmail:      c.MoreEmail,
		//RegDate:     time.Parse("2006-01-02", c.EstablishmentDate), // 成立日期
		//RegDate: c.EstablishmentDate,
		SocialCode:  c.UnifiedSocialCreditCode,
		Province:    c.Province,
		City:        c.City,
		District:    c.District,
		TaxNumber:   c.TaxpayerIdentificationNumber,
		OrgCode:     c.OrganizationCode,
		CompanyType: c.EnterpriseType,
		CompanySize: c.EnterpriseSize,
	}
	parse, err := time.Parse("2006-01-02", c.EstablishmentDate)
	if err == nil {
		enterprise.RegDate = &parse
	}
	return enterprise
}
func (repo *CompanyRepo) CreateLog(ctx context.Context, l domain.ParserLogger) error {

	return repo.daoParser.Insert(ctx, dao.ParserLogger{
		Mark:      l.Mark,
		FileName:  l.FileName,
		FileSize:  l.FileSize,
		Rows:      l.Rows,
		IOTime:    l.IOTime,
		ParseTime: l.ParseTime,
		WaitTime:  l.WaitTime,
	})
}
func (repo *CompanyRepo) Create(ctx context.Context, c domain.Company) error {
	// 这一层去重嘛？ 还是说 在dao层去重？
	//key := repo.key(c.UnifiedSocialCreditCode)
	key := "company:uni:V1"
	exist, err := repo.cache.Exist(ctx, key, c.UnifiedSocialCreditCode)
	if err != nil {
		// TODO 报错了怎么办？ 基本上是网络错误之类的吧 只能做监控啦还能怎么办
	}
	if !exist {
		err = repo.cache.Add(ctx, key, c.UnifiedSocialCreditCode)
		if err != nil {
			return err
		}
		return repo.dao.Insert(ctx, repo.DoToModel(c))
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

func NewCompanyRepo(dao dao.IEnterpriseDao, cache cache.CompanyCache, daoParser dao.IParseLog) ICompany {
	return &CompanyRepo{
		dao:       dao,
		cache:     cache,
		daoParser: daoParser,
	}
}
