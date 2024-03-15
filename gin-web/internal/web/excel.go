package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 18:06

type ExcelHandler struct {
}

func (h *ExcelHandler) ParseExcel(c gin.Context) {
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

	// 解析Excel

	// 返回值

}
