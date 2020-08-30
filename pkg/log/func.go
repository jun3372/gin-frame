package log

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func log() *logrus.Entry {
	_, file, line, _ := runtime.Caller(2)
	caller := fmt.Sprintf("%s:%d", file, line)
	return GetLog().WithField("caller", caller)
}

func Log(level logrus.Level, args ...interface{}) {
	if log().Logger.IsLevelEnabled(level) {
		log().Log(level, fmt.Sprint(args...))
	}
}

func Trace(args ...interface{}) {
	log().Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	log().Log(logrus.DebugLevel, args...)
}

func Print(args ...interface{}) {
	log().Info(args...)
}

func Info(args ...interface{}) {
	log().Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	log().Log(logrus.WarnLevel, args...)
}

func Warning(args ...interface{}) {
	log().Warn(args...)
}

func Error(args ...interface{}) {
	log().Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	log().Log(logrus.FatalLevel, args...)
	log().Logger.Exit(1)
}

func Panic(args ...interface{}) {
	log().Log(logrus.PanicLevel, args...)
	panic(fmt.Sprint(args...))
}

// Entry Printf family functions
func Logf(level logrus.Level, format string, args ...interface{}) {
	if log().Logger.IsLevelEnabled(level) {
		log().Log(level, fmt.Sprintf(format, args...))
	}
}

func Tracef(format string, args ...interface{}) {
	log().Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	log().Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	log().Logf(logrus.InfoLevel, format, args...)
}

func Printf(format string, args ...interface{}) {
	log().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log().Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	log().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log().Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log().Logf(logrus.FatalLevel, format, args...)
	log().Logger.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	log().Logf(logrus.PanicLevel, format, args...)
}

// Entry Println family functions

func Logln(level logrus.Level, args ...interface{}) {
	if log().Logger.IsLevelEnabled(level) {

		Log(level, sprintlnn(args...))
		// log().Log(level, log().sprintlnn(args...))
	}
}

func Traceln(args ...interface{}) {
	log().Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	log().Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	log().Logln(logrus.InfoLevel, args...)
}

func Println(args ...interface{}) {
	log().Infoln(args...)
}

func Warnln(args ...interface{}) {
	log().Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	log().Warnln(args...)
}

func Errorln(args ...interface{}) {
	log().Logln(logrus.ErrorLevel, args...)
}

func Fatalln(args ...interface{}) {
	log().Logln(logrus.FatalLevel, args...)
	log().Logger.Exit(1)
}

func Panicln(args ...interface{}) {
	log().Logln(logrus.PanicLevel, args...)
}

// Sprintlnn => Sprint no newline. This is to get the behavior of how
// fmt.Sprintln where spaces are always added between operands, regardless of
// their type. Instead of vendoring the Sprintln implementation to spare a
// string allocation, we do the simplest thing.
func sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return GetLog().WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return GetLog().WithField(key, value)
}
