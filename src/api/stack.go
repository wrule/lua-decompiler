package api

// LuaStack Lua栈
type LuaStack struct {
	slots []LuaValue
	top   int
}

// Top 获取栈顶位置索引（内部索引）
func (me *LuaStack) Top() int {
	return me.top
}

// Size 获取栈容量
func (me *LuaStack) Size() int {
	return len(me.slots)
}

// LuaNilValue 生成一个LuaTypeNil值
func (me *LuaStack) LuaNilValue() LuaValue {
	return LuaValue{
		vtype: LuaTypeNil,
		value: nil,
	}
}

// Check 检查栈中是否有满足n个LuaValue的足够空间，如果没有则扩容
func (me *LuaStack) Check(n int) {
	var diff = n - (me.Size() - me.top)
	for i := 0; i < diff; i++ {
		me.slots = append(me.slots, me.LuaNilValue())
	}
}

// Push 向栈中压入一个值
func (me *LuaStack) Push(value LuaValue) {
	if me.top >= me.Size() {
		panic("栈空间不足，无法Push")
	}
	me.slots[me.top] = value
	me.top++
}

// Pop 从栈中弹出一个值
func (me *LuaStack) Pop() LuaValue {
	if me.top < 1 {
		panic("栈已经为空，无法Pop")
	}
	me.top--
	value := me.slots[me.top]
	me.slots[me.top] = me.LuaNilValue()
	return value
}

// AbsIndex 索引转换成为绝对索引
func (me *LuaStack) AbsIndex(index int) int {
	if index >= 0 {
		return index
	}
	return index + me.top + 1
}

// IsValid 判断一个索引是否有效
func (me *LuaStack) IsValid(index int) bool {
	absIndex := me.AbsIndex(index)
	return absIndex > 0 && absIndex <= me.top
}

// Get 按索引取Lua值
func (me *LuaStack) Get(index int) LuaValue {
	if me.IsValid(index) {
		absIndex := me.AbsIndex(index)
		return me.slots[absIndex-1]
	}
	return me.LuaNilValue()
}

// Set 按索引设置Lua值
func (me *LuaStack) Set(index int, value LuaValue) {
	if me.IsValid(index) {
		absIndex := me.AbsIndex(index)
		me.slots[absIndex-1] = value
		return
	}
	panic("错误的索引，无法对栈进行Set")
}

// Reverse 区间倒序
func (me *LuaStack) Reverse(fromIndex, toIndex int) {
	slots := me.slots
	for fromIndex < toIndex {
		slots[fromIndex], slots[toIndex] = slots[toIndex], slots[fromIndex]
		fromIndex++
		toIndex--
	}
}

// NewLuaStack 构造函数，新建一个Lua栈
func NewLuaStack(size int) *LuaStack {
	return &LuaStack{
		slots: make([]LuaValue, size),
		top:   0,
	}
}
