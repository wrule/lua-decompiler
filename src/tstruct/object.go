package tstruct

// Object 结构
type Object struct {
	desc string
}

// Type 结构类型
func (me *Object) Type() ETStructType {
	return TStructObject
}

// Desc 结构描述
func (me *Object) Desc() string {
	return me.desc
}

// Hash 结构Hash值
func (me *Object) Hash() string {
	return "object"
}

// NewObject Object结构构造函数
func NewObject(desc string) *Object {
	return &Object{desc}
}
