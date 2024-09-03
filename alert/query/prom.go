package query

import (
	"alarm_collector/internal/models"
	"alarm_collector/pkg/utils/common"
	"alarm_collector/pkg/utils/expression_util"
	"errors"
	"fmt"
	"strings"
)

// operatorPrecedence 定义运算符优先级映射
var operatorPrecedence = map[string]int{
	"(":  0,
	")":  0,
	"||": 1,
	"&&": 2,
	">":  3,
}

// isOperator 检查子表达式是否是运算符
func isOperator(exp string) bool {
	_, ok := operatorPrecedence[exp]
	return ok
}

// evaluateRPN 计算逆波兰表达式的值
func evaluateRPN(alarmRule string, metricsMap map[string]interface{}) (bool, string, error) {
	var stack []bool
	var severity string
	//
	expressions := toRPN(alarmRule)

	for _, expr := range expressions {
		if !isOperator(expr) {
			// 如果是操作数，入栈
			switch expr {
			case "true":
				stack = append(stack, true)
			case "false":
				stack = append(stack, false)
			default:
				processOperator(expr, &severity, metricsMap, &stack)
			}
		} else {
			// 如果是运算符，进行运算
			switch expr {
			case "&&":
				operand2 := stack[len(stack)-1]
				operand1 := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, operand1 && operand2)
			case "||":
				operand2 := stack[len(stack)-1]
				operand1 := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, operand1 || operand2)
			}
		}
	}

	if len(stack) != 1 {
		return false, "", errors.New("逆波兰表达式计算出现错误")
	}
	return stack[0], severity, nil
}

// toRPN 将中缀布尔表达式转换为逆波兰表达式
func toRPN(expression string) []string {
	var output []string    // 输出队列
	var operators []string // 操作符栈

	// 使用逗号分割表达式
	tokens := strings.Split(expression, ",")
	for _, token := range tokens {
		switch {
		case token == "(":
			operators = append(operators, token)
		case token == ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = operators[:len(operators)-1]
		case isOperator(token):
			for len(operators) > 0 && operatorPrecedence[operators[len(operators)-1]] >= operatorPrecedence[token] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		default:
			output = append(output, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	return output
}

// 处理操作符，计算栈顶的两个表达式的结果，并将结果入栈
func processOperator(rule string, severity *string, metricsMap map[string]interface{}, conditionStack *[]bool) {
	// 解析告警规则
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
	case "!=":
		conditionMet = dataSource != alarmValue

	default:
		fmt.Println("无效的规则:", rule)
	}
	if conditionMet {
		*conditionStack = append(*conditionStack, true)
	} else {
		delete(metricsMap, targetMapping)
		*conditionStack = append(*conditionStack, false)
	}
}

// parseRule 解析告警规则，获取参数、比较符和值
func parseRule(rule string, severity *string) (paramStr string, operator string, alarmValue float64, err error) {
	// 修剪输入字符串，去除首尾多余的空格
	rule = strings.TrimSpace(rule)
	// 分割规则字符串
	parts := strings.Fields(rule)
	// 提取参数、比较符和值
	paramStr = strings.TrimSpace(parts[0])
	operator = strings.TrimSpace(parts[1])
	valueStr := strings.TrimSpace(parts[2])
	if len(parts) == 4 {
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

// 处理告警源
func (rq *RuleQuery) handleAlertSource(rule models.AlertRule) (map[string]interface{}, string, error) {
	var (
		alarmRule         = new(common.MyString)
		targetMapping     = new(common.MyString)
		rules             = rule.PrometheusConfig.Rules
		alertSourceMap    = rule.PrometheusConfig.AlertSource
		isUnionRule       = rule.PrometheusConfig.IsUnionRule
		complexExpression = expression_util.TransformExpression(rule.PrometheusConfig.ComplexExpression)
	)
	size := len(rules)
	for i, r := range rules {
		switch isUnionRule {
		case 0:
			alarmRule.A(r.TargetKey + " " + r.Operator + " " + common.StrVal(r.Value) + " " + r.Severity)
			targetMapping.A(r.TargetMapping)
			if i < size-1 {
				targetMapping.A(",")
				alarmRule.A(",||,")
			}
		case 1:
			alarmRule.A(r.TargetKey + " " + r.Operator + " " + common.StrVal(r.Value) + " " + r.Severity)
			targetMapping.A(r.TargetMapping)
			if i < size-1 {
				targetMapping.A(",")
				alarmRule.A(",&&,")
			}
		case 2:
			//( 1 AND 2 ) OR ( 3 AND 4 )
			complexExpression = strings.ReplaceAll(complexExpression, fmt.Sprintf("\\u%04X", i+1), fmt.Sprintf("%s %s %s %s", r.TargetKey, r.Operator, common.StrVal(r.Value), r.Severity))
			targetMapping.A(r.TargetMapping)
			if i < size-1 {
				targetMapping.A(",")
			}
		}
	}
	//获取数据源的值  如果达到告警的阈值 那么就写入redis缓冲中
	s := models.PrometheusDataSourceQuery{
		MetricType: alertSourceMap["metricType"],
		MetricName: alertSourceMap["metricName"],
		MetricHost: alertSourceMap["metricHost"],
		//Pid:           alertSourceMap["pid"],
		TenantId:      rule.TenantId,
		TargetMapping: targetMapping.Str(),
	}
	//获取告警源
	oldAlertSource, err := rq.ctx.CK.PrometheusDataSource().Get(s)
	if err != nil {
		return nil, "", err
	}
	parameters := make(map[string]interface{})
	newAlertSource := make(map[string]interface{})
	for _, r := range rules {
		targetFields := strings.Split(r.TargetMapping, ",")
		for _, field := range targetFields {
			if value, exists := oldAlertSource[field]; exists {
				parameters[field] = value
			}
		}
		var result interface{}
		if !common.IsEmptyStr(r.TargetExpression) {
			//执行表达式
			result = expression_util.EvaluateExpression(r.TargetExpression, parameters)
		} else {
			result = oldAlertSource[r.TargetMapping]
		}
		//单位换算
		result = expression_util.HandlerConvertMetricUnit(result, r.FromUnit, r.ToUnit, r.Precision)
		newAlertSource[r.TargetKey] = result
	}
	if isUnionRule == 2 {
		alarmRule.A(complexExpression)
	}
	return newAlertSource, alarmRule.Str(), nil
}

/*
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
*/
