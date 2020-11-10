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

// Name Lua值的类型的名称
func (me ELuaValueType) Name() string {
	switch me {
	case LuaTypeNone:
		return "no value"
	case LuaTypeNil:
		return "nil"
	case LuaTypeBoolean:
		return "boolean"
	case LuaTypeNumber:
		return "number"
	case LuaTypeString:
		return "string"
	case LuaTypeTable:
		return "table"
	case LuaTypeFunction:
		return "function"
	case LuaTypeUserData:
		return "userdata"
	case LuaTypeLightUserData:
		return "light_userdata"
	case LuaTypeThread:
		return "thread"
	default:
		return ""
	}
}

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
