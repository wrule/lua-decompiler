package api

import (
	"fmt"
	"strconv"
)

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

// ToBoolean 尝试取得布尔类型值
func (me *LuaValue) ToBoolean() bool {
	switch me.Type() {
	// 确定这个不是false？
	case LuaTypeNone:
		return false
	case LuaTypeNil:
		return false
	case LuaTypeBoolean:
		return me.Value().(bool)
	default:
		return true
	}
}

// ToNumberX 尝试取得float64的值
func (me *LuaValue) ToNumberX() (float64, bool) {
	switch me.Type() {
	case LuaTypeInteger:
		num := float64(me.Value().(int64))
		return num, true
	case LuaTypeNumber:
		return me.Value().(float64), true
	case LuaTypeString:
		num, err := strconv.ParseFloat(me.Value().(string), 64)
		return num, err == nil
	default:
		return float64(0), false
	}
}

// ToIntegerX 尝试取得int64的值
func (me *LuaValue) ToIntegerX() (int64, bool) {
	value := me.Value()
	num, ok := value.(int64)
	return num, ok
}

// ToStringX s
func (me *LuaValue) ToStringX() (string, bool) {
	switch me.Type() {
	case LuaTypeString:
		return me.Value().(string), true
	case LuaTypeInteger, LuaTypeNumber:
		str := fmt.Sprintf("%v", me.Value())
		me.vtype = LuaTypeString
		me.value = str
		return str, true
	default:
		return "", false
	}
}
