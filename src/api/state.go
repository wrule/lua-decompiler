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

// func (me *LuaState) Pop(n int) {
// 	for i := 0; i < n; i++ {
// 		me.stack.Pop()
// 	}
// }

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

// Insert s
func (me *LuaState) Insert(index int) {
	me.Rotate(index, 1)
}

// Remove s
func (me *LuaState) Remove(index int) {
	me.Rotate(index, -1)
	me.Pop(1)
}

// Rotate 区间旋转
func (me *LuaState) Rotate(index, n int) {
	// [1, 2, 3, 4, 5]
	// len = 5 index = 2 n = 2
	// [1，3，2，5，4]
	// [1，4，5，2，3]
	t := me.stack.top - 1
	p := me.stack.AbsIndex(index) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	me.stack.Reverse(p, m)
	me.stack.Reverse(m+1, t)
	me.stack.Reverse(p, t)
}

// SetTop s
func (me *LuaState) SetTop(index int) {

}

func (me *LuaState) Pop(n int) {
	me.SetTop(-n - 1)
}

// NewLuaState 构造函数创建一个LuaState
func NewLuaState() *LuaState {
	return &LuaState{
		stack: NewLuaStack(20),
	}
}
