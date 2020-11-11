package number

import (
	"strconv"
)

// ParseInteger 字符串转换成为int64
func ParseInteger(str string) (int64, bool) {
	num, err := strconv.ParseInt(str, 10, 64)
	return num, err == nil
}

// ParseFloat 字符串转换成为float64
func ParseFloat(str string) (float64, bool) {
	num, err := strconv.ParseFloat(str, 64)
	return num, err == nil
}
