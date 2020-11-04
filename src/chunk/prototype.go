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
	me.source, _ = reader.ReadString()
	me.lineDefined, _ = reader.ReadUint32()
	me.lastLineDefined, _ = reader.ReadUint32()
	me.numParams, _ = reader.ReadByte()
	me.isVararg, _ = reader.ReadByte()
	me.maxStackSize, _ = reader.ReadByte()
	me.code = reader.ReadCodes()
	me.constants = reader.ReadConstants()
	me.upvalues = reader.ReadUpvalues()
}
