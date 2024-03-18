package web

import (
	"gin-web/internal/service"
	"gin-web/pkg/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 18:06

type ExcelHandler struct {
	svc *service.ExcelParserService
}

func NewExcelHandler(s *service.ExcelParserService) *ExcelHandler {
	return &ExcelHandler{
		svc: s,
	}
}

func (h *ExcelHandler) ParseExcel(c *gin.Context) {
	type Req struct {
		FileName string `json:"file_name"`
		FilePath string `json:"file_path"`
	}
	// 校验参数
	var form Req
	if err := c.Bind(&form); err != nil {
		c.String(http.StatusUnauthorized, "数据不合法")
		return
	}
	// 校验路径是否合法
	err := h.svc.ParserExcel(form.FilePath)
	if err != nil {
		tools.JsonErrorStrResp(c, err.Error())
		return
	}
	tools.JsonSuccessData(c, "解析成功", nil)
}
