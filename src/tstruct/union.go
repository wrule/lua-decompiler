package tstruct

// Union 结构
type Union struct {
	desc string
}

// Type 结构类型
func (me *Union) Type() ETStructType {
	return TStructUnion
}

// Desc 结构描述
func (me *Union) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Union) Hash() string {
	return "union"
}

// NewUnion Union结构构造函数
func NewUnion(desc string) *Union {
	return &Union{desc}
}
