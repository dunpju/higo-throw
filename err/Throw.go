package err

import (
	"gitee.com/dengpju/higo-parameter/parameter"
	"github.com/dengpju/higo-logger/logger"
	"runtime"
)

type LogContent struct {
	Msg, File  string
	Code, Line int
}

type Throwable struct{}

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
