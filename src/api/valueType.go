package api

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
