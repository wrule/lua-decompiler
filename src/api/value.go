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
