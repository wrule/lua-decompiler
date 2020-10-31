package main

// JsField Js字段
type JsField struct {
	name  string
	value JsValue
}

// Name 获取Js字段名称
func (me *JsField) Name() string {
	return me.name
}

// Value 获取Js字段值
func (me *JsField) Value() JsValue {
	return me.value
}

// Type 获取Js字段类型
func (me *JsField) Type() EJsType {
	return me.value.Type()
}

// NewJsField 构造函数
func NewJsField(name string, value JsValue) *JsField {
	return &JsField{
		name:  name,
		value: value,
	}
}
