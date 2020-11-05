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

// Source 代码文件位置
func (me *Prototype) Source() string {
	return me.source
}

func (me *Prototype) LineDefined() uint32 {
	return me.lineDefined
}

func (me *Prototype) LastLineDefined() uint32 {
	return me.lastLineDefined
}

func (me *Prototype) NumParams() byte {
	return me.numParams
}

func (me *Prototype) IsVararg() byte {
	return me.isVararg
}

func (me *Prototype) MaxStackSize() byte {
	return me.maxStackSize
}

func (me *Prototype) Codes() []uint32 {
	return me.codes
}

func (me *Prototype) Constants() []interface{} {
	return me.constants
}

func (me *Prototype) Upvalues() []Upvalue {
	return me.upvalues
}

func (me *Prototype) Prototypes() []*Prototype {
	return me.protos
}

func (me *Prototype) LineInfos() []uint32 {
	return me.lineInfos
}

func (me *Prototype) LocVars() []LocVar {
	return me.locVars
}

func (me *Prototype) UpvalueNames() []string {
	return me.upvalueNames
}
