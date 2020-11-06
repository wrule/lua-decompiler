package chunk

import (
	"fmt"

	"github.com/fatih/color"
)

// Prototype 函数原型
type Prototype struct {
	source          string
	lineDefined     uint32
	lastLineDefined uint32
	numParams       byte
	isVararg        byte
	maxStackSize    byte
	codes           []uint32
	constants       []interface{}
	upvalues        []Upvalue
	protos          []*Prototype
	lineInfos       []uint32
	locVars         []LocVar
	upvalueNames    []string
}

// Source 获取源文件名
func (me *Prototype) Source() string {
	return me.source
}

// LineDefined 获取开始行号
func (me *Prototype) LineDefined() uint32 {
	return me.lineDefined
}

// LastLineDefined 获取结束行号
func (me *Prototype) LastLineDefined() uint32 {
	return me.lastLineDefined
}

// NumParams 获取固定参数个数
func (me *Prototype) NumParams() byte {
	return me.numParams
}

// IsVararg 是否是Vararg函数
func (me *Prototype) IsVararg() byte {
	return me.isVararg
}

// MaxStackSize 获取需要寄存器数量
func (me *Prototype) MaxStackSize() byte {
	return me.maxStackSize
}

// Codes 获取指令表
func (me *Prototype) Codes() []uint32 {
	return me.codes
}

// Constants 获取常量表
func (me *Prototype) Constants() []interface{} {
	return me.constants
}

// Upvalues 获取Upvalue表
func (me *Prototype) Upvalues() []Upvalue {
	return me.upvalues
}

// Prototypes 获取子函数原型表
func (me *Prototype) Prototypes() []*Prototype {
	return me.protos
}

// LineInfos 获取行号表
func (me *Prototype) LineInfos() []uint32 {
	return me.lineInfos
}

// LocVars 获取局部变量表
func (me *Prototype) LocVars() []LocVar {
	return me.locVars
}

// UpvalueNames 获取Upvalue名列表
func (me *Prototype) UpvalueNames() []string {
	return me.upvalueNames
}

// ListCodes 指令列表信息
func (me *Prototype) ListCodes() []string {
	var codeNum = len(me.Codes())
	var lines = make([]string, codeNum+1)
	lines[0] = fmt.Sprintf("指令数: %d", codeNum)
	for index, code := range me.Codes() {
		pindex := index + 1
		lines[pindex] = fmt.Sprintf(
			"\t%d\t[%d]\t0x%08X",
			pindex,
			me.LineInfos()[index],
			code,
		)
	}
	return lines
}

// ListConstants 打印常量列表
func (me *Prototype) ListConstants() []string {
	var constantNum = len(me.Constants())
	var lines = make([]string, constantNum+1)
	lines[0] = fmt.Sprintf("常量数: %d", constantNum)
	for index := range lines[1:] {
		pindex := index + 1
		lines[pindex] = fmt.Sprintf("\t%d\t", pindex)
	}
	return lines
}

// ListLocVars 局部变量信息
func (me *Prototype) ListLocVars() []string {
	var locVarNum = len(me.LocVars())
	var lines = make([]string, locVarNum+1)
	lines[0] = fmt.Sprintf("局部变量数: %d", locVarNum)
	for index := range lines[1:] {
		pindex := index + 1
		lines[index] = fmt.Sprintf("\t%d\t", pindex)
	}
	return lines
}

// ListUpvalues Upvalue信息
func (me *Prototype) ListUpvalues() []string {
	var upvalueNum = len(me.Upvalues())
	var lines = make([]string, upvalueNum+1)
	lines[0] = fmt.Sprintf("Upvalue数: %d", upvalueNum)
	for index := range lines[1:] {
		pindex := index + 1
		lines[pindex] = fmt.Sprintf("\t%d\t", pindex)
	}
	return lines
}

// List 输出函数原型信息
func (me *Prototype) List() {
	fmt.Printf(
		"<文件名: %s [%d:%d]> (%d 个指令)\n",
		me.Source(),
		me.LineDefined(),
		me.LastLineDefined(),
		len(me.Codes()),
	)
	varargFlag := ""
	if me.IsVararg() > 0 {
		varargFlag = "+"
	}
	fmt.Printf(
		"参数个数: %d%s  寄存器数量: %d  Upvalue数量: %d  局部变量数量: %d  常量数量: %d  子函数数量: %d\n",
		me.NumParams(),
		varargFlag,
		me.MaxStackSize(),
		len(me.Upvalues()),
		len(me.LocVars()),
		len(me.Constants()),
		len(me.Prototypes()),
	)

	for _, line := range me.ListCodes() {
		fmt.Println(line)
	}

	for _, line := range me.ListConstants() {
		fmt.Println(line)
	}

	for _, line := range me.ListLocVars() {
		fmt.Println(line)
	}

	for _, line := range me.ListUpvalues() {
		fmt.Println(line)
	}

	color.Green("结束，颜色测试")
}
