package service

import (
	"gin-web/internal/repository"
	"gin-web/internal/service/sms"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 19:39

type CodeService struct {
	sms  sms.ISMSService
	repo *repository.CodeRepository
}

func NewCodeService(sms sms.ISMSService, repo *repository.CodeRepository) *CodeService {
	return &CodeService{}
}
