package main

// EJsType Js值的类型枚举
type EJsType string

const (
	// JsUndefined Js值 Undefined类型
	JsUndefined EJsType = "undefined"
	// JsNull Js值 Null类型
	JsNull EJsType = "null"
	// JsBoolean Js值 Boolean类型
	JsBoolean EJsType = "boolean"
	// JsNumber Js值 Number类型
	JsNumber EJsType = "number"
	// JsString Js值 String类型
	JsString EJsType = "string"
	// JsDate Js值 Date类型
	JsDate EJsType = "Date"
	// JsObject Js值 Object类型
	JsObject EJsType = "object"
	// JsArray Js值 Array类型
	JsArray EJsType = "array"
	// JsUnknow Js值 未知类型
	JsUnknow EJsType = "any"
)
