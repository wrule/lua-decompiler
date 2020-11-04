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

// NewReader 构造函数
func NewReader(reader io.Reader) *Reader {
	return &Reader{reader}
}
