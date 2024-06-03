package common

import (
	"encoding/json"
	"fmt"
	nanoid "github.com/matoous/go-nanoid/v2"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Int32ToString int32转换为string
func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

// StringToInt 字符串转int
func StringToInt(str string) int {
	if IsEmptyStr(str) {
		return 0
	}
	num, _ := strconv.Atoi(str)
	return num
}

// StringToInt32 字符串转int32
func StringToInt32(str string) int32 {
	j, _ := strconv.ParseInt(str, 10, 32)

	return int32(j)

}

// ToValidUTF8 treats s as UTF-8-encoded bytes and returns a copy with each run of bytes
// representing invalid UTF-8 replaced with the bytes in replacement, which may be empty.
func ToValidUTF8(s, replacement []byte) []byte {
	b := make([]byte, 0, len(s)+len(replacement))
	invalid := false // previous byte was from an invalid UTF-8 sequence
	for i := 0; i < len(s); {
		c := s[i]
		if c < utf8.RuneSelf {
			i++
			invalid = false
			b = append(b, byte(c))
			continue
		}
		_, wid := utf8.DecodeRune(s[i:])
		if wid == 1 {
			i++
			if !invalid {
				invalid = true
				b = append(b, replacement...)
			}
			continue
		}
		invalid = false
		b = append(b, s[i:i+wid]...)
		i += wid
	}
	return b
}

// MapToJson map转换为string
func MapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

/**
 * StrVal 获取变量的字符串值, interface 转 string
 * 浮点型 3.0将会转换成字符串3, "3"
 * 非数值或字符类型的变量将会被转换成JSON格式字符串
 */

func StrVal(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func IntVal(value interface{}) int {
	return StringToInt(StrVal(value))
}

// MapVal interface转map
func MapVal(value interface{}) map[string]interface{} {
	var key map[string]interface{}
	if value == nil {
		return key
	}
	switch value := value.(type) {
	case map[string]interface{}:
		key = value
	default:
		return key
	}
	return key
}

// 将interface 转换为[]map[string]interface{}

func ArrayMapVal(value interface{}) []map[string]interface{} {
	var keys []map[string]interface{}
	switch values := value.(type) {
	case []map[string]interface{}:
		keys = values
		fmt.Printf("valuesss: %v\n", value)
	default:

		return keys
	}
	return keys
}

// ArrayVal interface转[]interface
func ArrayVal(value interface{}) []interface{} {
	var key []interface{}
	if value == nil {
		return key
	}
	switch value.(type) {
	case []interface{}:
		key = value.([]interface{})
	default:
		return key
	}
	return key
}

// NewNanoid 生成Nanoid 仅有22位
func NewNanoid() string {
	id, _ := nanoid.New()
	return id
}

// RemoveDuplicates 删除重复的元素，并且删除空元素
func RemoveDuplicates(metrics []string) []string {
	// 只需要检查key，因此value使用空结构体，不占内存
	processed := make(map[string]struct{})

	uniqUserIDs := make([]string, 0)
	for _, uid := range metrics {
		// 删除空元素
		if uid == "" || uid == "," {
			continue
		}
		if _, ok := processed[uid]; ok {
			continue
		}
		// 将唯一的用户ID加到切片中
		uniqUserIDs = append(uniqUserIDs, uid)

		// 将用户ID标记为已存在
		processed[uid] = struct{}{}

	}

	return uniqUserIDs
}

func RegexpDuplicates(regularExpression, str, flag string) []string {
	re := regexp.MustCompile(regularExpression)
	regexpExpression := re.ReplaceAllString(str, "")
	expressionList := strings.Split(regexpExpression, flag)
	newexpressionList := RemoveDuplicates(expressionList)
	return newexpressionList
}
