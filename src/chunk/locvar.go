package chunk

// LocVar 局部变量
type LocVar struct {
	varName string
	startPC uint32
	endPC   uint32
}
