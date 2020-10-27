package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./1.txt", os.O_APPEND|os.O_WRONLY, os.ModeCharDevice)
	if err != nil {
		log.Fatalln(err)
		return
	}
	count, err := file.WriteString("你好")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(count)
	file.Close()
}
