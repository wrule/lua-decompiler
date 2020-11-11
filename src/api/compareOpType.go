package api

// ECompareOpType Lua比较运算符
type ECompareOpType int

const (
	LuaOpEQ ECompareOpType = iota // ==
	LuaOpLT                       // <
	LuaOpLE                       // <=
)
