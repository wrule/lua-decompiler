package api

// Len 获取栈中值的长度
func (me *LuaState) Len(index int) int {
	value := me.stack.Get(index)
	if str, ok := value.Value().(string); ok {
		me.PushInteger(int64(len(str)))
	}
	panic("长度计算错误")
}
