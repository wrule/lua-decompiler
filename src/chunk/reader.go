package chunk

import (
	"encoding/binary"
	"io"
	"math"
)

// Reader 读取器
type Reader struct {
	reader io.Reader
}

// readError 用于错误处理
func readError(err error) {
	panic("读取chunk出现错误")
}

// ReadByte 读取一个字节
func (me *Reader) ReadByte() byte {
	var bytes = make([]byte, 1)
	_, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return bytes[0]
}

// ReadBytes 读取字节数组
func (me *Reader) ReadBytes(size uint64) []byte {
	var bytes = make([]byte, size)
	count, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return bytes[:count]
}

// ReadUint32 读取一个uint32
func (me *Reader) ReadUint32() uint32 {
	var bytes = make([]byte, 4)
	_, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return binary.LittleEndian.Uint32(bytes)
}

// ReadUint64 读取一个uint64
func (me *Reader) ReadUint64() uint64 {
	var bytes = make([]byte, 8)
	_, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return binary.LittleEndian.Uint64(bytes)
}

// ReadLuaInteger 读取一个Lua整数
func (me *Reader) ReadLuaInteger() int64 {
	num := me.ReadUint64()
	return int64(num)
}

// ReadLuaNumber 读取一个Lua浮点数
func (me *Reader) ReadLuaNumber() float64 {
	num := me.ReadUint64()
	return math.Float64frombits(num)
}

// ReadString 读取一个字符串
func (me *Reader) ReadString() string {
	var size = uint64(me.ReadByte())
	if size == 0 {
		return ""
	}
	if size == 0xFF {
		size = me.ReadUint64()
	}
	bytes := me.ReadBytes(size - 1)
	return string(bytes)
}

// ReadCodes 读取Lua虚拟机指令
func (me *Reader) ReadCodes() []uint32 {
	size := me.ReadUint32()
	codes := make([]uint32, size)
	for index := range codes {
		codes[index] = me.ReadUint32()
	}
	return codes
}

// ReadConstant 读取常量
func (me *Reader) ReadConstant() Constant {
	ctype := EConstantType(me.ReadByte())
	switch ctype {
	case ConstantTypeNil:
		return Constant{ConstantTypeNil, nil}
	case ConstantTypeBoolean:
		return Constant{ConstantTypeBoolean, me.ReadByte() != 0x00}
	case ConstantTypeNumber:
		return Constant{ConstantTypeNumber, me.ReadLuaNumber()}
	case ConstantTypeInteger:
		return Constant{ConstantTypeInteger, me.ReadLuaInteger()}
	case ConstantTypeShortStr:
		return Constant{ConstantTypeShortStr, me.ReadString()}
	case ConstantTypeLongStr:
		return Constant{ConstantTypeLongStr, me.ReadString()}
	default:
		panic("常量类型解析错误")
	}
}

// ReadConstants 读取常量列表
func (me *Reader) ReadConstants() []Constant {
	size := me.ReadUint32()
	constants := make([]Constant, size)
	for index := range constants {
		constants[index] = me.ReadConstant()
	}
	return constants
}

// ReadUpvalues 读取Upvalues列表
func (me *Reader) ReadUpvalues() []Upvalue {
	size := me.ReadUint32()
	upvalues := make([]Upvalue, size)
	for index := range upvalues {
		instack := me.ReadByte()
		idx := me.ReadByte()
		upvalues[index] = Upvalue{instack, idx}
	}
	return upvalues
}

// ReadPrototype 读取函数原型
func (me *Reader) ReadPrototype(parentSource string) *Prototype {
	source := me.ReadString()
	if source == "" {
		source = parentSource
	}
	return &Prototype{
		source:          source,
		lineDefined:     me.ReadUint32(),
		lastLineDefined: me.ReadUint32(),
		numParams:       me.ReadByte(),
		isVararg:        me.ReadByte(),
		maxStackSize:    me.ReadByte(),
		codes:           me.ReadCodes(),
		constants:       me.ReadConstants(),
		upvalues:        me.ReadUpvalues(),
		protos:          me.ReadPrototypes(source),
		lineInfos:       me.ReadLineInfos(),
		locVars:         me.ReadLocVars(),
		upvalueNames:    me.ReadUpvalueNames(),
	}
}

// ReadPrototypes 读取函数原型列表
func (me *Reader) ReadPrototypes(parentSource string) []*Prototype {
	size := me.ReadUint32()
	protos := make([]*Prototype, size)
	for index := range protos {
		protos[index] = me.ReadPrototype(parentSource)
	}
	return protos
}

// ReadLineInfos 读取行号表
func (me *Reader) ReadLineInfos() []uint32 {
	size := me.ReadUint32()
	lineInfos := make([]uint32, size)
	for index := range lineInfos {
		lineInfos[index] = me.ReadUint32()
	}
	return lineInfos
}

// ReadLocVars 读取本地变量列表
func (me *Reader) ReadLocVars() []LocVar {
	size := me.ReadUint32()
	locVars := make([]LocVar, size)
	for index := range locVars {
		locVars[index] = LocVar{
			varName: me.ReadString(),
			startPC: me.ReadUint32(),
			endPC:   me.ReadUint32(),
		}
	}
	return locVars
}

// ReadUpvalueNames 读取Upvalue名称列表
func (me *Reader) ReadUpvalueNames() []string {
	size := me.ReadUint32()
	names := make([]string, size)
	for index := range names {
		names[index] = me.ReadString()
	}
	return names
}

// NewReader 构造函数
func NewReader(reader io.Reader) *Reader {
	return &Reader{reader}
}
