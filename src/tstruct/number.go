package tstruct

// Number 结构
type Number struct {
	desc string
}

// Type 结构类型
func (me *Number) Type() ETStructType {
	return TStructNumber
}

// Desc 结构描述
func (me *Number) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Number) Hash() string {
	return "null"
}

// NewNumber Number结构构造函数
func NewNumber(desc string) *Number {
	return &Number{desc}
}
