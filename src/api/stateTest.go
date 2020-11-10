package api

import "fmt"

// List 打印状态信息
func (me *LuaState) List() {
	for _, value := range me.stack.slots[:me.GetTop()] {
		fmt.Printf("[%v]", value.Value())
	}
	fmt.Println()
}
