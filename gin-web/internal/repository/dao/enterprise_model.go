package dao

import (
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 17:52
//

type Enterprise struct {
	gorm.Model
	CompanyName               string     `gorm:"column:company_name;comment:企业名称;not null;type:varchar(200);"`
	RegStatus                 string     `gorm:"column:reg_status;comment:登记状态;type:varchar(32);"`
	LegalBoss                 string     `gorm:"column:legal_boss;comment:法定代表人;type:varchar(64);"`
	RegCapital                string     `gorm:"column:reg_capital;comment:注册资本;type:varchar(64);"`
	RegDate                   *time.Time `gorm:"column:reg_date;comment:成立日期;"`
	SocialCode                string     `gorm:"column:social_code;comment:统一社会信用代码;not null;type:varchar(32);"`
	CompanyRegAddr            string     `gorm:"column:company_reg_addr;comment:企业注册地址;type:varchar(255);"`
	Phone                     string     `gorm:"column:phone;comment:电话;type:varchar(32);"`
	MorePhone                 string     `gorm:"column:more_phone;comment:更多电话;type:varchar(255);"`
	Email                     string     `gorm:"column:email;comment:邮箱;type:varchar(100);"`
	MoreEmail                 string     `gorm:"column:more_email;comment:更多邮箱;type:varchar(255);"`
	Province                  string     `gorm:"column:province;comment:所属省份;type:varchar(64);"`
	City                      string     `gorm:"column:city;comment:所属城市;type:varchar(100);"`
	District                  string     `gorm:"column:district;comment:所属区县;type:varchar(100);"`
	TaxNumber                 string     `gorm:"column:tax_number;comment:纳税人识别号;type:varchar(32);"`
	RegNumber                 string     `gorm:"column:reg_number;comment:注册号;type:varchar(32);"`
	OrgCode                   string     `gorm:"column:org_code;comment:组织机构代码;type:varchar(32);"`
	InsuredNumber             int        `gorm:"column:insured_number;comment:参保人数;type:varchar(32);"`
	InsuredNumberReport       string     `gorm:"column:insured_number_report;comment:参保人数所属年报;type:varchar(32);"`
	CompanyType               string     `gorm:"column:company_type;comment:企业(机构)类型;type:varchar(32);"`
	CompanySize               string     `gorm:"column:company_size;comment:企业规模;type:varchar(32);"`
	ApprovalDate              *time.Time `gorm:"column:approval_date;comment:核准日期;type:varchar(32);"`
	BusinessTerm              string     `gorm:"column:business_term;comment:营业期限;type:varchar(64);"`
	GBIndustryCategory        string     `gorm:"column:gb_industry_category;comment:国标行业门类;type:varchar(100);"`
	GBIndustryMajorCategory   string     `gorm:"column:gb_industry_major_category;comment:国标行业大类;type:varchar(100);"`
	GBIndustryMediumCategory  string     `gorm:"column:gb_industry_medium_category;comment:国标行业中类;type:varchar(100);"`
	GBIndustrySmallCategory   string     `gorm:"column:gb_industry_small_category;comment:国标行业小类;type:varchar(100);"`
	FormerName                string     `gorm:"column:former_name;comment:曾用名;type:varchar(255);"`
	EnglishName               string     `gorm:"column:english_name;comment:英文名;type:varchar(200);"`
	OfficialWebsite           string     `gorm:"column:official_website;comment:官网网址;type:varchar(255);"`
	CommunicationAdd          string     `gorm:"column:communication_add;comment:通信地址;type:varchar(500);"`
	CompanyProfile            string     `gorm:"column:company_profile;comment:企业简介;type:varchar(500);"`
	BusinessScope             string     `gorm:"column:business_scope;comment:经营范围;type:varchar(500);"`
	QCCIndustryCategory       string     `gorm:"column:qcc_industry_category;comment:企查查行业门类;type:varchar(100);"`
	QCCIndustryMajorCategory  string     `gorm:"column:qcc_industry_major_category;comment:企查查行业大类;type:varchar(100);"`
	QCCIndustryMediumCategory string     `gorm:"column:qcc_industry_medium_category;comment:企查查行业中类;type:varchar(100);"`
	QCCIndustrySmallCategory  string     `gorm:"column:qcc_industry_small_category;comment:企查查行业小类;type:varchar(100);"`
}
