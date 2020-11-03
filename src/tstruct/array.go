package tstruct

// Array 结构
type Array struct {
	desc string
}

// Type 结构类型
func (me *Array) Type() ETStructType {
	return TStructArray
}

// Desc 结构描述
func (me *Array) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Array) Hash() string {
	return "array"
}

// NewArray Array结构构造函数
func NewArray(desc string) *Array {
	return &Array{desc}
}
