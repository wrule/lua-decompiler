package api

import (
	"math"

	"../number"
)

// Arith 在栈上执行算术和位运算
func (me *LuaState) Arith(op EArithOpType) {
	a := me.stack.Pop()
	b := a
	dstOp := Operators[op]
	if dstOp.OperandNums == 2 {
		b = me.stack.Pop()
	}
	result := arith(a, b, dstOp)
	if result.Type() != LuaTypeNil {
		me.stack.Push(result)
	} else {
		panic("算术或位运算执行错误")
	}
}

// arith 拉倒吧，虚拟机的运算搞那么复杂
func arith(a, b LuaValue, op Operator) LuaValue {
	// 如果是位运算，则尝试进行整数转换而后求值
	if op.FloatFunc == nil {
		if x, ok := a.ToIntegerX(); ok {
			if y, ok := b.ToIntegerX(); ok {
				return LuaValue{
					vtype: LuaTypeInteger,
					value: op.IntegerFunc(x, y),
				}
			}
		}
	} else {
		if op.IntegerFunc != nil {
			if x, ok := a.ToIntegerX(); ok {
				if y, ok := b.ToIntegerX(); ok {
					return LuaValue{
						vtype: LuaTypeInteger,
						value: op.IntegerFunc(x, y),
					}
				}
			}
		}
		if op.FloatFunc != nil {
			if x, ok := a.ToNumberX(); ok {
				if y, ok := b.ToNumberX(); ok {
					return LuaValue{
						vtype: LuaTypeNumber,
						value: op.FloatFunc(x, y),
					}
				}
			}
		}
	}
	return LuaValue{
		vtype: LuaTypeNil,
		value: nil,
	}
}

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
	// 整型运算方法
	IntegerFunc func(int64, int64) int64
	// 浮点型运算方法
	FloatFunc func(float64, float64) float64
	// 操作数个数
	OperandNums int
}

// Operators 运算映射表
var Operators = []Operator{
	Operator{IADD, FADD, 2},
	Operator{ISUB, FSUB, 2},
	Operator{IMUL, FMUL, 2},
	Operator{IMOD, FMOD, 2},
	Operator{nil, POW, 2},
	Operator{nil, DIV, 2},
	Operator{IIDIV, FIDIV, 2},
	Operator{BAND, nil, 2},
	Operator{BOR, nil, 2},
	Operator{BXOR, nil, 2},
	Operator{SHL, nil, 2},
	Operator{SHR, nil, 2},
	Operator{IUNM, FUNM, 1},
	Operator{BNOT, nil, 1},
}
