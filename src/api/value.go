package api

// LuaValue Lua值
type LuaValue struct {
	vtype ELuaValueType
	value interface{}
}

// Type Lua值的类型
func (me *LuaValue) Type() ELuaValueType {
	return me.vtype
}

// Value Lua值的真实值
func (me *LuaValue) Value() interface{} {
	return me.value
}

// ELuaValueType Lua值的类型的枚举
type ELuaValueType int

const (
	// LuaTypeNone -1
	LuaTypeNone ELuaValueType = iota - 1
	LuaTypeNil
	LuaTypeBoolean
	LuaTypeNumber
	LuaTypeString
	LuaTypeTable
	LuaTypeFunction
	LuaTypeUserData
	LuaTypeLightUserData
	LuaTypeThread
)
