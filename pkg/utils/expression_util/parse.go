package expression_util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func TransformExpression(expression string) string {
	// 替换 AND 为 &&
	expression = strings.ReplaceAll(expression, "AND", "&&")
	// 替换 OR 为 ||
	expression = strings.ReplaceAll(expression, "OR", "||")
	// 去除空格
	expression = strings.ReplaceAll(expression, " ", ",")
	// 使用正则表达式找到数字并替换为Unicode转义序列
	re := regexp.MustCompile(`(\d+)`)
	unicodeStr := re.ReplaceAllStringFunc(expression, func(s string) string {
		num, _ := strconv.Atoi(s)
		return fmt.Sprintf("\\u%04X", num) // 使用 %04X 格式化为Unicode转义序列
	})
	return unicodeStr
}
