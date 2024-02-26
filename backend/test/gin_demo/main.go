package gin_demo

import "github.com/gin-gonic/gin"

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 15:17

func main() {
	server := gin.Default()

	server.Run(":8001")
}
