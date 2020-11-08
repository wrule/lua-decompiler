package vm

import "fmt"

// Instruction 指令
type Instruction uint32

// Opcode 获取指令操作码
func (me Instruction) Opcode() EInstruction {
	return EInstruction(me & 0x3f)
}

// ABC 获取ABC模式下的参数（6，8，9，9）
func (me Instruction) ABC() (a, b, c int) {
	a = int(me >> 6 & 0xff)
	b = int(me >> 14 & 0x1ff)
	c = int(me >> 23)
	return a, b, c
}

// ABx 获取ABx模式下的参数（6，8，18）
func (me Instruction) ABx() (a, bx int) {
	a = int(me >> 6 & 0xff)
	bx = int(me >> 14)
	return a, bx
}

// MaxArgBx （比如说255）
const MaxArgBx = 1<<18 - 1

// MaxArgSBx （比如说127）
const MaxArgSBx = MaxArgBx >> 1

// AsBx 获取AsBx模式下的参数（6，8，18）
func (me Instruction) AsBx() (a, sbx int) {
	a, bx := me.ABx()
	// 比如[0 ~ 126][127][128 ~ 255]，可以表示127个负数和128个正数，一个0
	return a, bx - MaxArgSBx
}

// Ax 获取Ax模式下的参数（6，26）
func (me Instruction) Ax() int {
	return int(me >> 6)
}

// OpName 获取指令名称
func (me Instruction) OpName() string {
	return InstructionInfos[me.Opcode()].name
}

// Operands 获取指令操作数信息
func (me Instruction) Operands() string {
	switch me.OpMode() {
	case IABC:
		a, b, c := me.ABC()
		var outText = fmt.Sprintf("%d", a)
		if me.BMode() != InstructionArgN {
			// 9位数的最高位为1的话，视为常量表索引，低8位按负数输出
			if b > 0xff {
				outText += fmt.Sprintf(" %d", -1-(b&0xff))
			} else {
				outText += fmt.Sprintf(" %d", b)
			}
		}
		if me.CMode() != InstructionArgN {
			if c > 0xff {
				outText += fmt.Sprintf(" %d", -1-(c&0xff))
			} else {
				outText += fmt.Sprintf(" %d", c)
			}
		}
		return fmt.Sprintf("%s [ABC]", outText)
	case IABx:
		a, bx := me.ABx()
		var outText = fmt.Sprintf("%d", a)
		if me.BMode() == InstructionArgK {
			outText += fmt.Sprintf(" %d", -1-bx)
		} else if me.BMode() == InstructionArgU {
			outText += fmt.Sprintf(" %d", bx)
		}
		return fmt.Sprintf("%s [ABx]", outText)
	case IAsBx:
		a, sbx := me.AsBx()
		return fmt.Sprintf("%d %d [AsBx]", a, sbx)
	case IAx:
		ax := me.Ax()
		return fmt.Sprintf("%d [Ax]", -1-ax)
	default:
		return ""
	}
}

// OpMode 获取指令模式
func (me Instruction) OpMode() EInstructionMode {
	return InstructionInfos[me.Opcode()].opMode
}

// BMode 获取操作数B的使用模式
func (me Instruction) BMode() EInstructionArgType {
	return InstructionInfos[me.Opcode()].argBMode
}

// CMode 获取操作数C的使用模式
func (me Instruction) CMode() EInstructionArgType {
	return InstructionInfos[me.Opcode()].argCMode
}

// EInstruction 指令操作码
type EInstruction byte

const (
	OpMove EInstruction = iota
	OpLoadK
	OpLoadKx
	OpLoadBool
	OpLoadNil
	OpGetUpval
	OpGetTabup
	OpGetTable
	OpSetTabup
	OpSetUpval
	OpSetTable
	OpNewTable
	OpSelf
	OpAdd
	OpSub
	OpMul
	OpMod
	OpPow
	OpDiv
	OpIDiv
	OpBand
	OpBor
	OpBxor
	OpShl
	OpShr
	OpUnm
	OpBNot
	OpNot
	OpLen
	OpConcat
	OpJmp
	OpEq
	OpLt
	OpLe
	OpTest
	OpTestSet
	OpCall
	OpTailCall
	OpReturn
	OpForLoop
	OpForPrep
	OpTForCall
	OpTForLoop
	OpSetList
	OpClosure
	OpVararg
	OpExtraarg
)
