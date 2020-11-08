package chunk

import (
	"fmt"

	"../vm"
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
	constants       []Constant
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
func (me *Prototype) Constants() []Constant {
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

// ListName 名称信息
func (me *Prototype) ListName() []string {
	var lines = make([]string, 1)
	var funcType = "main"
	if me.LineDefined() > 0 {
		funcType = "function"
	}
	lines[0] = fmt.Sprintf(
		"%s <文件名: %s [%d:%d]>(%d 个指令)",
		funcType,
		me.Source(),
		me.LineDefined(),
		me.LastLineDefined(),
		len(me.Codes()),
	)
	return lines
}

// ListMeta 基本信息
func (me *Prototype) ListMeta() []string {
	var lines = make([]string, 2)
	varargFlag := ""
	if me.IsVararg() > 0 {
		varargFlag = "+"
	}
	lines[0] = fmt.Sprintf(
		"参数个数: %d%s\t寄存器数: %d\tUpvalue数量: %d",
		me.NumParams(),
		varargFlag,
		me.MaxStackSize(),
		len(me.Upvalues()),
	)
	lines[1] = fmt.Sprintf(
		"局部变量: %d\t常量数量: %d\t子函数数量: %d",
		len(me.LocVars()),
		len(me.Constants()),
		len(me.Prototypes()),
	)
	return lines
}

// ListCodes 指令列表信息
func (me *Prototype) ListCodes() []string {
	var codeNum = len(me.Codes())
	var lines = make([]string, codeNum+1)
	lines[0] = fmt.Sprintf("指令数(%d):", codeNum)
	for index, code := range me.Codes() {
		pindex := index + 1
		var codeNum = ""
		if index < len(me.LineInfos()) {
			codeNum = fmt.Sprintf("%d", me.LineInfos()[index])
		}
		lines[pindex] = fmt.Sprintf(
			"\t%d.\t[%s]\t%s %s",
			pindex,
			codeNum,
			vm.Instruction(code).OpName(),
			vm.Instruction(code).Operands(),
		)
	}
	return lines
}

// ListConstants 打印常量列表
func (me *Prototype) ListConstants() []string {
	var constantNum = len(me.Constants())
	var lines = make([]string, constantNum+1)
	lines[0] = fmt.Sprintf("常量数(%d):", constantNum)
	for index, constant := range me.Constants() {
		pindex := index + 1
		lines[pindex] = fmt.Sprintf(
			"\t%d.\t%s:\t%s",
			pindex,
			constant.TypeString(),
			constant.ValueString(),
		)
	}
	return lines
}

// ListLocVars 局部变量信息
func (me *Prototype) ListLocVars() []string {
	var locVarNum = len(me.LocVars())
	var lines = make([]string, locVarNum+1)
	lines[0] = fmt.Sprintf("局部变量数(%d):", locVarNum)
	for index, locVar := range me.LocVars() {
		pindex := index + 1
		lines[pindex] = fmt.Sprintf(
			"\t%d\t%s\tStartPC: %d\tEndPC: %d",
			pindex,
			locVar.VarName(),
			locVar.StartPC()+1,
			locVar.EndPC()+1,
		)
	}
	return lines
}

// ListUpvalues Upvalue信息
func (me *Prototype) ListUpvalues() []string {
	var upvalueNum = len(me.Upvalues())
	var lines = make([]string, upvalueNum+1)
	lines[0] = fmt.Sprintf("Upvalue数(%d):", upvalueNum)
	for index, upvalue := range me.Upvalues() {
		pindex := index + 1
		var upvalueName string = ""
		if index < len(me.UpvalueNames()) {
			upvalueName = me.UpvalueNames()[index]
		}
		lines[pindex] = fmt.Sprintf(
			"\t%d.\t%s\tInstack: %d\tIdx: %d",
			pindex,
			upvalueName,
			upvalue.Instack(),
			upvalue.Idx(),
		)
	}
	return lines
}

// ListSubPrototypes 子函数原型信息
func (me *Prototype) ListSubPrototypes() []string {
	var lines []string = []string{}
	lines = append(lines, fmt.Sprintf("子函数数(%d):", len(me.Prototypes())))
	for _, subProto := range me.Prototypes() {
		lines = append(lines, subProto.List()...)
	}
	for index := range lines[1:] {
		lines[index+1] = "\t" + lines[index+1]
	}
	return lines
}

// addFrame 为函数信息添加绿色边框
func (me *Prototype) addFrame(lines []string) []string {
	var frameLines = make([]string, len(lines)+2)
	var index = 0
	frameLines[index] = color.GreenString(".------------------- Func Start")
	for index = 1; index < len(frameLines)-1; index++ {
		frameLines[index] = fmt.Sprintf("%s %s", color.GreenString("|"), lines[index-1])
	}
	frameLines[index] = color.GreenString("'------------------- Func End")
	return frameLines
}

// List 函数原型信息
func (me *Prototype) List() []string {
	var lines []string
	lines = append(lines, me.ListName()...)
	lines = append(lines, me.ListMeta()...)
	lines = append(lines, me.ListCodes()...)
	lines = append(lines, me.ListConstants()...)
	lines = append(lines, me.ListLocVars()...)
	lines = append(lines, me.ListUpvalues()...)
	lines = append(lines, me.ListSubPrototypes()...)
	return me.addFrame(lines)
}

// PrintList 打印信息
func (me *Prototype) PrintList() {
	for _, line := range me.List() {
		fmt.Println(line)
	}
}
