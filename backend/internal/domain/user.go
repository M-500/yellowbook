package domain

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 10:07

// User
// @Description: 用户本身就是一个领域对象 对应的是DDD中的聚合根(entity)
// 有些系统也叫做 BO(Business Object)
type User struct {
	Email    string
	Password string
	//ConfirmPwd string 这个密码就不需要了，因为他不需要传递给service层
}
