package chunk

// Prototype 函数原型
type Prototype struct {
	source          string
	lineDefined     uint32
	lastLineDefined uint32
	numParams       byte
	isVararg        byte
	maxStackSize    byte
	codes           []uint32
	constants       []interface{}
	upvalues        []Upvalue
	protos          []*Prototype
	lineInfos       []uint32
	locVars         []LocVar
	upvalueNames    []string
}

// Source 获取源文件名
func (me *Prototype) Source() string {
	return me.source
}

// LineDefined 获取开始行号
func (me *Prototype) LineDefined() uint32 {
	return me.lineDefined
}

// LastLineDefined 获取结束行号
func (me *Prototype) LastLineDefined() uint32 {
	return me.lastLineDefined
}

// NumParams 获取固定参数个数
func (me *Prototype) NumParams() byte {
	return me.numParams
}

// IsVararg 是否是Vararg函数
func (me *Prototype) IsVararg() byte {
	return me.isVararg
}

// MaxStackSize 获取需要寄存器数量
func (me *Prototype) MaxStackSize() byte {
	return me.maxStackSize
}

// Codes 获取指令表
func (me *Prototype) Codes() []uint32 {
	return me.codes
}

// Constants 获取常量表
func (me *Prototype) Constants() []interface{} {
	return me.constants
}

// Upvalues 获取Upvalue表
func (me *Prototype) Upvalues() []Upvalue {
	return me.upvalues
}

// Prototypes 获取子函数原型表
func (me *Prototype) Prototypes() []*Prototype {
	return me.protos
}

// LineInfos 获取行号表
func (me *Prototype) LineInfos() []uint32 {
	return me.lineInfos
}

// LocVars 获取局部变量表
func (me *Prototype) LocVars() []LocVar {
	return me.locVars
}

// UpvalueNames 获取Upvalue名列表
func (me *Prototype) UpvalueNames() []string {
	return me.upvalueNames
}
