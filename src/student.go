package main

import (
	"fmt"
)

// Student 学生
type Student struct {
	name string
	sex  bool
	age  int
}

// SexDesc 性别描述
func (me *Student) SexDesc() string {
	if me.sex == true {
		return "男"
	}
	return "女"
}

// ShowMetaInfos 显示基本信息
func (me *Student) ShowMetaInfos() {
	fmt.Printf("姓名: %s\n", me.name)
	fmt.Printf("性别: %s\n", me.SexDesc())
	fmt.Printf("年龄: %d\n", me.age)
}
