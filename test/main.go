package main

import (
	"gitee.com/dengpju/higo-code/code"
	"github.com/dengpju/higo-throw/exception"
)

type TestCode int64

func (this TestCode) Message() string {
	return code.Get(this)
}

const (
	None TestCode = 1
	CPU  TestCode = 2 // 中央处理器
	GPU  TestCode = 3 // 图形处理器
)

func (this TestCode) Register() code.Message {
	return code.Container().
		Put(None, "None1").
		Put(CPU, "CPU1").
		Put(GPU, "GPU1")
}

func main()  {
	//自定义日志输出格式
	//throw.LogHandle = func() {
	//	throw.LogFormat = "%s (code: %d)"
	//	throw.LogInfo = fmt.Sprintf(throw.LogFormat, throw.LogPayload.Msg, throw.LogPayload.Code)
	//}
	//exception.Throw(exception.Code(GPU))
	exception.Throw(exception.Message("uuujjj333"),exception.Message("uuuj11"), exception.Code("1"), exception.RealMessage("这才是真实错误"))
	exception.Throw(exception.Message("uuujjj33344"), exception.Code(4))
	exception.Throw(exception.Message("uuujjj33355"), exception.Code(5))
}
