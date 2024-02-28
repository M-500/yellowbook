package main

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-22 10:12

func main() {
	server := InitWebServer()
	server.Run(":8081")
}
