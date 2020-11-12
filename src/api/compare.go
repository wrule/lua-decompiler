package api

// Compare 比较Lua值
func (me *LuaState) Compare(index1, index2 int, op ECompareOpType) bool {
	a := me.stack.Get(index1)
	b := me.stack.Get(index2)
	switch op {
	case LuaOpEQ:
		return compareEQ(a, b)
	case LuaOpLT:
		return compareLT(a, b)
	case LuaOpLE:
		return compareLE(a, b)
	default:
		panic("错误的比较运算符")
	}
	return false
}

func compareEQ(a, b LuaValue) bool {
	return true
}

func compareLT(a, b LuaValue) bool {
	return true
}

func compareLE(a, b LuaValue) bool {
	return true
}
