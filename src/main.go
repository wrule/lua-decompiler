package main

import (
	"fmt"
)

func main() {
	fmt.Println("程序开始")
	stu := Student{"鸡毛巾", true, 13}
	stu.ShowMetaInfos()
}
