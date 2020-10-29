package main

// JsValue Js值
type JsValue struct {
	jsType       EJsType
	objectFields []JsField
	arrayValues  []JsValue
}

// Type 获取Js值的类型
func (me *JsValue) Type() EJsType {
	return me.jsType
}

// ObjectFields 获取JsObject类型的字段列表
func (me *JsValue) ObjectFields() []JsField {
	return me.objectFields
}

// ArrayValues 获取JsArray类型的值列表
func (me *JsValue) ArrayValues() []JsValue {
	return me.arrayValues
}

// NewJsValue 构造函数
func NewJsValue(value interface{}) *JsValue {
	return &JsValue{
		jsType:       JsUndefined,
		objectFields: []JsField{},
		arrayValues:  []JsValue{},
	}
}
