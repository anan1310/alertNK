package common

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// EncodeMD5 MD5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// Base64Md5 先base64，然后MD5
func Base64Md5(params string) string {
	return EncodeMD5(base64.StdEncoding.EncodeToString([]byte(params)))
}

// DeBase64 base64解密
func DeBase64(encStr string) string {
	// 解密规则
	// data := "hello"
	// data = data + data[1:2] helloe
	decStr, _ := base64.URLEncoding.DecodeString(encStr)
	return string(string(decStr[0 : len(decStr)-1]))

}

// Base64 base64加密
func Base64(params string) []byte {
	return []byte(base64.RawStdEncoding.EncodeToString([]byte(params)))

}
