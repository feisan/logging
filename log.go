package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

type (
	LogLevel int
)

const (
	LOG_CRITICAL = LogLevel(50)
	LOG_ERROR    = LogLevel(40)
	LOG_WARNING  = LogLevel(30)
	LOG_INFO     = LogLevel(20)
	LOG_DEBUG    = LogLevel(10)
	LOG_NOSET    = LogLevel(0)
)

////////////////////////////////////////////////////////////////////////////////
func SetLevel(lvl LogLevel) {
	_logger.SetLevel(lvl)
}

func SetPrefix(prefix string) {
	_logger.SetPrefix(prefix)
}

func SetFlags(flag int) {
	_logger.SetFlags(flag)
}

func Debug(v ...interface{}) {
	_logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	_logger.Debugf(format, v...)
}

func Info(v ...interface{}) {
	_logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	_logger.Infof(format, v...)
}

func Warning(v ...interface{}) {
	_logger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	_logger.Warningf(format, v...)
}

func Error(v ...interface{}) {
	_logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	_logger.Errorf(format, v...)
}

func Critical(v ...interface{}) {
	_logger.Critical(v...)
}

func Criticalf(format string, v ...interface{}) {
	_logger.Criticalf(format, v...)
}

type logger struct {
	llogger   *log.Logger
	lvl       LogLevel
	highlight bool
}

func (self *logger) SetLevel(lvl LogLevel) {
	self.lvl = lvl
}

func (self *logger) SetPrefix(prefix string) {
	self.llogger.SetPrefix(prefix)
}

func (self *logger) SetFlags(flag int) {
	self.llogger.SetFlags(flag)
}

func (self *logger) Debug(v ...interface{}) {
	self.log(LOG_DEBUG, v...)
}

func (self *logger) Debugf(format string, v ...interface{}) {
	self.logf(LOG_DEBUG, format, v...)
}

func (self *logger) Info(v ...interface{}) {
	self.log(LOG_INFO, v...)
}

func (self *logger) Infof(format string, v ...interface{}) {
	self.logf(LOG_INFO, format, v...)
}

func (self *logger) Warning(v ...interface{}) {
	self.log(LOG_WARNING, v...)
}

func (self *logger) Warningf(format string, v ...interface{}) {
	self.logf(LOG_WARNING, format, v...)
}

func (self *logger) Error(v ...interface{}) {
	self.log(LOG_ERROR, v...)
}

func (self *logger) Errorf(format string, v ...interface{}) {
	self.logf(LOG_ERROR, format, v...)
}

func (self *logger) Critical(v ...interface{}) {
	self.log(LOG_CRITICAL, v...)
}

func (self *logger) Criticalf(format string, v ...interface{}) {
	self.logf(LOG_CRITICAL, format, v...)
}

func (self *logger) log(lvl LogLevel, v ...interface{}) {
	if lvl < self.lvl {
		return
	}

	vv := make([]interface{}, len(v)+2)
	logStr, logColor := getLogConfig(lvl)
	if self.highlight && logColor != "" {
		vv[0] = logColor + "[" + logStr + "]"
		copy(vv[1:], v)
		vv[len(v)+1] = "\033[1;0m"
	} else {
		vv[0] = "[" + logStr + "]"
		copy(vv[1:], v)
		vv[len(v)+1] = ""
	}
	s := fmt.Sprintln(vv...)
	self.llogger.Output(4, s)
}

func (self *logger) logf(lvl LogLevel, format string, v ...interface{}) {
	if lvl < self.lvl {
		return
	}
	logStr, logColor := getLogConfig(lvl)
	var s string
	if self.highlight {
		s = logColor + "[" + logStr + "] " + fmt.Sprintf(format, v...) + "\033[1;0m"
	} else {
		s = "[" + logStr + "] " + fmt.Sprintf(format, v...)
	}
	self.llogger.Output(4, s)
}

func getLogConfig(lvl LogLevel) (string, string) {
	switch lvl {
	case LOG_CRITICAL:
		return "critical", "\033[1;31m"
	case LOG_ERROR:
		return "error", "\033[1;31m"
	case LOG_WARNING:
		return "warning", "\033[1;33m"
	case LOG_INFO:
		return "info", "\033[1;32m"
	case LOG_DEBUG:
		return "debug", ""
	}
	return "unknown", ""
}

func NewLogger(w io.Writer, prefix string) *logger {
	return &logger{log.New(w, prefix, log.LstdFlags|log.Lshortfile), LOG_NOSET, true}
}

func New() *logger {
	return NewLogger(os.Stdout, "")
}

var _logger *logger = New()
