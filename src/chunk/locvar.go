package chunk

// LocVar 局部变量
type LocVar struct {
	varName string
	startPC uint32
	endPC   uint32
}

// VarName 局部变量名称
func (me *LocVar) VarName() string {
	return me.varName
}

// StartPC s
func (me *LocVar) StartPC() uint32 {
	return me.startPC
}

// EndPC s
func (me *LocVar) EndPC() uint32 {
	return me.endPC
}
