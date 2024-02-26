package _struct

import (
	"fmt"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-26 9:41

type User struct {
	Name string
}

func TestStruct(t *testing.T) {
	a := User{Name: "琳琳"}
	b := User{Name: "琳琳"}
	fmt.Println("两个变量类型是同一个结构体，并且字段值相同，可以比较吗: ", a == b)
	c := &User{Name: "琳琳"}
	d := &User{Name: "琳琳"}
	fmt.Println("两个指针变量类型是同一个结构体，并且字段值相同，可以比较吗: ", c == d)
}
