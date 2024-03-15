package web

import (
	"gin-web/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 15:34

type CaptchaHandler struct {
	codeSvc service.ICaptcha
}

func NewCaptchaHandler(codeSvc service.ICaptcha) *CaptchaHandler {
	return &CaptchaHandler{
		codeSvc: codeSvc,
	}
}

func (h *CaptchaHandler) CaptchaImgAPI(ctx *gin.Context) {
	captcha, err := h.codeSvc.MakeCaptcha()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Data: captcha,
	})
	return
}
