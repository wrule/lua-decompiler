package api

type VM interface {
	// LuaState
	// PC 获取当前程序计数器
	PC() int
	// AddPC 修改程序计数器，用于跳转指令
	AddPC(n int)
	// Fetch 取出当前指令，并将PC指向下一条指令
	Fetch() uint32
	// GetConst 将指定常量推入栈顶
	GetConst(index int)
	// GetRK 将指定常量或栈值推入栈顶
	GetRK(index int)
}
