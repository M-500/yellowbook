package dao

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 17:52

type Enterprise struct {
	ID                            uint      `gorm:"primaryKey" json:"id"`
	EnterpriseName                string    `gorm:"type:varchar(255);comment:企业名称" json:"enterprise_name"`
	RegistrationStatus            string    `gorm:"type:varchar(255);comment:登记状态" json:"registration_status"`
	LegalRepresentative           string    `gorm:"type:varchar(255);comment:法定代表人" json:"legal_representative"`
	RegisteredCapital             string    `gorm:"type:varchar(255);comment:注册资本" json:"registered_capital"`
	EstablishmentDate             time.Time `gorm:"comment:成立日期" json:"establishment_date"`
	UnifiedSocialCreditCode       string    `gorm:"type:varchar(255);comment:统一社会信用代码" json:"unified_social_credit_code"`
	EnterpriseRegistrationAddress string    `gorm:"type:varchar(255);comment:企业注册地址" json:"enterprise_registration_address"`
	Phone                         string    `gorm:"type:varchar(255);comment:电话" json:"phone"`
	MorePhone                     string    `gorm:"type:varchar(255);comment:更多电话" json:"more_phone"`
	Email                         string    `gorm:"type:varchar(255);comment:邮箱" json:"email"`
	MoreEmail                     string    `gorm:"type:varchar(255);comment:更多邮箱" json:"more_email"`
	Province                      string    `gorm:"type:varchar(255);comment:所属省份" json:"province"`
	City                          string    `gorm:"type:varchar(255);comment:所属城市" json:"city"`
	District                      string    `gorm:"type:varchar(255);comment:所属区县" json:"district"`
	TaxpayerIdentificationNumber  string    `gorm:"type:varchar(255);comment:纳税人识别号" json:"taxpayer_identification_number"`
	RegistrationNumber            string    `gorm:"type:varchar(255);comment:注册号" json:"registration_number"`
	OrganizationCode              string    `gorm:"type:varchar(255);comment:组织机构代码" json:"organization_code"`
	InsuredNumber                 int       `gorm:"comment:参保人数" json:"insured_number"`
	InsuredNumberAnnualReport     string    `gorm:"type:varchar(255);comment:参保人数所属年报" json:"insured_number_annual_report"`
	EnterpriseType                string    `gorm:"type:varchar(255);comment:企业(机构)类型" json:"enterprise_type"`
	EnterpriseSize                string    `gorm:"type:varchar(255);comment:企业规模" json:"enterprise_size"`
	ApprovalDate                  time.Time `gorm:"comment:核准日期" json:"approval_date"`
	BusinessTerm                  string    `gorm:"type:varchar(255);comment:营业期限" json:"business_term"`
	GBIndustryCategory            string    `gorm:"type:varchar(255);comment:国标行业门类" json:"gb_industry_category"`
	GBIndustryMajorCategory       string    `gorm:"type:varchar(255);comment:国标行业大类" json:"gb_industry_major_category"`
	GBIndustryMediumCategory      string    `gorm:"type:varchar(255);comment:国标行业中类" json:"gb_industry_medium_category"`
	GBIndustrySmallCategory       string    `gorm:"type:varchar(255);comment:国标行业小类" json:"gb_industry_small_category"`
	FormerName                    string    `gorm:"type:varchar(255);comment:曾用名" json:"former_name"`
	EnglishName                   string    `gorm:"type:varchar(255);comment:英文名" json:"english_name"`
	OfficialWebsite               string    `gorm:"type:varchar(255);comment:官网网址" json:"official_website"`
	CommunicationAddress          string    `gorm:"type:varchar(255);comment:通信地址" json:"communication_address"`
	EnterpriseProfile             string    `gorm:"type:text;comment:企业简介" json:"enterprise_profile"`
	BusinessScope                 string    `gorm:"type:text;comment:经营范围" json:"business_scope"`
	QCCIndustryCategory           string    `gorm:"type:varchar(255);comment:企查查行业门类" json:"qcc_industry_category"`
	QCCIndustryMajorCategory      string    `gorm:"type:varchar(255);comment:企查查行业大类" json:"qcc_industry_major_category"`
	QCCIndustryMediumCategory     string    `gorm:"type:varchar(255);comment:企查查行业中类" json:"qcc_industry_medium_category"`
	QCCIndustrySmallCategory      string    `gorm:"type:varchar(255);comment:企查查行业小类" json:"qcc_industry_small_category"`
}
