package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(`{"data":`+`
{
	"name": "jimao",
	"sex": true,
	"tags": ["3", "2", "1", 1]
}
`+`}`), &jsonObj)
	if err != nil {
		log.Fatalln("JSON反序列化失败")
		return
	}

	var value = NewJsValue(jsonObj)

	fmt.Println("结束 ", value.ObjectFields()[0].Value().ObjectFields()[2].Value().Type())
}
