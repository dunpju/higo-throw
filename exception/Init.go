package exception

import (
	"fmt"
	"gitee.com/dengpju/higo-parameter/parameter"
	"github.com/dengpju/higo-utils/utils"
	"sync"
)

const (
	REAL    = "real"
	MESSAGE = "message"
	CODE    = "code"
	DATA    = "data"
)

var (
	once          sync.Once
	ThrowInstance IThrowable         //异常实例(可自定义)
	Handle        parameter.Callable //参数处理函数(可自定义)
	LogHandle     func()             //日志处理函数(可自定义)
	ArrayMap      *utils.ArrayMap
	LogPayload    *LogContent
	LogFormat     = "%s (code: %d) %s at %s:%d" //日志格式(可自定义)
	LogInfo       = ""
)

func init() { 
	once.Do(func() {
		ThrowInstance = &Throwable{}
		LogPayload = &LogContent{}
		ArrayMap = utils.Array()
		//初始化参数处理函数
		Handle = func(p *parameter.Parameter) {
			if p.Name == REAL {
				LogPayload.Real = ErrorToString(p.Value)
			}
			if p.Name == MESSAGE {
				LogPayload.Msg = ErrorToString(p.Value)
			}
			if p.Name == CODE {
				LogPayload.Code = p.Value.(int)
			}
			ArrayMap.Put(p.Name, p.Value)
		}
		//初始化日志处理函数
		LogHandle = func() {
			LogInfo = fmt.Sprintf(LogFormat, LogPayload.Msg, LogPayload.Code, LogPayload.Real, LogPayload.File, LogPayload.Line)
		}
	})
}
