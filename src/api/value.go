package api

// LuaValue 老泪纵横
type LuaValue struct {
	vtype ELuaValueType
	value interface{}
}

func (me *LuaValue) Type() ELuaValueType {
	return me.vtype
}

func (me *LuaValue) Value() interface{} {
	return me.value
}

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
