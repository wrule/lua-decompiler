package api

// LuaTable 表实现
type LuaTable struct {
	lArray []LuaValue
	lMap   map[string]LuaValue
}

// Get 从表中获取值
func (me *LuaTable) Get(key LuaValue) LuaValue {

}

// Put 向表里写入键值对
func (me *LuaTable) Put(key, value LuaValue) {

}

// NewLuaTable LuaTable的构造函数
func NewLuaTable(nArr, nRec int) *LuaTable {
	result := &LuaTable{}
	if nArr > 0 {
		result.lArray = make([]LuaValue, 0, nArr)
	}
	if nRec > 0 {
		result.lMap = make(map[string]LuaValue, nRec)
	}
	return result
}
