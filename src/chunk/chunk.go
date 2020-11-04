package chunk

// Chunk 结构
type Chunk struct {
	header           Header
	sizeUpdatevalues byte
	mainFunc         *Prototype
}
