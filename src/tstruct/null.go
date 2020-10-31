package tstruct

// Null 结构
type Null struct {
	desc string
}

// Type 结构类型
func (me *Null) Type() ETStructType {
	return TStructNull
}

// Desc 结构描述
func (me *Null) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Null) Hash() string {
	return "null"
}

// NewNull Null结构构造函数
func NewNull(desc string) *Null {
	return &Null{desc}
}
