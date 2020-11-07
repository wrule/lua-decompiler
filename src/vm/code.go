package vm

type ECodeMode int

const (
	// IABC 6 8 9 9
	IABC = iota
	// IABx 6 8 18
	IABx
	// IAsBx 6 8 18
	IAsBx
	// IAx 6 26
	IAx
)

const (
	OpMove = iota
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
