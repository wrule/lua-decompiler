package js

import (
	"reflect"
	"regexp"
)

// Value Js值
type Value struct {
	jsType       EJsType
	objectFields []*Field
	arrayValues  []*Value
}

// Type 获取Js值的类型
func (me *Value) Type() EJsType {
	return me.jsType
}

// ObjectFields 获取JsObject类型的字段列表
func (me *Value) ObjectFields() []*Field {
	return me.objectFields
}

// ArrayValues 获取JsArray类型的值列表
func (me *Value) ArrayValues() []*Value {
	return me.arrayValues
}

// NewJsValue 构造函数
func NewJsValue(value interface{}) *Value {
	var jsType = getJsType(value)
	var objectFields []*Field
	var arrayValues []*Value
	if jsType == JsObject {
		objectFields = getObjectFields(value.(map[string]interface{}))
	} else if jsType == JsArray {
		arrayValues = getArrayValues(value.([]interface{}))
	}
	return &Value{
		jsType:       jsType,
		objectFields: objectFields,
		arrayValues:  arrayValues,
	}
}

// getObjectFields 获取对象字段列表
func getObjectFields(value map[string]interface{}) []*Field {
	var mapSize = len(value)
	var result = make([]*Field, mapSize)
	var index = 0
	for itemKey, itemValue := range value {
		result[index] = NewJsField(itemKey, NewJsValue(itemValue))
		index++
	}
	return result
}

// getArrayValues 获取数组值列表
func getArrayValues(value []interface{}) []*Value {
	var arraySize = len(value)
	var result = make([]*Value, arraySize)
	for index, arrayValue := range value {
		result[index] = NewJsValue(arrayValue)
		index++
	}
	return result
}

// jsDateRegexpClosure 用于封装Date正则表达式的闭包函数
func jsDateRegexpClosure() func() *regexp.Regexp {
	jsDateRegexp, err := regexp.Compile("^\\d+-\\d+-\\d+T\\d+:\\d+:\\d+.\\d+Z$")
	if err != nil {
		panic("正则表达式解析错误")
	}
	return func() *regexp.Regexp {
		return jsDateRegexp
	}
}

// getJsType 传入空接口类型的值获取其Js类型
func getJsType(value interface{}) EJsType {
	if value == nil {
		return JsNull
	}
	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Bool:
		return JsBoolean
	case reflect.Float64:
		return JsNumber
	case reflect.String:
		if jsDateRegexpClosure()().MatchString(value.(string)) {
			return JsDate
		}
		return JsString
	case reflect.Map:
		return JsObject
	case reflect.Slice:
		return JsArray
	default:
		return JsUnknow
	}
}
