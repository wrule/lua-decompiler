package number

import (
	"strconv"
)

// ParseInteger 字符串转换成为int64
// 如果字符串完全符合int64规范（没有小数点）则返回标志为成功
func ParseInteger(str string) (int64, bool) {
	num, err := strconv.ParseInt(str, 10, 64)
	return num, err == nil
}

// ParseFloat 字符串转换成为float64
func ParseFloat(str string) (float64, bool) {
	num, err := strconv.ParseFloat(str, 64)
	return num, err == nil
}

// TryParseInteger 尝试把字符串转换成为int64
// 此函数在需要时也会尝试进行float64转换
func TryParseInteger(str string) (int64, bool) {
	if num, ok := ParseInteger(str); ok {
		return num, true
	}
	if num, ok := ParseFloat(str); ok {
		return FloatToInteger(num)
	}
	return 0, false
}
