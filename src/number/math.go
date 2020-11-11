package number

import "math"

// IntegerFloorDiv 实现Lua的整数整除
func IntegerFloorDiv(a, b int64) int64 {
	// 不必向下取整
	if a > 0 && b > 0 ||
		a < 0 && b < 0 ||
		a%b == 0 {
		return a / b
	}
	return a/b - 1
}

// FloatFloorDiv 实现Lua的浮点数整除
func FloatFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

// IntegerMod 实现Lua的整数取模运算
func IntegerMod(a, b int64) int64 {
	return a - IntegerFloorDiv(a, b)*b
}

// FloatMod 实现Lua的浮点数取模运算
func FloatMod(a, b float64) float64 {
	return a - FloatFloorDiv(a, b)*b
}

// ShiftLeft 实现Lua的位左移运算
func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	}
	return ShiftRight(a, -n)
}

// ShiftRight 实现Lua的位右移运算
func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		return int64(uint64(a) >> uint64(n))
	}
	return ShiftLeft(a, -n)
}
