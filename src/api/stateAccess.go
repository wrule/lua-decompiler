package api

// TypeName 获取类型的描述
func (me *LuaState) TypeName(tp ELuaValueType) string {
	return tp.Name()
}

// Type 获取指定索引处的值的类型
func (me *LuaState) Type(index int) ELuaValueType {
	if me.stack.IsValid(index) {
		value := me.stack.Get(index)
		return value.Type()
	}
	return LuaTypeNone
}

// IsNone 判断指定索引处的值是否为None类型
func (me *LuaState) IsNone(index int) bool {
	return me.Type(index) == LuaTypeNone
}

// IsNil 判断指定索引处的值是否为Nil类型
func (me *LuaState) IsNil(index int) bool {
	return me.Type(index) == LuaTypeNil
}

// IsNoneOrNil 判断指定索引处的值是否为None或者Nil类型
func (me *LuaState) IsNoneOrNil(index int) bool {
	return me.IsNone(index) || me.IsNil(index)
}

// IsBoolean 判断指定索引处的值是否为Boolean类型
func (me *LuaState) IsBoolean(index int) bool {
	return me.Type(index) == LuaTypeBoolean
}

// IsString 判断指定索引处的值是否为String或者Number类型
func (me *LuaState) IsString(index int) bool {
	vtype := me.Type(index)
	return vtype == LuaTypeString || vtype == LuaTypeNumber
}

// IsNumber 判断指定索引处的值是否为Number类型
func (me *LuaState) IsNumber(index int) bool {
	return true
}

// IsInteger s
func (me *LuaState) IsInteger(index int) bool {
	dstValue := me.stack.Get(index)
	value := dstValue.Value()
	_, ok := value.(int64)
	return ok
}

// ToBoolean s
func (me *LuaState) ToBoolean(index int) bool {
	return true
}

func (me *LuaState) ToNumber(index int) float64 {
	return 0
}

func (me *LuaState) ToNumberX(index int) (float64, bool) {
	return 0, true
}

func (me *LuaState) ToInteger(index int) int64 {
	return 0
}

func (me *LuaState) ToIntegerX(index int) (int64, bool) {
	return 0, true
}

func (me *LuaState) ToString(index int) string {
	return ""
}

func (me *LuaState) ToStringX(index int) (string, bool) {
	return "", true
}
