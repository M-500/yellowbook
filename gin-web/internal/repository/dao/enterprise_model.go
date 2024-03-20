package dao

import (
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 17:52

type Enterprise struct {
	gorm.Model
	CompanyName               string     `gorm:"comment:企业名称;type:;size:;"`
	RegStatus                 string     `gorm:"comment:登记状态;type:;size:;"`
	LegalBoss                 string     `gorm:"comment:;type:;size:;"`
	RegCapital                string     `gorm:"comment:;type:;size:;"`
	RegDate                   *time.Time `gorm:"comment:;type:;size:;"`
	SocialCode                string     `gorm:"comment:;type:;size:;"`
	CompanyRegAddr            string     `gorm:"comment:;type:;size:;"`
	Phone                     string     `gorm:"comment:;type:;size:;"`
	MorePhone                 string     `gorm:"comment:;type:;size:;"`
	Email                     string     `gorm:"comment:;type:;size:;"`
	MoreEmail                 string     `gorm:"comment:;type:;size:;"`
	Province                  string     `gorm:"comment:;type:;size:;"`
	City                      string     `gorm:"comment:;type:;size:;"`
	District                  string     `gorm:"comment:;type:;size:;"`
	TaxNumber                 string     `gorm:"comment:;type:;size:;"`
	RegNumber                 string     `gorm:"comment:;type:;size:;"`
	OrgCode                   string     `gorm:"comment:;type:;size:;"`
	InsuredNumber             int        `gorm:"comment:;type:;size:;"`
	InsuredNumberReport       string     `gorm:"comment:;type:;size:;"`
	CompanyType               string     `gorm:"comment:;type:;size:;"`
	CompanySize               string     `gorm:"comment:;type:;size:;"`
	ApprovalDate              *time.Time `gorm:"comment:;type:;size:;"`
	BusinessTerm              string     `gorm:"comment:;type:;size:;"`
	GBIndustryCategory        string     `gorm:"comment:;type:;size:;"`
	GBIndustryMajorCategory   string     `gorm:"comment:;type:;size:;"`
	GBIndustryMediumCategory  string     `gorm:"comment:;type:;size:;"`
	GBIndustrySmallCategory   string     `gorm:"comment:;type:;size:;"`
	FormerName                string     `gorm:"comment:;type:;size:;"`
	EnglishName               string     `gorm:"comment:;type:;size:;"`
	OfficialWebsite           string     `gorm:"comment:;type:;size:;"`
	CommunicationAdd          string     `gorm:"comment:;type:;size:;"`
	CompanyProfile            string     `gorm:"comment:;type:;size:;"`
	BusinessScope             string     `gorm:"comment:;type:;size:;"`
	QCCIndustryCategory       string     `gorm:"comment:;type:;size:;"`
	QCCIndustryMajorCategory  string     `gorm:"comment:;type:;size:;"`
	QCCIndustryMediumCategory string     `gorm:"comment:;type:;size:;"`
	QCCIndustrySmallCategory  string     `gorm:"comment:;type:;size:;"`
}
