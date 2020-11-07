package main

import (
	"fmt"
	"log"
	"os"

	"./chunk"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("请输入需反编译的chunk文件")
	}
	var fileName = os.Args[1]
	fmt.Println(fileName)
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModeType)
	if err != nil {
		log.Fatalln("打开文件失败")
		return
	}
	var reader = chunk.NewReader(file)
	var ck chunk.Chunk
	ck.CheckLoad(reader)
	ck.PrintList()
	file.Close()
}
