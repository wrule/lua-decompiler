package chunk

// Constant 常量空接口
type Constant struct {
	ctype EConstantType
	value interface{}
}

// Type 常量的类型
func (me *Constant) Type() EConstantType {
	return me.ctype
}

// Value 常量的值
func (me *Constant) Value() interface{} {
	return me.value
}

// ValueString 常量的值用字符串形式表达
func (me *Constant) ValueString() string {
	return ""
}

// EConstantType Lua常量类型枚举
type EConstantType byte

const (
	// ConstantTypeNil nil常量类型
	ConstantTypeNil EConstantType = 0x00
	// ConstantTypeBoolean 布尔值常量类型
	ConstantTypeBoolean EConstantType = 0x01
	// ConstantTypeNumber 浮点数常量类型
	ConstantTypeNumber EConstantType = 0x03
	// ConstantTypeInteger 整数常量类型
	ConstantTypeInteger EConstantType = 0x13
	// ConstantTypeShortStr 短字符串常量类型
	ConstantTypeShortStr EConstantType = 0x04
	// ConstantTypeLongStr 长字符串常量类型
	ConstantTypeLongStr EConstantType = 0x14
)
