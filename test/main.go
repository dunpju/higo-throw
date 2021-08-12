package main

import (
	"github.com/dengpju/higo-throw/exception"
)

func main()  {
	//自定义日志输出格式
	//throw.LogHandle = func() {
	//	throw.LogFormat = "%s (code: %d)"
	//	throw.LogInfo = fmt.Sprintf(throw.LogFormat, throw.LogPayload.Msg, throw.LogPayload.Code)
	//}
	exception.Throw(exception.Message("uuujjj333"),exception.Message("uuuj11"), exception.Code(1), exception.RealMessage("这才是真实错误"))
	exception.Throw(exception.Message("uuujjj33344"), exception.Code(4))
	exception.Throw(exception.Message("uuujjj33355"), exception.Code(5))
}
