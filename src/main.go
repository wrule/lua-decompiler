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

	// ls := api.NewLuaState()
	// ls.PushInteger(1)
	// ls.PushString("2.0")
	// ls.PushString("3.0")
	// ls.PushNumber(4.0)
	// ls.List()
	// ls.Arith(api.LuaOpADD)
	// ls.List()
	// ls.Arith(api.LuaOpBNOT)
	// ls.List()
	// ls.Len(2)
	// ls.List()
	// ls.Concat(3)
	// ls.List()
	// ls.PushBoolean(ls.Compare(1, 2, api.LuaOpEQ))
	// ls.List()

	nums := make([]int, 5, 10)
	// nums[9] = 13
	fmt.Println(nums, len(nums), cap(nums))
}
