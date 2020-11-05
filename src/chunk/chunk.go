package chunk

// Chunk 结构
type Chunk struct {
	header           Header
	sizeUpdatevalues byte
	mainFunc         *Prototype
}

// CheckLoad 校验并加载Chunk
func (me *Chunk) CheckLoad(reader *Reader) {
	me.header.CheckLoad(reader)
	me.sizeUpdatevalues = reader.ReadByte()
	me.mainFunc = reader.ReadPrototype("")
}
