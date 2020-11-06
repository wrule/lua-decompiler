package chunk

// Upvalue 元素
type Upvalue struct {
	instack byte
	idx     byte
}

// Instack s
func (me *Upvalue) Instack() byte {
	return me.instack
}

// Idx s
func (me *Upvalue) Idx() byte {
	return me.idx
}
