package exception

import (
	"fmt"
	"gitee.com/dengpju/higo-code/code"
	"gitee.com/dengpju/higo-parameter/parameter"
	"github.com/dengpju/higo-utils/utils"
	"strconv"
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
				if ic, ok := p.Value.(code.ICode); ok {
					cm := code.New(ic)
					LogPayload.Code = int(cm.Code)
					LogPayload.Msg = cm.Message
					ArrayMap.Put(MESSAGE, LogPayload.Msg)
				} else if sc, ok := p.Value.(string); ok {
					si, err := strconv.Atoi(sc)
					if err != nil {
						panic(err)
					}
					LogPayload.Code = si
				} else {
					LogPayload.Code = p.Value.(int)
				}
			}
			ArrayMap.Put(p.Name, p.Value)
		}
		//初始化日志处理函数
		LogHandle = func() {
			LogInfo = fmt.Sprintf(LogFormat, LogPayload.Msg, LogPayload.Code, LogPayload.Real, LogPayload.File, LogPayload.Line)
		}
	})
}
