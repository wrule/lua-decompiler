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

func readError(err error) {
	panic("文件读取错误")
}

// ReadByte 读取一个字节
func (me *Reader) ReadByte() byte {
	var bytes []byte = make([]byte, 1)
	_, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return bytes[0]
}

// ReadBytes 读取字节数组
func (me *Reader) ReadBytes(size uint64) []byte {
	var bytes []byte = make([]byte, size)
	count, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return bytes[:count]
}

// ReadUint32 读取一个uint32
func (me *Reader) ReadUint32() uint32 {
	var bytes []byte = make([]byte, 4)
	_, err := me.reader.Read(bytes)
	if err != nil {
		readError(err)
	}
	return binary.LittleEndian.Uint32(bytes)
}

// ReadUint64 读取一个uint64
func (me *Reader) ReadUint64() uint64 {
	var bytes []byte = make([]byte, 8)
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
	bytes := me.ReadBytes(size)
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
func (me *Reader) ReadConstant() interface{} {
	ctype := me.ReadByte()
	switch ctype {
	case TagNil:
		return nil
	case TagBoolean:
		return me.ReadByte() != 0x00
	case TagNumber:
		return me.ReadLuaNumber()
	case TagInteger:
		return me.ReadLuaInteger()
	case TagShortStr:
		return me.ReadString()
	case TagLongStr:
		return me.ReadString()
	default:
		panic("常量类型解析错误")
	}
}

// ReadConstants 读取常量列表
func (me *Reader) ReadConstants() []interface{} {
	size := me.ReadUint32()
	constants := make([]interface{}, size)
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

// NewReader 构造函数
func NewReader(reader io.Reader) *Reader {
	return &Reader{reader}
}
