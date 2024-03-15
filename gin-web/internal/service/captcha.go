package service

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 15:28

import (
	"errors"
	"gin-web/internal/domain"
	"github.com/mojocn/base64Captcha"
)

type ICaptcha interface {
	MakeCaptcha() (*domain.CaptchaResponse, error)
	CheckCaptcha(id, code string, clear bool) bool
}

type captchaService struct {
	store base64Captcha.Store
}

func NewCaptchaService() ICaptcha {
	return &captchaService{
		store: base64Captcha.DefaultMemStore,
	}
}

func (c captchaService) MakeCaptcha() (*domain.CaptchaResponse, error) {
	driver := base64Captcha.NewDriverDigit(
		80,  // 高度 png 像素高度
		240, // 宽度 png 像素高度
		5,   // 验证码默认位数
		0.7, //单个数字的最大绝对倾斜因子
		180) // 背景圆圈的数量
	cp := base64Captcha.NewCaptcha(driver, c.store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		return nil, errors.New("生成验证码失败")
	}

	return &domain.CaptchaResponse{
		CaptchaID: id,
		PicPath:   b64s,
	}, nil
}

func (c captchaService) CheckCaptcha(id, code string, clear bool) bool {
	if c.store.Verify(id, code, clear) {
		return true
	}
	return false
}
