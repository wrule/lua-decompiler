package tstruct

// Boolean 结构
type Boolean struct {
	desc string
}

// Type 结构类型
func (me *Boolean) Type() ETStructType {
	return TStructBoolean
}

// Desc 结构描述
func (me *Boolean) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Boolean) Hash() string {
	return "boolean"
}

// NewBoolean Boolean结构构造函数
func NewBoolean(desc string) *Boolean {
	return &Boolean{desc}
}
