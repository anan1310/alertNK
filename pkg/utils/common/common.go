package common

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MyString struct {
	builder strings.Builder
}

// A 添加要拼接的字符串
func (ms *MyString) A(str string) *MyString {
	ms.builder.WriteString(str)
	return ms
}

// Str 生成字符串
func (ms *MyString) Str() string {
	return ms.builder.String()
}

// SubStr 字符串截取，支持字符串中含有中文
func SubStr(str string, begin int, end int) (res string) {
	if IsEmptyStr(str) {
		res = ""
		return
	}
	res = string([]rune(str)[begin:end])
	return
}

// SubIndexStr 截取"/"开头的字符串
func SubIndexStr(str, sub string) (res string) {
	if IsEmptyStr(str) {
		res = ""
		return
	}
	index := strings.Index(str, sub)

	res = str[index+1:]
	return res

}

// ToFloat64 str转float64
func ToFloat64(str string) (float64, error) {
	if IsEmptyStr(str) {
		return 0, nil
	}
	fv, err := strconv.ParseFloat(str, 64)
	return fv, err
}

// ToStr int转字符串
func ToStr(num int) string {
	return strconv.Itoa(num)
}

// ToInt string 转换int
func ToInt(str string) int {
	if IsEmptyStr(str) {
		return 0
	}
	num, _ := strconv.Atoi(str)
	return num
}

// ToBool string转bool
func ToBool(str string) bool {
	if IsEmptyStr(str) {
		return false
	}
	parseBool, _ := strconv.ParseBool(str)
	return parseBool
}

// ToInt32 字符串转 int32
func ToInt32(str string) int32 {
	if IsEmptyStr(str) {
		return 0
	}
	num, _ := strconv.Atoi(str)
	return int32(num)
}

// ContainStr 判断某字符串是否包含在字符串数组中
func ContainStr(arr []string, str string) bool {
	if !IsEmptyArr(arr) && !IsEmptyStr(str) {
		for i := 0; i < len(arr); i++ {
			if str == arr[i] {
				return true
			}
		}
	}
	return false
}

// NameFromPath 路径中提取文件名称（包含后缀）
func NameFromPath(fullPath string) string {
	if IsEmptyStr(fullPath) {
		return ""
	}
	/*
		lastIndex := strings.LastIndex(fullPath, "/")
		return string([]rune(fullPath)[lastIndex+1:])
	*/
	// 用于返回路径的最后一个元素，通常是文件名或目录名
	return filepath.Base(fullPath)
}

// GenUUID 生成UUID
func GenUUID() string {
	return uuid.New().String()
}

// IsEmptyStr 判断字符串是否为空
func IsEmptyStr(str string) bool {
	return len(strings.Trim(str, " ")) == 0
}

// Trim 去掉首尾空白
func Trim(str string) string {
	if IsEmptyStr(str) {
		return ""
	}
	return strings.Trim(str, " ")
}

// HasEmptyStr 判断参数中是否含有空字符串
func HasEmptyStr(strs ...string) bool {
	for _, str := range strs {
		if IsEmptyStr(str) {
			return true
		}
	}
	return false
}

// IsEmptyArr 判断数组是否为空
func IsEmptyArr(arr []string) bool {
	if arr == nil {
		return true
	}
	if len(arr) == 0 {
		return true
	}
	return false
}

// GetNameWithoutSuffix 获取不带后缀的文件名，比如：myWeb.tar.gz的名称为 myWeb
func GetNameWithoutSuffix(name string) string {
	if IsEmptyStr(name) {
		return ""
	}
	name = Trim(name)
	lastIndex := strings.LastIndex(name, ".")
	// 如果后缀是 ".tar.gz"，则进一步查找倒数第二个点的索引
	if strings.HasSuffix(name, ".tar.gz") {
		subName := name[:lastIndex]                 // 去掉最后一个点后的部分
		lastIndex = strings.LastIndex(subName, ".") // 查找倒数第二个点的索引
	}
	return name[:lastIndex]
}

// FixString 字符串类型断言
func FixString(obj interface{}) string {
	switch obj.(type) {
	case string:
		return obj.(string)
	default:
		return ""
	}
}

// PathWithSuffix 路径中提取文件名称（包含后缀）
func PathWithSuffix(fullPath string) string {
	if IsEmptyStr(fullPath) {
		return ""
	}
	lastIndex := strings.LastIndex(fullPath, "/")
	return string([]rune(fullPath)[lastIndex+1:])
}

// EndWith 判断字符串是否以某一子字符串结尾
func EndWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// StartWith 判断字符串是否以某一子字符串开头
func StartWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func CleanStr(str string) string {
	if IsEmptyStr(str) {
		return ""
	}
	str = Trim(str)
	str = strings.ReplaceAll(str, "\r", "")
	return str
}

// CleanBlank 使用正则清除空白
func CleanBlank(str string) string {
	// /s 相当于[\t\n\f\r ]
	re := regexp.MustCompile("\\s+")
	str = re.ReplaceAllString(str, "")
	return str
}

// DirFromPath 路径中提取目录
func DirFromPath(fullPath string) string {
	if IsEmptyStr(fullPath) {
		return ""
	}
	/*
		//如果不是目录，直接返回当前路径
		lastIndex := strings.LastIndex(fullPath, "/")
		if lastIndex == -1 {
			return fullPath
		} else {
			return string([]rune(fullPath)[0:lastIndex])
		}
	*/
	//用于返回路径的目录部分
	return filepath.Dir(fullPath)
}

func MergeMap(baseMap, patchMap map[string]string) map[string]string {
	if baseMap == nil {
		return patchMap
	}
	if patchMap == nil {
		return baseMap
	}
	for k := range baseMap {
		newValue := patchMap[k]
		if !IsEmptyStr(newValue) {
			baseMap[k] = newValue
		}
	}
	return baseMap
}

func MergeDataByMap(data string, replaceMap map[string]string) string {
	if data == "" || replaceMap == nil {
		return data
	}
	for k, v := range replaceMap {
		if !IsEmptyStr(v) {
			oldValue := new(MyString).A("${").A(k).A("}").Str()
			data = strings.ReplaceAll(data, oldValue, v)
		}

	}
	return data
}

// InterfaceToStr interface转换为String
func InterfaceToStr(param interface{}) string {
	dataType, err := json.Marshal(&param)
	if err != nil {
		fmt.Println(err)
	}
	dataString := string(dataType)
	return dataString
}

// Decimal 保留小数点后几位
func Decimal(value float64, prec int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', prec, 64), 64)
	return value
}

// ConvertToFloat64 将interface转换为float4
func ConvertToFloat64(data interface{}) float64 {
	// 将data断言为float64类型
	value, ok := data.(float64)
	if !ok {
		return 0
	}

	return value
}

func TimeUnixNano() int64 {
	return time.Now().UnixNano() / 1e6
}

// ExecuteRegular 执行正则，返回结果
func ExecuteRegular(str, pattern string) string {
	// 编译正则表达式模式
	regex := regexp.MustCompile(pattern)

	// 使用正则表达式替换字符串
	result := regex.ReplaceAllString(str, "")

	return result
}

// ExecuteIpAddress 解析IP地址
func ExecuteIpAddress(address string) string {
	// 定义匹配IP地址的正则表达式
	ipPattern := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`

	// 编译正则表达式
	re := regexp.MustCompile(ipPattern)

	// 使用正则表达式查找匹配的IP地址
	ips := re.FindAllString(address, -1)

	// 将匹配到的IP地址组成一个字符串，使用逗号分隔
	ipsString := strings.Join(ips, ",")

	return ipsString
}
