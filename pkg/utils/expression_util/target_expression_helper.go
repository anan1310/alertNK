package expression_util

import (
	"alarm_collector/global"
	"alarm_collector/pkg/utils/common"
	"github.com/Knetic/govaluate"
	"go.uber.org/zap"
	"regexp"
)

// EvaluateExpression 解析并执行表达式
func EvaluateExpression(expressionStr string, parameters map[string]interface{}) interface{} {
	// 使用正则表达式提取表达式并替换掉特殊字符
	expressionPattern := `@\[|]@`
	re := regexp.MustCompile(expressionPattern)
	expressionStr = re.ReplaceAllString(expressionStr, "")

	expression, err := govaluate.NewEvaluableExpression(expressionStr)
	if err != nil {
		// 记录错误到日志中
		global.Logger.Error("解析表达式错误_1：", zap.Error(err))
		return 0
	}

	result, err := expression.Evaluate(parameters)
	if err != nil || result == "NAN" {
		global.Logger.Error("解析表达式错误_2：", zap.Error(err))
		//单位转换失败
		return 0
	}

	return result
}

// HandlerConvertMetricUnit 转换单位 为了统一precision设置为字符串
func HandlerConvertMetricUnit(value interface{}, fromUnit, toUnit, precision string) interface{} {
	if common.IsEmptyStr(precision) {
		//默认精度保留两位小数
		precision = "2"
	}
	toPrecision := common.ToInt(precision)
	//如果不是float64类型 返回当前值
	coverValue, ok := value.(float64)
	if !ok {
		return value
	}
	//以下情况不需要转换 ：1：默认单位和转换单位相同并且不为空 2、默认单位和转换单位为"" 不需要进行转换
	if (fromUnit == toUnit && fromUnit != "" && toUnit != "") || (fromUnit == "" && toUnit == "") {
		//计算精度
		return roundToPrecision(coverValue, toPrecision)

	}
	//单位转换：如果错误返回原来的value
	convertedValue, err := convertMetricUnit(coverValue, fromUnit, toUnit, toPrecision)
	if err != nil {
		return value
	}

	return convertedValue
}
