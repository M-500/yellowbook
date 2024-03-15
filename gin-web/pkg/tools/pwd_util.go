// Package tools
// Date        : 2023/2/15 20:49
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :
package tools

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

// GeneratePwd
// @Description: 创建密码
// @param plaintext
// @return ciphertext
// @return err
func GeneratePwd(plaintext string) (ciphertext string, err error) {
	if len(strings.TrimSpace(plaintext)) == 0 {
		return "", errors.New("明文不能为空")
	}
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(plaintext, options)
	ciphertext = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	return
}

// CheckPwd
// @Description: 校验密码
// @return res
func CheckPwd(oldPassword, checkPassword string) bool {
	//options := &password.Options{16, 100, 32, sha512.New}
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	if !strings.Contains(oldPassword, "$") {
		return false
	}
	passwordInfo := strings.Split(oldPassword, "$")
	return password.Verify(checkPassword, passwordInfo[2], passwordInfo[3], options)
}
