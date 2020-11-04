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
	protos          []*Prototype
	lineInfo        []uint32
	upvalueNames    []string
}
