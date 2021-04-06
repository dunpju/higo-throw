package throw

import (
	"fmt"
	"gitee.com/dengpju/higo-parameter/parameter"
	"github.com/dengpju/higo-logger/logger"
	"github.com/dengpju/higo-utils/utils"
	"runtime"
	"sync"
)

const (
	MESSAGE = "message"
	CODE    = "code"
	DATA    = "data"
)

var (
	once          sync.Once
	ThrowInstance IThrowable         //异常实例(可自定义)
	Handle        parameter.Callable //参数处理函数(可自定义)
	LogHandle     func()             //日志处理函数(可自定义)
	MapString     utils.MapString
	LogPayload    *LogContent
	LogFormat     = "%s (code: %d) at %s:%d" //日志格式(可自定义)a
	LogInfo       = ""
)

type LogContent struct {
	Msg, File  string
	Code, Line int
}

type Throwable struct{}

func init() {
	once.Do(func() {
		ThrowInstance = &Throwable{}
		LogPayload = &LogContent{}
		MapString = make(utils.MapString)
		//初始化参数处理函数
		Handle = func(p *parameter.Parameter) {
			if p.Name == MESSAGE {
				LogPayload.Msg = ErrorToString(p.Value)
			}
			if p.Name == CODE {
				LogPayload.Code = p.Value.(int)
			}
			MapString.Put(p.Name, p.Value)
		}
		//初始化日志处理函数
		LogHandle = func() {
			LogInfo = fmt.Sprintf(LogFormat, LogPayload.Msg, LogPayload.Code, LogPayload.File, LogPayload.Line)
		}
	})
}

func (this *Throwable) Exception(parameters ...*parameter.Parameter) {
	_, file, line, _ := runtime.Caller(2)
	LogPayload.File = file
	LogPayload.Line = line
	parameter.Parameters(parameters).ForEach(Handle)
	LogHandle()
	logger.Logrus.Init()
	logger.Logrus.Info(LogInfo)
	panic(MapString)
}

func Message(value interface{}) *parameter.Parameter {
	return parameter.New(MESSAGE, value)
}

func Code(value int) *parameter.Parameter {
	return parameter.New(CODE, value)
}

func Data(value interface{}) *parameter.Parameter {
	return parameter.New(DATA, value)
}

// 抛出异常
func Throw(parameters ...*parameter.Parameter) {
	ThrowInstance.Exception(parameters...)
}

// recover 转 string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	case []uint8:
		return BytesToString(r.([]uint8))
	default:
		return r.(string)
	}
}

// []uint8 转 string
func BytesToString(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
