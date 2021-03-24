package throw

import "gitee.com/dengpju/higo-parameter/parameter"

// 异常接口
type IThrowable interface {
	Exception(parameters ...*parameter.Parameter)
}