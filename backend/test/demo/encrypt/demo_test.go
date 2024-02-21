package encrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-02-21 14:34

func TestPwdEncrypt(t *testing.T) {
	pwdPlaintext := "abc*.123" // 密码明文

	// 加密
	ciphertext, err := bcrypt.GenerateFromPassword([]byte(pwdPlaintext), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败")
	}
	fmt.Println("加密后的密文", string(ciphertext))
	// 比较
	err = bcrypt.CompareHashAndPassword(ciphertext, []byte("abc*.123")) // 返回nil则通过，非nil则失败
	if err != nil {
		fmt.Println("密码输入错误")
	}
	fmt.Println("密码校验通过")
}
