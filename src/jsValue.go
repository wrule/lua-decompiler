package main

import (
	"fmt"
	"reflect"
	"regexp"
)

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
	var jsType = getJsType(value)
	if jsType == JsObject {
		fmt.Println(reflect.TypeOf(value))
		getObjectFields(value.(map[string]interface{}))
	} else if jsType == JsArray {
		fmt.Println(value)
	}
	return &JsValue{
		jsType:       jsType,
		objectFields: []JsField{},
		arrayValues:  []JsValue{},
	}
}

func getObjectFields(value map[string]interface{}) []JsField {
	var mapSize = len(value)
	var result = make([]JsField, mapSize)
	var index = 0
	for itemKey, itemValue := range value {
		result[index] = NewJsField(itemKey, NewJsValue(itemValue))
	}
	return result
}

func getArrayValues() []JsValue {
	return []JsValue{}
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
