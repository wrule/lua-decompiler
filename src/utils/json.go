package utils

import (
	"encoding/json"
)

// Stringify JSON序列化
func Stringify(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic("JSON序列化错误")
	}
	return string(bytes)
}

// Parse JSON反序列化
func Parse(jsonText string) map[string]interface{} {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(jsonText), &jsonObj)
	if err != nil {
		panic("JSON反序列化错误")
	}
	return jsonObj
}

// Wrap 在JSON外层包裹一个对象结构
func Wrap(jsonText string) string {
	return `{"data":` + jsonText + "}"
}

// WrapStringify 自带拆包的JSON序列化
func WrapStringify(value interface{}) string {
	return Stringify(value.(map[string]interface{})["data"])
}

// WrapParse 自带包裹的JSON反序列化
func WrapParse(jsonText string) map[string]interface{} {
	return Parse(Wrap(jsonText))
}
