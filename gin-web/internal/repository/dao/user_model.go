package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:50

type UserModel struct {
	gorm.Model
	Phone        string `gorm:"type:string;size:11;index;default '';comment:手机号码," json:"phone"`
	UserName     string `gorm:"type:string;size:32;not null;unique;comment:昵称" json:"user_name"`
	Password     string `gorm:"type:string;size:128;comment:密码" json:"password"`
	CreatorId    int    `gorm:"type:string;not null;comment:创建人ID" json:"creator_id"`
	Gender       string `gorm:"type:string;size:1;comment:性别" json:"gender"`
	Active       bool   `gorm:"type:tinyint;not null;default 0;comment:状态,是否已经禁用" json:"active"`
	IsAdmin      bool   `gorm:"type:tinyint;not null;default 0;comment:是否是管理员 0不是，1是，默认0" json:"is_admin"`
	CoverImgLink string `gorm:"type:string;size:255;comment:头像链接" json:"cover_img_link"`
}

func (UserModel) TableName() string {
	return "user"
}
