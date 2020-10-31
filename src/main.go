package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// fmt.Println("你好世界")
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(`{"data":`+`
null
`+`}`), &jsonObj)
	if err != nil {
		log.Fatalln("JSON反序列化失败")
		return
	}

	var value = NewJsValue(jsonObj)

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
