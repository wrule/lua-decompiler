package vm

// ECodeMode 指令模式
type ECodeMode int

const (
	// IABC 6 8 9 9
	IABC ECodeMode = iota
	// IABx 6 8 18
	IABx
	// IAsBx 6 8 18
	IAsBx
	// IAx 6 26
	IAx
)

// CodeInfo 指令信息
type CodeInfo struct {
	testFlag byte
	setAFlag byte
	argBMode ECodeArgType
	argCMode ECodeArgType
	opMode   ECodeMode
	name     string
}

// OpCodes 指令信息列表
var OpCodes = []CodeInfo{
	/*   T  A  B       C       mode        name    */
	CodeInfo{0, 1, CodeArgR, CodeArgN, IABC /* */, "MOVE    "}, // R(A) := R(B)
	CodeInfo{0, 1, CodeArgK, CodeArgN, IABx /* */, "LOADK   "}, // R(A) := Kst(Bx)
	CodeInfo{0, 1, CodeArgN, CodeArgN, IABx /* */, "LOADKX  "}, // R(A) := Kst(extra arg)
	CodeInfo{0, 1, CodeArgU, CodeArgU, IABC /* */, "LOADBOOL"}, // R(A) := (bool)B; if (C) pc++
	CodeInfo{0, 1, CodeArgU, CodeArgN, IABC /* */, "LOADNIL "}, // R(A), R(A+1), ..., R(A+B) := nil
	CodeInfo{0, 1, CodeArgU, CodeArgN, IABC /* */, "GETUPVAL"}, // R(A) := UpValue[B]
	CodeInfo{0, 1, CodeArgU, CodeArgK, IABC /* */, "GETTABUP"}, // R(A) := UpValue[B][RK(C)]
	CodeInfo{0, 1, CodeArgR, CodeArgK, IABC /* */, "GETTABLE"}, // R(A) := R(B)[RK(C)]
	CodeInfo{0, 0, CodeArgK, CodeArgK, IABC /* */, "SETTABUP"}, // UpValue[A][RK(B)] := RK(C)
	CodeInfo{0, 0, CodeArgU, CodeArgN, IABC /* */, "SETUPVAL"}, // UpValue[B] := R(A)
	CodeInfo{0, 0, CodeArgK, CodeArgK, IABC /* */, "SETTABLE"}, // R(A)[RK(B)] := RK(C)
	CodeInfo{0, 1, CodeArgU, CodeArgU, IABC /* */, "NEWTABLE"}, // R(A) := {} (size = B,C)
	CodeInfo{0, 1, CodeArgR, CodeArgK, IABC /* */, "SELF    "}, // R(A+1) := R(B); R(A) := R(B)[RK(C)]
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "ADD     "}, // R(A) := RK(B) + RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "SUB     "}, // R(A) := RK(B) - RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "MUL     "}, // R(A) := RK(B) * RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "MOD     "}, // R(A) := RK(B) % RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "POW     "}, // R(A) := RK(B) ^ RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "DIV     "}, // R(A) := RK(B) / RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "IDIV    "}, // R(A) := RK(B) // RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "BAND    "}, // R(A) := RK(B) & RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "BOR     "}, // R(A) := RK(B) | RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "BXOR    "}, // R(A) := RK(B) ~ RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "SHL     "}, // R(A) := RK(B) << RK(C)
	CodeInfo{0, 1, CodeArgK, CodeArgK, IABC /* */, "SHR     "}, // R(A) := RK(B) >> RK(C)
	CodeInfo{0, 1, CodeArgR, CodeArgN, IABC /* */, "UNM     "}, // R(A) := -R(B)
	CodeInfo{0, 1, CodeArgR, CodeArgN, IABC /* */, "BNOT    "}, // R(A) := ~R(B)
	CodeInfo{0, 1, CodeArgR, CodeArgN, IABC /* */, "NOT     "}, // R(A) := not R(B)
	CodeInfo{0, 1, CodeArgR, CodeArgN, IABC /* */, "LEN     "}, // R(A) := length of R(B)
	CodeInfo{0, 1, CodeArgR, CodeArgR, IABC /* */, "CONCAT  "}, // R(A) := R(B).. ... ..R(C)
	CodeInfo{0, 0, CodeArgR, CodeArgN, IAsBx /**/, "JMP     "}, // pc+=sBx; if (A) close all upvalues >= R(A - 1)
	CodeInfo{1, 0, CodeArgK, CodeArgK, IABC /* */, "EQ      "}, // if ((RK(B) == RK(C)) ~= A) then pc++
	CodeInfo{1, 0, CodeArgK, CodeArgK, IABC /* */, "LT      "}, // if ((RK(B) <  RK(C)) ~= A) then pc++
	CodeInfo{1, 0, CodeArgK, CodeArgK, IABC /* */, "LE      "}, // if ((RK(B) <= RK(C)) ~= A) then pc++
	CodeInfo{1, 0, CodeArgN, CodeArgU, IABC /* */, "TEST    "}, // if not (R(A) <=> C) then pc++
	CodeInfo{1, 1, CodeArgR, CodeArgU, IABC /* */, "TESTSET "}, // if (R(B) <=> C) then R(A) := R(B) else pc++
	CodeInfo{0, 1, CodeArgU, CodeArgU, IABC /* */, "CALL    "}, // R(A), ... ,R(A+C-2) := R(A)(R(A+1), ... ,R(A+B-1))
	CodeInfo{0, 1, CodeArgU, CodeArgU, IABC /* */, "TAILCALL"}, // return R(A)(R(A+1), ... ,R(A+B-1))
	CodeInfo{0, 0, CodeArgU, CodeArgN, IABC /* */, "RETURN  "}, // return R(A), ... ,R(A+B-2)
	CodeInfo{0, 1, CodeArgR, CodeArgN, IAsBx /**/, "FORLOOP "}, // R(A)+=R(A+2); if R(A) <?= R(A+1) then { pc+=sBx; R(A+3)=R(A) }
	CodeInfo{0, 1, CodeArgR, CodeArgN, IAsBx /**/, "FORPREP "}, // R(A)-=R(A+2); pc+=sBx
	CodeInfo{0, 0, CodeArgN, CodeArgU, IABC /* */, "TFORCALL"}, // R(A+3), ... ,R(A+2+C) := R(A)(R(A+1), R(A+2));
	CodeInfo{0, 1, CodeArgR, CodeArgN, IAsBx /**/, "TFORLOOP"}, // if R(A+1) ~= nil then { R(A)=R(A+1); pc += sBx }
	CodeInfo{0, 0, CodeArgU, CodeArgU, IABC /* */, "SETLIST "}, // R(A)[(C-1)*FPF+i] := R(A+i), 1 <= i <= B
	CodeInfo{0, 1, CodeArgU, CodeArgN, IABx /* */, "CLOSURE "}, // R(A) := closure(KPROTO[Bx])
	CodeInfo{0, 1, CodeArgU, CodeArgN, IABC /* */, "VARARG  "}, // R(A), R(A+1), ..., R(A+B-2) = vararg
	CodeInfo{0, 0, CodeArgU, CodeArgU, IAx /*  */, "EXTRAARG"}, // extra (larger) argument for previous Code
}

// ECodeArgType 指令参数类型
type ECodeArgType byte

const (
	// CodeArgN 不表示任何信息
	CodeArgN ECodeArgType = iota
	// CodeArgR IABC模式表示寄存器索引，IAsBx模式表示跳转偏移
	CodeArgR
	// CodeArgK 常量表索引或寄存器索引
	CodeArgK
	// CodeArgU 其他情况（布尔值，整数值，Upvalue索引，子函数索引等）
	CodeArgU
)

// ECode 指令操作码
type ECode byte

const (
	OpMove ECode = iota
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
