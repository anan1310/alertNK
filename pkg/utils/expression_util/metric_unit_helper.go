package expression_util

import (
	"alarm_collector/constant"
	"fmt"
	"strconv"
)

// convertMetricUnit 用于将指标值从原始单位转换为目标单位，保留指定的精度
func convertMetricUnit(value float64, fromUnit, toUnit string, precision int) (float64, error) {

	// 检查输入的单位是否存在
	if _, ok := constant.UnitConversion[fromUnit]; !ok {
		return 0, fmt.Errorf("原始单位 %s 不存在", fromUnit)
	}
	if _, ok := constant.UnitConversion[toUnit]; !ok {
		return 0, fmt.Errorf("目标单位 %s 不存在", toUnit)
	}

	// 通过字典检查单位是否属于同一个类别
	fromCategory, toCategory := constant.Categories[fromUnit], constant.Categories[toUnit]

	if fromCategory != toCategory {
		return handleSpecialUnits(fromCategory, toCategory, fromUnit, toUnit, value, precision)
	} else {
		//处理相同类型的单位
		return unitsConvert(fromUnit, toUnit, value, precision)
	}

}

// roundToPrecision 将浮点数四舍五入到指定的精度
func roundToPrecision(convertedValue float64, precision int) float64 {
	// 格式化为指定的精度
	//formattedValue := fmt.Sprintf(fmt.Sprintf("%%.%df", precision), convertedValue)

	// 转换为 float64 类型
	//roundedValue, _ := strconv.ParseFloat(formattedValue, 64)
	// 保留指定的精度
	roundedValue, _ := strconv.ParseFloat(strconv.FormatFloat(convertedValue, 'f', precision, 64), 64)
	return roundedValue

}

// 处理特殊单位
func handleSpecialUnits(fromCategory, toCategory, fromUnit, toUnit string, value float64, precision int) (float64, error) {

	switch {
	case fromCategory == "0" && toCategory == "2":
		// Bytes转换为KiByte
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "2" && toCategory == "0":
		//KiBytes转换为Bytes
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "1" && toCategory == "3":
		//Bytes/s转换为KiBytes/s
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "4" && toCategory == "6":
		//Bit转换为KiBit
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "6" && toCategory == "4":
		//KiBit转换为Bit
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "5" && toCategory == "7":
		//Bit/s转换为KBit/s
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "7" && toCategory == "5":
		//KBit/s转换为Bit/s
		return unitsConvert(fromUnit, toUnit, value, precision)
	case fromCategory == "0" && toCategory == "4":
		//Bytes转换为bit 1Bytes=8bit
		return bytesConvertToBit(fromUnit, toUnit, value, precision)
	// 可以添加其他特殊转换情况
	case fromCategory == "4" && toCategory == "0":
		//Bytes转换为bit 1Bytes=8bit
		return bitConvertToBytes(fromUnit, toUnit, value, precision)
	// 可以添加其他特殊转换情况
	default:
		return 0, fmt.Errorf("不支持从 %s 到 %s 的转换", fromUnit, toUnit)
	}
}

// 单位转换
func unitsConvert(fromUnit, toUnit string, value float64, precision int) (float64, error) {
	//统一处理单位
	convertedValue := value * constant.UnitConversion[fromUnit] / constant.UnitConversion[toUnit]
	return roundToPrecision(convertedValue, precision), nil
}

// Bytes转换为bit
func bytesConvertToBit(fromUnit, toUnit string, value float64, precision int) (float64, error) {
	//Bytes-bit
	convertedValue := 8 * value * constant.UnitConversion[fromUnit] / constant.UnitConversion[toUnit]
	return roundToPrecision(convertedValue, precision), nil
}

// bit转换为Bytes
func bitConvertToBytes(fromUnit, toUnit string, value float64, precision int) (float64, error) {
	//bit-Bytes
	convertedValue := 1 / 8 * value * constant.UnitConversion[fromUnit] / constant.UnitConversion[toUnit]
	return roundToPrecision(convertedValue, precision), nil
}
