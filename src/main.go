package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// fmt.Println("你好世界")
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(`
{
	"name": "nihao",
	"sex": true,
	"age": 13,
	"tags": ["c", "c++", "go", 1, {}],
	"meta": {
		"num": 13,
		"keys": [true, false, true]
	},
	"date":"2020-10-30T15:59:27.225Z",
	"ssss": null
}
`), &jsonObj)
	if err != nil {
		log.Fatalln("JSON反序列化失败")
		return
	}

	value := NewJsValue(jsonObj)

	fmt.Println(value.Type())

	// fmt.Printf("%T\n", jsonObj["name"])
	// for key := range jsonObj {
	// 	value := jsonObj[key]
	// 	fmt.Println(
	// 		key,
	// 		getJsType(value),
	// 		// value,
	// 		// reflect.TypeOf(value),
	// 		// reflect.ValueOf(value).IsValid(),
	// 		// reflect.TypeOf(value).Kind() == reflect.Slice,
	// 		// reflect.TypeOf(value).Kind() == reflect.Map,
	// 		// reflect.TypeOf(value).Kind(),
	// 	)
	// }
}
