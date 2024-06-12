package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//告警信息生成指纹

func GenerateFingerprint_(input string) string {
	// 获取当前时间戳
	timestamp := time.Now().UnixNano()

	// 将输入和时间戳拼接起来
	data := fmt.Sprintf("%s:%d", input, timestamp)

	// 计算MD5哈希值
	hash := md5.New()
	hash.Write([]byte(data))
	fingerprint := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串，并截取前16个字符
	return hex.EncodeToString(fingerprint)[:16]
}

func GenerateFingerprint(ruleID string) string {
	// 计算MD5哈希值
	hash := md5.New()
	hash.Write([]byte(ruleID))
	fingerprint := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串，并截取前16个字符
	return hex.EncodeToString(fingerprint)[:16]
}
