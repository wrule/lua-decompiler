package api

// LuaState s
type LuaState struct {
	stack *LuaStack
}

func (me *LuaState) GetTop() int {
	return me.stack.top
}

func (me *LuaState) AbsIndex(index int) int {
	return me.stack.AbsIndex(index)
}

func (me *LuaState) CheckStack(n int) bool {
	me.stack.Check(n)
	return true
}

func (me *LuaState) Pop(n int) {
	for i := 0; i < n; i++ {
		me.stack.Pop()
	}
}

func (me *LuaState) Copy(fromIndex int, toIndex int) {
	value := me.stack.Get(fromIndex)
	me.stack.Set(toIndex, value)
}

func (me *LuaState) PushValue(index int) {
	value := me.stack.Get(index)
	me.stack.Push(value)
}

func (me *LuaState) Replace(index int) {
	value := me.stack.Pop()
	me.stack.Set(index, value)
}

// NewLuaState 构造函数创建一个LuaState
func NewLuaState() *LuaState {
	return &LuaState{
		stack: NewLuaStack(20),
	}
}
