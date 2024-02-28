package web

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-27 16:07

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
