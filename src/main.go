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

	var ck chunk.Chunk

	ck.CheckLoad(reader)

	fmt.Println(ck.MainFunc().Source())

	file.Close()
}
