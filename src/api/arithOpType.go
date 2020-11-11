package api

// EArithOpType 算术或按位运算符类型枚举
type EArithOpType int

const (
	LuaOpADD  EArithOpType = iota // +
	LuaOpSUB                      // -
	LuaOpMUL                      // *
	LuaOpMOD                      // %
	LuaOpPOW                      // ^
	LuaOpDIV                      // /
	LuaOpIDIV                     // //
	LuaOpBAND                     // &
	LuaOpBOR                      // |
	LuaOpBXOR                     // ~
	LuaOpSHL                      // <<
	LuaOpSHR                      // >>
	LuaOpUNM                      // -
	LuaOpBNOT                     // ~
)
