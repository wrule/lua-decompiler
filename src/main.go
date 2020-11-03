package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("你好Lua")
	file, err := os.OpenFile("./test/1.out", os.O_RDONLY, os.ModeType)
	if err != nil {
		log.Fatalln("打开文件失败")
		return
	}
	info, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(info.Size())
	bytes := make([]byte, 10)
	size, err := file.Read(bytes)
	if err != nil {
		return
	}
	fmt.Println(size)
	for index, item := range bytes {
		fmt.Printf("%d %02X %s\n", index, item, string(item))
	}
	file.Close()
}
