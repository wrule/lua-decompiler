package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"./api"
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

	// fmt.Println(len(vm.InstructionInfos))

	state := api.NewLuaState()
	state.PushBoolean(true)
	state.List()
	state.PushInteger(10)
	state.List()
	state.PushNil()
	state.List()
	state.PushString("hello")
	state.List()
	state.PushValue(-4)
	state.List()
	state.Replace(3)
	state.List()
	state.SetTop(6)
	state.List()
	state.Remove(-3)
	state.List()
	state.SetTop(-5)
	state.List()

	fmt.Println(math.Floor(5.0 / -3.0))

	// fmt.Println(strconv.ParseInt("234.0", 10, 64))
	fmt.Println(strconv.ParseFloat(".123s", 64))
}
