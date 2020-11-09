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
	// [1，4，5，2，3]
	// 4
	t := me.stack.top - 1 /* end of stack segment being rotated */
	// 1
	p := me.stack.AbsIndex(index) - 1 /* start of segment */
	var m int                         /* end of prefix */
	if n >= 0 {
		// 2
		m = t - n
	} else {
		m = p - n - 1
	}
	// 1, 2
	// [1，3，2，4，5]
	me.stack.Reverse(p, m) /* reverse the prefix with length 'n' */
	// 3, 4
	// [1，3，2，5，4]
	me.stack.Reverse(m+1, t) /* reverse the suffix */
	// 1, 4
	// [1，4，5，2，3]
	me.stack.Reverse(p, t) /* reverse the entire segment */
}

// NewLuaState 构造函数创建一个LuaState
func NewLuaState() *LuaState {
	return &LuaState{
		stack: NewLuaStack(20),
	}
}
