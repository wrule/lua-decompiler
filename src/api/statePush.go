package api

// PushNil 向栈中压入一个nil值
func (me *LuaState) PushNil() {
	me.stack.Push(me.stack.LuaNilValue())
}

// PushBoolean 向栈中压入一个布尔值
func (me *LuaState) PushBoolean(value bool) {
	me.stack.Push(LuaValue{
		vtype: LuaTypeBoolean,
		value: value,
	})
}

// PushInteger 向栈中压入一个整数值
func (me *LuaState) PushInteger(value int64) {
	me.stack.Push(LuaValue{
		vtype: LuaTypeInteger,
		value: value,
	})
}

// PushNumber 向栈中压入一个浮点数值
func (me *LuaState) PushNumber(value float64) {
	me.stack.Push(LuaValue{
		vtype: LuaTypeNumber,
		value: value,
	})
}

// PushString 向栈中压入一个字符串值
func (me *LuaState) PushString(value string) {
	me.stack.Push(LuaValue{
		vtype: LuaTypeString,
		value: value,
	})
}
