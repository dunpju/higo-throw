package main

import (
	"github.com/dengpju/higo-throw/throw"
)

func main()  {
	//自定义日志输出格式
	//throw.LogHandle = func() {
	//	throw.LogFormat = "%s (code: %d)"
	//	throw.LogInfo = fmt.Sprintf(throw.LogFormat, throw.LogPayload.Msg, throw.LogPayload.Code)
	//}
	throw.Throw(throw.Message("uuujjj333"), throw.Code(1))
	throw.Throw(throw.Message("uuujjj33344"), throw.Code(4))
	throw.Throw(throw.Message("uuujjj33355"), throw.Code(5))
}
