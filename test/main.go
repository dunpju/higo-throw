package main

import (
	"github.com/dengpju/higo-throw/err"
)

func main()  {
	//自定义日志输出格式
	//throw.LogHandle = func() {
	//	throw.LogFormat = "%s (code: %d)"
	//	throw.LogInfo = fmt.Sprintf(throw.LogFormat, throw.LogPayload.Msg, throw.LogPayload.Code)
	//}
	err.Throw(err.Message("uuujjj333"), err.Code(1))
	err.Throw(err.Message("uuujjj33344"), err.Code(4))
	err.Throw(err.Message("uuujjj33355"), err.Code(5))
}
