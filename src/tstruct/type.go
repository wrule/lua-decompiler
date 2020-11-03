package tstruct

// ETStructType s
type ETStructType string

const (
	// TStructNull Null类型结构
	TStructNull ETStructType = "Null"
	// TStructUndefined Undefined类型结构
	TStructUndefined ETStructType = "Undefined"
	// TStructBoolean 布尔类型结构
	TStructBoolean ETStructType = "Boolean"
	// TStructNumber Number类型结构
	TStructNumber ETStructType = "Number"
	// TStructString String类型结构
	TStructString ETStructType = "String"
	// TStructDate Date类型结构
	TStructDate ETStructType = "Date"
	// TStructObject Object类型结构
	TStructObject ETStructType = "Object"
	// TStructArray Array类型结构
	TStructArray ETStructType = "Array"
	// TStructTuple Tuple类型结构
	TStructTuple ETStructType = "Tuple"
	// TStructUnion Union类型结构
	TStructUnion ETStructType = "Union"
)
