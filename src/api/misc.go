package api

// Len 获取栈中值的长度
func (me *LuaState) Len(index int) int {
	value := me.stack.Get(index)
	if str, ok := value.Value().(string); ok {
		me.PushInteger(int64(len(str)))
	}
	panic("长度计算错误")
}

// Concat 没有按照书上做
func (me *LuaState) Concat(n int) {
	if n == 0 {
		me.PushString("")
	} else if n >= 2 {
		var result = ""
		for i := 0; i < n; i++ {
			if me.IsString(-1) {
				value := me.stack.Pop()
				str, _ := value.ToStringX()
				result += str
			} else {
				panic("Concat连接错误")
			}
		}
		me.PushString(result)
	}
	// 为1的话什么都不做
}
