package chunk

// Upvalue 元素
type Upvalue struct {
	instack byte
	idx     byte
}
