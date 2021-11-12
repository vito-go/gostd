package mylog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

const logFlag = log.Lmicroseconds | log.Ldate | log.Lshortfile

//  infoPrefix warnPrefix errPrefix panicPrefix fatalPrefix 日志输出前缀.
const (
	infoPrefix  = "[INFO] "
	warnPrefix  = "[WARN] "
	errPrefix   = "[ERROR] "
	panicPrefix = "[PANIC] "
	fatalPrefix = "[FATAL] "
)

/*
 开头部分：\033[显示方式;前景色;背景色m + 结尾部分：\033[0m
分类和数值表示的参数含义：
显示方式: 0（默认值）、1（高亮）、22（非粗体）、4（下划线）、24（非下划线）、 5（闪烁）、25（非闪烁）、7（反显）、27（非反显）
前景色: 30（黑色）、31（红色）、32（绿色）、 33（黄色）、34（蓝色）、35（洋 红）、36（青色）、37（白色）
背景色: 40（黑色）、41（红色）、42（绿色）、 43（黄色）、44（蓝色）、45（洋 红）、46（青色）、47（白色）
*/
var (
	infoLogger  = log.New(os.Stdout, infoPrefix, logFlag)
	warnLogger  = log.New(os.Stdout, warnPrefix, logFlag)
	errLogger   = log.New(os.Stdout, errPrefix, logFlag)
	panicLogger = log.New(os.Stdout, panicPrefix, logFlag)
	fatalLogger = log.New(os.Stdout, fatalPrefix, logFlag)
	once        = new(sync.Once)
)

type Config struct {
	InfoWriter  io.Writer
	WarnWriter  io.Writer
	ErrWriter   io.Writer
	PanicWriter io.Writer
	FatalWriter io.Writer
}

// Init set the Writer. It does nothing when after doing first.
func Init(c *Config) {
	once.Do(func() {
		if c.InfoWriter != nil {
			infoLogger.SetOutput(c.InfoWriter)
			infoLogger.SetPrefix(infoPrefix)
		}
		if c.WarnWriter != nil {
			warnLogger.SetOutput(c.WarnWriter)
			warnLogger.SetPrefix(warnPrefix)
		}
		if c.ErrWriter != nil {
			errLogger.SetOutput(c.ErrWriter)
			errLogger.SetPrefix(errPrefix)
		}
		if c.PanicWriter != nil {
			panicLogger.SetOutput(c.PanicWriter)
			panicLogger.SetPrefix(panicPrefix)
		}
		if c.FatalWriter != nil {
			fatalLogger.SetOutput(c.FatalWriter)
			fatalLogger.SetPrefix(fatalPrefix)
		}
		jsonLogger.setConfig(c)
	})
}

func Info(v ...interface{}) {
	_ = infoLogger.Output(2, fmt.Sprintln(v...))
}
func Infof(format string, v ...interface{}) {
	_ = infoLogger.Output(2, fmt.Sprintf(format, v...))
}
func Warn(v ...interface{}) {
	_ = warnLogger.Output(2, fmt.Sprintln(v...))
}
func Warnf(format string, v ...interface{}) {
	_ = warnLogger.Output(2, fmt.Sprintf(format, v...))
}
func Error(v ...interface{}) {
	_ = errLogger.Output(2, fmt.Sprintln(v...))
}
func Errorf(format string, v ...interface{}) {
	_ = errLogger.Output(2, fmt.Sprintf(format, v...))
}
func Panic(v ...interface{}) {
	s := fmt.Sprintln(v...)
	_ = panicLogger.Output(2, s)
	panic(s)
}
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_ = panicLogger.Output(2, s)
	panic(s)
}
func Fatal(v ...interface{}) {
	_ = fatalLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
func Fatalf(format string, v ...interface{}) {
	_ = fatalLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func init() {
	if runtime.GOOS != "windows" {
		// 仅仅当为非windows平台,且为标准输出时才为带颜色输出. json不支持
		infoLogger.SetPrefix("\033[1;32m[INFO] \033[0m")
		warnLogger.SetPrefix("\033[1;33m[WARN] \033[0m")
		errLogger.SetPrefix("\033[1;31m[ERROR] \033[0m")
		panicLogger.SetPrefix("\033[1;37;31m[PANIC] \033[0m")
		fatalLogger.SetPrefix("\033[5;32m[FATAL] \033[0m")
	}
}
