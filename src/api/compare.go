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
}

func compareEQ(a, b LuaValue) bool {
	switch a.Type() {
	case LuaTypeNil:
		return b.Type() == LuaTypeNil
	case LuaTypeBoolean:
		x := a.Value().(bool)
		y, ok := b.Value().(bool)
		return ok && x == y
	case LuaTypeString:
		x := a.Value().(string)
		y, ok := b.Value().(string)
		return ok && x == y
	case LuaTypeInteger:
		switch b.Type() {
		case LuaTypeInteger:
			return a.Value() == b.Value()
		case LuaTypeNumber:
			return float64(a.Value().(int64)) == b.Value()
		default:
			return false
		}
	case LuaTypeNumber:
		switch b.Type() {
		case LuaTypeInteger:
			return a.Value() == float64(b.Value().(int64))
		case LuaTypeNumber:
			return a.Value() == b.Value()
		default:
			return false
		}
	default:
		return a.Value() == b.Value()
	}
}

func compareLT(a, b LuaValue) bool {
	return true
}

func compareLE(a, b LuaValue) bool {
	return true
}
