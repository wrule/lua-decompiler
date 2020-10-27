package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatalln("文件打开失败")
	}
	jsonText := string(bytes)
	fmt.Println(jsonText)

	var any interface{}
	any = []int{1, 2, 3}
	any = 123
	any = "空接口可以存储任何数据"
	fmt.Println(any)
}
