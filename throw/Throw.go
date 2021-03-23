package throw

import (
	"fmt"
	"gitee.com/dengpju/higo-parameter/parameter"
	"github.com/dengpju/higo-logger/logger"
	"github.com/dengpju/higo-utils/utils"
	"runtime"
	"sync"
)

var (
	throwable *Throwable
	once      sync.Once
)

const (
	MESSAGE = "message"
	CODE    = "code"
	DATA    = "data"
)

var (
	Handle     parameter.Callable
	LogHandle  func()
	MapString  utils.MapString
	LogPayload *LogContent
	LogFormat  = "%s (code: %d) at %s:%d"
	LogInfo    = ""
)

type LogContent struct {
	Msg, File  string
	Code, Line int
}

type Throwable struct{}

func init() {
	once.Do(func() {
		throwable = &Throwable{}
		LogPayload = &LogContent{}
		MapString = make(utils.MapString)
		Handle = func(p *parameter.Parameter) {
			if p.Name == MESSAGE {
				LogPayload.Msg = ErrorToString(p.Value)
			}
			if p.Name == CODE {
				LogPayload.Code = p.Value.(int)
			}
			MapString.Put(p.Name, p.Value)
		}
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
	throwable.Exception(parameters...)
}

// recover 转 string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	case []uint8:
		return B2S(r.([]uint8))
	default:
		return r.(string)
	}
}

// []uint8 转 string
func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
