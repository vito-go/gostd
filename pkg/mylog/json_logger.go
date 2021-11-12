package mylog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

// jsonLogger 完全基于标准库实现json日志格式的输出, 高效、简洁.兼容logrus等日志框架.
// For example:
//	introduction := "introduction"
//	mylog.WithField("nihao", "hah").Info("信息")
//	mylog.WithField("name", "zhangsan").WithField("age", 18).Info("个人信息")
//	mylog.WithField("name", "zhangsan").WithField("age", 18).Infof("个人信息: %s", introduction)
//	mylog.JsonOut.Infof("ads: %s", introduction)
// output:
//	{"level":"info","time":"2021/11/11 13:54:47.060","msg":"信息","file_no":"main.go:7","nihao":"hah"}
//	{"level":"info","time":"2021/11/11 13:54:47.060","msg":"个人信息","file_no":"main.go:8","name":"zhangsan","age":18}
//  {"level":"info","time":"2021/11/11 13:54:47.060","msg":"个人信息: introduction","file_no":"main.go:9","name":"zhangsan","age":18}
//	{"level":"info","time":"2021/11/11 13:54:47.060","msg":"ads: introduction","file_no":"main.go:10"}
var jsonLogger = fieldLog{mu: sync.Mutex{}, cfg: &Config{
	InfoWriter:  os.Stdout,
	WarnWriter:  os.Stdout,
	ErrWriter:   os.Stdout,
	PanicWriter: os.Stdout,
	FatalWriter: os.Stdout,
}}

type fieldLog struct {
	mu  sync.Mutex // ensures atomic writes; protects the following fields
	cfg *Config
}
type entry struct {
	fn     func() *entry
	level  level
	time   string
	fileNo string
	msg    string
	fields [][2]interface{}
}

var JsonOut = entry{fn: jsonOut}

func WithField(k, v string) *entry {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
	}
	return &entry{
		time:   time.Now().Format("2006/01/02 15:04:05.000"),
		fileNo: filepath.Base(file) + ":" + strconv.FormatInt(int64(line), 10),
		fields: [][2]interface{}{{k, v}}}
}
func (e *entry) WithField(k string, v interface{}) *entry {
	e.fields = append(e.fields, [2]interface{}{k, v})
	return e
}

type level string

const (
	infoLevel  = level("info")
	warnLevel  = level("warn")
	errLevel   = level("err")
	panicLevel = level("panic")
	fatalLevel = level("fatal")
)

func (e *entry) Info(a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputLn(infoLevel, a...)
}
func (e *entry) Infof(format string, a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputFln(infoLevel, format, a...)
}
func (e *entry) Warnf(format string, a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputFln(warnLevel, format, a...)
}
func (e *entry) Errorf(format string, a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputFln(warnLevel, format, a...)
}
func (e *entry) Panicf(format string, a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputFln(panicLevel, format, a...)
}
func (e *entry) Fatalf(format string, a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputFln(fatalLevel, format, a...)
}

func (e *entry) Warn(a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputLn(warnLevel, a...)
}
func (e *entry) Error(a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputLn(errLevel, a...)
}
func (e *entry) Panic(a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputLn(panicLevel, a...)
}
func (e *entry) Fatal(a ...interface{}) {
	if e.fn != nil {
		*e = *e.fn()
	}
	_ = e.outputLn(fatalLevel, a...)
}

func (e *entry) outputLn(l level, a ...interface{}) error {
	return e.output(l, "", a...)

}
func (e *entry) outputFln(l level, format string, a ...interface{}) error {
	return e.output(l, format, a...)
}
func (e *entry) toJsonStr() string {
	result := fmt.Sprintf(`{"level":"%s","time":"%s","msg":"%s","file_no":"%s"`,
		e.level, e.time, e.msg, e.fileNo)
	for i := 0; i < len(e.fields); i++ {
		switch e.fields[i][1].(type) {
		case string, bool:
			result += fmt.Sprintf(`,"%s":"%s"`, e.fields[i][0], e.fields[i][1].(string))
		case int, int64, int32, float64, float32, uint, uint8, uint16, uint32, uint64:
			result += fmt.Sprintf(`,"%s":%d`, e.fields[i][0], e.fields[i][1])
		default:
			result += fmt.Sprint(e.fields[i][1])
		}
	}
	result += "}"
	return result
}

func (e *entry) output(l level, format string, a ...interface{}) error {
	var msg string
	if format == "" {
		msg = fmt.Sprintln(a...)
		msg = msg[:len(msg)-1] // 去除末尾的\n符号
	} else {
		msg = fmt.Sprintf(format, a...)
	}
	e.msg = msg
	e.level = l
	output := e.toJsonStr()
	output += "\n"
	jsonLogger.mu.Lock()
	defer jsonLogger.mu.Unlock()
	var outWriter io.Writer
	switch l {
	case infoLevel:
		outWriter = jsonLogger.cfg.InfoWriter
	case warnLevel:
		outWriter = jsonLogger.cfg.WarnWriter
	case errLevel:
		outWriter = jsonLogger.cfg.ErrWriter
	case panicLevel:
		outWriter = jsonLogger.cfg.PanicWriter
	case fatalLevel:
		outWriter = jsonLogger.cfg.FatalWriter
	}
	_, err := outWriter.Write(*(*[]byte)(unsafe.Pointer(&output)))
	switch l {
	case panicLevel:
		panic(msg)
	case fatalLevel:
		os.Exit(1)
	}
	return err
}

func jsonOut() *entry {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
	}
	return &entry{
		time:   time.Now().Format("2006/01/02 15:04:05.000"),
		fileNo: filepath.Base(file) + ":" + strconv.FormatInt(int64(line), 10),
	}
}
func (f *fieldLog) setConfig(cfg *Config) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.cfg = cfg
}
