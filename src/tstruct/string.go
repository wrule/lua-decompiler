package tstruct

// String 结构
type String struct {
	desc string
}

// Type 结构类型
func (me *String) Type() ETStructType {
	return TStructString
}

// Desc 结构描述
func (me *String) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *String) Hash() string {
	return "null"
}

// NewString String结构构造函数
func NewString(desc string) *String {
	return &String{desc}
}
