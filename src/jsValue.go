package main

// JsValue Js值
type JsValue struct {
	jsType       EJsType
	objectFields []JsField
	arrayValues  []JsValue
}
