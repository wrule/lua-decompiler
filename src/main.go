package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("程序开始")
	stu := Student{"鸡毛巾", true, 13}
	stu.ShowMetaInfos()
	bytes, err := json.Marshal(stu)
	if err != nil {
		log.Fatalln("JSON转换失败")
	}
	jsonText := string(bytes)
	fmt.Println(jsonText)
	var stu2 Student
	err = json.Unmarshal([]byte(jsonText), &stu2)
	if err != nil {
		log.Fatalln("JSON反序列化失败")
	}
	stu2.ShowMetaInfos()
}
