package student

import "fmt"

// Student 结构
type Student struct {
	Name    string
	Sex     bool
	Age     int
	Address string
	Remark  string
}

// ShowName 函数
func (me Student) ShowName() {
	fmt.Print(me.Name)
}
