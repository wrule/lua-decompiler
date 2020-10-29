package student

// Student 结构体
type Student struct {
	name    string
	sex     bool
	age     int
	address string
}

// SexDesc 获取性别的描述
func (stu *Student) SexDesc() string {
	if stu.sex {
		return "男"
	}
	return "女"
}
