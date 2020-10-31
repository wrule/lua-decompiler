package main

import (
	"fmt"

	"./js"
	"./utils"
)

func main() {
	var jsonText = `
{
	"name": "jimao",
	"sex": true,
	"tags": ["3", "2", "1", 1]
}	
`

	var jsonObj = utils.WrapParse(jsonText)

	var value = js.NewJsValue(jsonObj)

	fmt.Println("结束 ", value.Type())
}
