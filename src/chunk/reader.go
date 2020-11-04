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

// ReadByte 读取一个字节
func (me *Reader) ReadByte() (byte, error) {
	var bytes []byte = make([]byte, 1)
	_, err := me.reader.Read(bytes)
	return bytes[0], err
}

// ReadUint32 读取一个uint32
func (me *Reader) ReadUint32() (uint32, error) {
	var bytes []byte = make([]byte, 4)
	_, err := me.reader.Read(bytes)
	return binary.LittleEndian.Uint32(bytes), err
}

// ReadUint64 读取一个uint64
func (me *Reader) ReadUint64() (uint64, error) {
	var bytes []byte = make([]byte, 8)
	_, err := me.reader.Read(bytes)
	return binary.LittleEndian.Uint64(bytes), err
}

// ReadLuaInteger 读取一个Lua整数
func (me *Reader) ReadLuaInteger() (int64, error) {
	num, err := me.ReadUint64()
	return int64(num), err
}

// ReadLuaNumber 读取一个Lua浮点数
func (me *Reader) ReadLuaNumber() (float64, error) {
	num, err := me.ReadUint64()
	return math.Float64frombits(num), err
}

// ReadString 读取一个字符串
func (me *Reader) ReadString() (string, error) {
	sizeByte, err := me.ReadByte()
	var size uint64 = uint64(sizeByte)
	if size == 0 {
		return "", err
	}
	if size == 0xFF {
		// 这个err被放弃了
		size, err = me.ReadUint64()
	}
	bytes, err := me.ReadBytes(size)
	return string(bytes), err
}

// ReadBytes 读取字节数组
func (me *Reader) ReadBytes(size uint64) ([]byte, error) {
	var bytes []byte = make([]byte, size)
	count, err := me.reader.Read(bytes)
	return bytes[:count], err
}

// ReadCodes 读取Lua虚拟机指令
func (me *Reader) ReadCodes() []uint32 {
	size, _ := me.ReadUint32()
	codes := make([]uint32, size)
	for index := range codes {
		codes[index], _ = me.ReadUint32()
	}
	return codes
}

// ReadConstant 读取常量
func (me *Reader) ReadConstant() interface{} {
	ctype, _ := me.ReadByte()
	switch ctype {
	case TagNil:
		return nil
	case TagBoolean:
		res, _ := me.ReadByte()
		return res != 0x00
	case TagNumber:
		res, _ := me.ReadLuaNumber()
		return res
	case TagInteger:
		res, _ := me.ReadLuaInteger()
		return res
	case TagShortStr:
		res, _ := me.ReadString()
		return res
	case TagLongStr:
		res, _ := me.ReadString()
		return res
	default:
		panic("常量类型解析错误")
	}
}

// ReadConstants 读取常量列表
func (me *Reader) ReadConstants() []interface{} {
	size, _ := me.ReadUint32()
	constants := make([]interface{}, size)
	for index := range constants {
		constants[index] = me.ReadConstant()
	}
	return constants
}

// ReadUpvalues 读取Upvalues列表
func (me *Reader) ReadUpvalues() []Upvalue {
	size, _ := me.ReadUint32()
	upvalues := make([]Upvalue, size)
	for index := range upvalues {
		instack, _ := me.ReadByte()
		idx, _ := me.ReadByte()
		upvalues[index] = Upvalue{instack, idx}
	}
	return upvalues
}

// NewReader 构造函数
func NewReader(reader io.Reader) *Reader {
	return &Reader{reader}
}
