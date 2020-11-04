package chunk

// Header Chunk的Header部分
type Header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizetSize       byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

const (
	// LuaSignature 魔数
	LuaSignature = "\x1bLua"
	// LuacVersion Luac版本
	LuacVersion = 0x53
	// LuacFormat Luac格式数字
	LuacFormat = 0x00
	// LuacData Luac预留验证数据
	LuacData = "\x19\x93\r\n\x1a\n"
	// CIntSize Lua虚拟机cint宽度
	CIntSize = 4
	// SizeTSize Lua虚拟机size_t宽度
	SizeTSize = 8
	// InstructionSize Lua虚拟机指令宽度
	InstructionSize = 4
	// LuaIntegerSize Lua虚拟机整数宽度
	LuaIntegerSize = 8
	// LuaNumberSize Lua虚拟机浮点数宽度
	LuaNumberSize = 8
	// LuacInt 整数验证数字
	LuacInt = 0x5678
	// LuacNum 浮点数验证数字
	LuacNum = 370.5
)

// CheckLoad 从Reader中检查并加载Header
func (me *Header) CheckLoad(reader *Reader) {
	{
		res := reader.ReadBytes(4)
		if string(res) != LuaSignature {
			panic("文件没有有效的Lua签名")
		}
		copy(me.signature[:], res)
	}
	{
		res := reader.ReadByte()
		if res != LuacVersion {
			panic("Luac版本校验失败")
		}
		me.version = res
	}
	{
		res := reader.ReadByte()
		if res != LuacFormat {
			panic("Luac格式数字校验失败")
		}
		me.format = res
	}
	{
		res := reader.ReadBytes(6)
		if string(res) != LuacData {
			panic("文件没有有效的Luac验证数据")
		}
		copy(me.luacData[:], res)
	}
	{
		res := reader.ReadByte()
		if res != CIntSize {
			panic("cint宽度校验失败")
		}
		me.cintSize = res
	}
	{
		res := reader.ReadByte()
		if res != SizeTSize {
			panic("sizetSize宽度校验失败")
		}
		me.sizetSize = res
	}
	{
		res := reader.ReadByte()
		if res != InstructionSize {
			panic("指令宽度校验失败")
		}
		me.instructionSize = res
	}
	{
		res := reader.ReadByte()
		if res != LuaIntegerSize {
			panic("LuaInteger宽度校验失败")
		}
		me.luaIntegerSize = res
	}
	{
		res := reader.ReadByte()
		if res != LuaNumberSize {
			panic("LuaNumber宽度校验失败")
		}
		me.luaNumberSize = res
	}
	{
		res := reader.ReadLuaInteger()
		if res != LuacInt {
			panic("Luac整数验证数字校验失败")
		}
		me.luacInt = res
	}
	{
		res := reader.ReadLuaNumber()
		if res != LuacNum {
			panic("Luac浮点数验证数字校验失败")
		}
		me.luacNum = res
	}
}
