package tstruct

// Tuple 结构
type Tuple struct {
	desc string
}

// Type 结构类型
func (me *Tuple) Type() ETStructType {
	return TStructTuple
}

// Desc 结构描述
func (me *Tuple) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Tuple) Hash() string {
	return "tuple"
}

// NewTuple Tuple结构构造函数
func NewTuple(desc string) *Tuple {
	return &Tuple{desc}
}
