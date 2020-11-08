package vm

// EInstructionArgType 指令参数类型
type EInstructionArgType byte

const (
	// InstructionArgN 不表示任何信息
	InstructionArgN EInstructionArgType = iota
	// InstructionArgR IABC模式表示寄存器索引，IAsBx模式表示跳转偏移
	InstructionArgR
	// InstructionArgK 常量表索引或寄存器索引
	InstructionArgK
	// InstructionArgU 其他情况（布尔值，整数值，Upvalue索引，子函数索引等）
	InstructionArgU
)
