package api

import (
	"math"

	"../number"
)

var (
	IADD  = func(a, b int64) int64 { return a + b }
	FADD  = func(a, b float64) float64 { return a + b }
	ISUB  = func(a, b int64) int64 { return a - b }
	FSUB  = func(a, b float64) float64 { return a - b }
	IMUL  = func(a, b int64) int64 { return a * b }
	FMUL  = func(a, b float64) float64 { return a * b }
	IMOD  = number.IntegerMod
	FMOD  = number.FloatMod
	POW   = math.Pow
	DIV   = func(a, b float64) float64 { return a / b }
	IIDIV = number.IntegerFloorDiv
	FIDIV = number.FloatFloorDiv
	BAND  = func(a, b int64) int64 { return a & b }
	BOR   = func(a, b int64) int64 { return a | b }
	BXOR  = func(a, b int64) int64 { return a ^ b }
	SHL   = number.ShiftLeft
	SHR   = number.ShiftRight
	IUNM  = func(a, _ int64) int64 { return -a }
	FUNM  = func(a, _ float64) float64 { return -a }
	BNOT  = func(a, _ int64) int64 { return ^a }
)

// Operator 整合整数和浮点数运算函数的结构体
type Operator struct {
	IntegerFunc func(int64, int64) int64
	FloatFunc   func(float64, float64) float64
}

// Operators 运算映射表
var Operators = []Operator{
	Operator{IADD, FADD},
	Operator{ISUB, FSUB},
	Operator{IMUL, FMUL},
	Operator{IMOD, FMOD},
	Operator{nil, POW},
	Operator{nil, DIV},
	Operator{IIDIV, FIDIV},
	Operator{BAND, nil},
	Operator{BOR, nil},
	Operator{BXOR, nil},
	Operator{SHL, nil},
	Operator{SHR, nil},
	Operator{IUNM, FUNM},
	Operator{BNOT, nil},
}
