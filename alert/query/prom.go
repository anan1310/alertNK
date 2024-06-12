package query

import (
	"alarm_collector/pkg/utils/common"
	"fmt"
	"strings"
)

// operatorPrecedence 定义操作符优先级
var (
	operatorPrecedence = map[string]int{
		"&&": 1,
		"||": 0,
		">":  2,
		"<":  2,
		">=": 2,
		"<=": 2,
		"=":  2,
	}
)

func ParsePromRule(alarmRule string, valueMap map[string]interface{}) (error, []bool, string) {
	// 2、解析告警规则
	conditionStack := make([]bool, 0)  // 用于存储表达式结果的栈
	operatorStack := make([]string, 0) // 用于存储操作符的栈

	var severity string

	// 遍历规则字符串，解析出表达式和操作符
	for _, condition := range strings.Split(alarmRule, ",") { //condition： system_cpu_usage < 0.1
		if p, ok := operatorPrecedence[condition]; ok { //operatorPrecedence[condition] 返回操作符对应的int类型
			// 如果是操作符，则根据优先级弹出栈中的元素进行计算
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1] // 取出操作符栈顶元素
				if q, ok := operatorPrecedence[top]; ok && q >= p {
					// 如果操作符栈顶元素的优先级大于等于当前操作符，弹出栈顶元素，并计算栈顶的两个表达式的值
					operatorStack = operatorStack[:len(operatorStack)-1] // 弹出操作符栈顶元素
					//保证可以获取到数据，防止数组越界问题
					if len(conditionStack) < 2 {
						break
					}
					right := conditionStack[len(conditionStack)-1]          // 取出右操作数
					left := conditionStack[len(conditionStack)-2]           // 取出左操作数
					conditionStack = conditionStack[:len(conditionStack)-2] // 弹出左右操作数
					// 根据操作符计算左右操作数的结果，并将结果入栈
					switch top {
					case "&&":
						//if !left || !right {
						//	return nil, append(conditionStack, false) // 短路求值，停止计算
						//}
						conditionStack = append(conditionStack, left && right)
					case "||":
						//if left || right {
						//	return nil, append(conditionStack, true) // 短路求值，停止计算
						//}
						conditionStack = append(conditionStack, left || right)
					}
				} else { //优先级小于，结束当前循环
					break
				}
			}
			operatorStack = append(operatorStack, condition)
		} else {
			// 如果是比较运算符，则将表达式的结果压入栈中
			_, conditionStack = recordAlarm(condition, &severity, valueMap, conditionStack)

		}
	}

	// 处理剩余的操作符,比如 ||
	for len(operatorStack) > 0 && len(conditionStack) >= 2 {
		top := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		right := conditionStack[len(conditionStack)-1]
		left := conditionStack[len(conditionStack)-2]
		conditionStack = conditionStack[:len(conditionStack)-2]
		switch top {
		case "&&":
			conditionStack = append(conditionStack, left && right)
		case "||":
			conditionStack = append(conditionStack, left || right)
		}
	}
	// 如果最终条件为真，则触发告警
	return nil, conditionStack, severity
}

// 记录告警信息
func recordAlarm(rule string, severity *string, metricsMap map[string]interface{}, conditionStack []bool) (error, []bool) {
	//解析告警规则
	targetMapping, operator, alarmValue, _ := parseRule(rule, severity)

	currentValue, ok := metricsMap[targetMapping]
	if !ok {
		fmt.Println("没有找到")
	}
	// 告警数据源
	dataSource := common.ConvertToFloat64(currentValue)
	// 判断是否满足告警规则
	conditionMet := false
	switch operator {
	case ">":
		conditionMet = dataSource > alarmValue
	case "<":
		conditionMet = dataSource < alarmValue
	case "=":
		conditionMet = dataSource == alarmValue
	case ">=":
		conditionMet = dataSource >= alarmValue
	case "<=":
		conditionMet = dataSource <= alarmValue
	default:
		fmt.Println("无效的规则:", rule)
	}
	if conditionMet {
		conditionStack = append(conditionStack, true)
	} else {
		delete(metricsMap, targetMapping)
		conditionStack = append(conditionStack, false)
	}
	return nil, conditionStack
}

// parseRule 解析告警规则，获取参数、比较符和值
func parseRule(rule string, severity *string) (paramStr string, operator string, alarmValue float64, err error) {
	// 修剪输入字符串，去除首尾多余的空格
	rule = strings.TrimSpace(rule)

	// 分割规则字符串
	parts := strings.Fields(rule)
	//if len(parts) != 3 {
	//	err = fmt.Errorf("无效的规则 '%s'，格式应为<参数> <比较符> <值>", rule)
	//	return
	//}

	// 提取参数、比较符和值
	paramStr = strings.TrimSpace(parts[0])
	operator = strings.TrimSpace(parts[1])
	valueStr := strings.TrimSpace(parts[2])
	if len(parts) == 4 {
		//severityTmp := strings.TrimSpace(parts[3])
		*severity = strings.TrimSpace(parts[3])
	}
	// 转换 alarmValue 为 float64 类型
	alarmValue, err = common.ToFloat64(valueStr)
	if err != nil {
		err = fmt.Errorf("无法将值 '%s' 转换为 float64 类型的数值: %v", valueStr, err)
		return
	}

	return paramStr, operator, alarmValue, nil
}
