package tstruct

// ITStruct 接口
type ITStruct interface {
	Type() ETStructType
	Desc() string
	Hash() string
}
