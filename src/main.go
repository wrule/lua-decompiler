package main

import (
	"fmt"

	"./js"
	"./tstruct"
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

	fmt.Println("结束 ", value.ObjectFields()[0].Value().ObjectFields()[0].Name())

	fmt.Println(tstruct.TStructNull)
	var nll = tstruct.NewNull("my")
	fmt.Println(nll.Type(), nll.Desc(), nll.Hash())

	fmt.Println(utils.Hash("sdf"))
}
