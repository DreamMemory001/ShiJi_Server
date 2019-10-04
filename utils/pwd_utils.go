package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
	"time"
)

/**
密码处理逻辑
*/
func PwdEncode(pwdString string) (string) {
	pwdBytes := []byte(pwdString)
	var salt = make([]byte, 8)
	binary.BigEndian.PutUint64(salt,
		uint64(time.Now().Unix())) // 盐，是一个随机字符串，每一个用户都不一样，在这里我们随机选择 "I1lrI7wqJOJZ" 作为盐
	iterations := 1000            // 加密算法的迭代次数，10000 次
	digest := sha256.New           // digest 算法，使用 sha256

	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key(pwdBytes, salt, iterations, 32, digest)

	// 第二步：Base64 编码
	str := base64.StdEncoding.EncodeToString(dk)

	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	return "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
}

func PwdDecode(pwdString string) (realPwdString string) {
}
