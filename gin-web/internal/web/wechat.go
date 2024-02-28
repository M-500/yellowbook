package web

import "github.com/gin-gonic/gin"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-28 9:58

type OAuth2WechatHandler struct {
}

func NewOAuth2WechatHandler() *OAuth2WechatHandler {
	return &OAuth2WechatHandler{}
}

func (h *OAuth2WechatHandler) AuthURL(ctx *gin.Context) {

}

func (h *OAuth2WechatHandler) CallBack(ctx *gin.Context) {

}
