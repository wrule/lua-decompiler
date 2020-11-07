package chunk

// Chunk 结构
type Chunk struct {
	header           Header
	sizeUpdatevalues byte
	mainFunc         *Prototype
}

// Header 获取头部信息
func (me *Chunk) Header() Header {
	return me.header
}

// SizeUpdatevalues 获取SizeUpdatevalues
func (me *Chunk) SizeUpdatevalues() byte {
	return me.sizeUpdatevalues
}

// MainFunc 获取主函数
func (me *Chunk) MainFunc() *Prototype {
	return me.mainFunc
}

// CheckLoad 校验并加载Chunk
func (me *Chunk) CheckLoad(reader *Reader) {
	me.header.CheckLoad(reader)
	me.sizeUpdatevalues = reader.ReadByte()
	me.mainFunc = reader.ReadPrototype("")
}

// PrintList 输出Chunk信息
func (me *Chunk) PrintList() {
	me.MainFunc().PrintList()
}
