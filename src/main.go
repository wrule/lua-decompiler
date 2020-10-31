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
}
