package main

// import (
// 	"./student"
// )

import (
	"fmt"
)

// Student 结构体
type Student struct {
	name string
	sex  bool
	age  int
}

// SexDesc 获取性别的描述
func (stu *Student) SexDesc() string {
	if stu.sex {
		return "男"
	}
	return "女"
}

func main() {
	var stu = Student{"德克萨斯", true, 13}

	fmt.Println(stu.SexDesc())

	// bytes, err := ioutil.ReadFile("./test.json")
	// if err != nil {
	// 	log.Fatalln("文件打开失败")
	// }
	// jsonText := string(bytes)
	// fmt.Println(jsonText)

	// var any interface{}
	// any = []int{1, 2, 3}
	// any = 123
	// any = "空接口可以存储任何数据"
	// fmt.Println(any)
}

// type PolicyType int32

// const (
// 	Policy_MIN PolicyType = iota
// 	Policy_MAX
// 	Policy_MID
// 	Policy_AVG
// )
