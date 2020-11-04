package chunk

// Prototype 函数原型
type Prototype struct {
	source          string
	lineDefined     uint32
	lastLineDefined uint32
	numParams       byte
	isVararg        byte
	maxStackSize    byte
	code            []uint32
	constants       []interface{}
	upvalues        []Upvalue
	protos          []*Prototype
	lineInfo        []uint32
	locVars         []LocVar
	upvalueNames    []string
}

// CheckLoad 加载函数原型
func (me *Prototype) CheckLoad(reader Reader, parentSource string) {
	me.source = reader.ReadString()
	me.lineDefined = reader.ReadUint32()
	me.lastLineDefined = reader.ReadUint32()
	me.numParams = reader.ReadByte()
	me.isVararg = reader.ReadByte()
	me.maxStackSize = reader.ReadByte()
	me.code = reader.ReadCodes()
	me.constants = reader.ReadConstants()
	me.upvalues = reader.ReadUpvalues()
}
