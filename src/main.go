package main

import (
	"fmt"
	"log"
	"os"

	"./chunk"
)

func main() {
	file, err := os.OpenFile("./test/1.out", os.O_RDONLY, os.ModeType)
	if err != nil {
		log.Fatalln("打开文件失败")
		return
	}

	var reader = chunk.NewReader(file)

	var header chunk.Header

	header.CheckLoad(reader)

	res, err := reader.ReadUint32()
	if err != nil {
		return
	}
	fmt.Println(res)

	// 1B 4C 75 61
	// 27 76 117 97

	file.Close()
}
