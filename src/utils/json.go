package json

// Wrap 在JSON外层包裹一个对象结构
func Wrap(json string) string {
	return `{"data":` + json + "}"
}
