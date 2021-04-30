package main

import (
	"github.com/dengpju/higo-throw/error"
)

func main()  {
	//自定义日志输出格式
	//throw.LogHandle = func() {
	//	throw.LogFormat = "%s (code: %d)"
	//	throw.LogInfo = fmt.Sprintf(throw.LogFormat, throw.LogPayload.Msg, throw.LogPayload.Code)
	//}
	error.Throw(error.Message("uuujjj333"), error.Code(1))
	error.Throw(error.Message("uuujjj33344"), error.Code(4))
	error.Throw(error.Message("uuujjj33355"), error.Code(5))
}
