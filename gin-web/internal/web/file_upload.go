package web

import (
	"gin-web/pkg/tools"
	"github.com/gin-gonic/gin"
	"path"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-15 17:59

type FileUploadResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

func UploadFileController(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		dst := path.Join("./static", file.Filename)
		saveErr := c.SaveUploadedFile(file, dst)
		if saveErr == nil {
			res := FileUploadResponse{
				FileName: file.Filename,
				FilePath: dst,
			}
			tools.JsonSuccessData(c, "文件上传成功", res)
		}

		tools.JsonErrorStrResp(c, saveErr.Error())
	}
	tools.JsonErrorStrResp(c, err.Error())
}
