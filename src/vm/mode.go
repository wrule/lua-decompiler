package vm

// EInstructionMode 指令编码模式枚举
type EInstructionMode int

const (
	// IABC 6 8 9 9
	IABC EInstructionMode = iota
	// IABx 6 8 18
	IABx
	// IAsBx 6 8 18
	IAsBx
	// IAx 6 26
	IAx
)
