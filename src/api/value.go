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

// ToBoolean 尝试转换为布尔类型
func (me *LuaValue) ToBoolean() *LuaValue {
	switch me.Type() {
	// 确定这个不是false？
	case LuaTypeNone:
		return &LuaValue{
			vtype: LuaTypeBoolean,
			value: false,
		}
	case LuaTypeNil:
		return &LuaValue{
			vtype: LuaTypeBoolean,
			value: false,
		}
	case LuaTypeBoolean:
		return me
	default:
		return &LuaValue{
			vtype: LuaTypeBoolean,
			value: true,
		}
	}
}

// ToNumberX s
func (me *LuaValue) ToNumberX() (*LuaValue, bool) {
	switch me.Type() {
	case LuaTypeNumber:
		return me, true
	default:
		return &LuaValue{
			vtype: LuaTypeNumber,
			value: float64(0),
		}, false
	}
}
