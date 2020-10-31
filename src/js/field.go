package js

// Field Js字段
type Field struct {
	name  string
	value *Value
}

// Name 获取Js字段名称
func (me *Field) Name() string {
	return me.name
}

// Value 获取Js字段值
func (me *Field) Value() *Value {
	return me.value
}

// Type 获取Js字段类型
func (me *Field) Type() EJsType {
	return me.value.Type()
}

// NewJsField 构造函数
func NewJsField(name string, value *Value) *Field {
	return &Field{
		name:  name,
		value: value,
	}
}
