package main

import (
	"fmt"

	"./student"
)

// Test 测试结构
type Test struct {
	name string
	age  int
}

func main() {
	fmt.Print("你好,世界\n")
	t := Test{"可是\n", 13}
	stu := student.Student{
		Name:    "特朗普\n",
		Sex:     true,
		Age:     13,
		Address: "",
		Remark:  "",
	}
	fmt.Print(t.name)
	stu.ShowName()
}
