package console

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func log(text ...any) string {

	var parsedText []interface{}
	var format string
	for i, t := range text {
		if i != len(text)-1 {
			format += "[%v] "
			parsedText = append(parsedText, fmt.Sprint(t))
		} else {
			format += "%v"
		}
	}

	parsedText = append(parsedText, func() any {
		if len(text) > 0 {
			return text[len(text)-1]
		} else {
			return ""
		}
	}())

	return fmt.Sprintf(format, parsedText...)

}

func Debug(text ...any) {
	logger.Debug(log(text...))
}

func Info(text ...any) {
	logger.Info(log(text...))
}

func Warn(text ...any) {
	logger.Warn(log(text...))
}

func Error(text ...any) {
	logger.Error(log(text...))
}

type LeveledLogger struct {
	level logrus.Level
}

func (l *LeveledLogger) Debugf(format string, v ...any) {
	if l.level >= logrus.DebugLevel {
		Debug(fmt.Sprintf(format, v...))
	}
}

func (l *LeveledLogger) Infof(format string, v ...any) {
	if l.level >= logrus.InfoLevel {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l *LeveledLogger) Warnf(format string, v ...any) {
	if l.level >= logrus.WarnLevel {
		Warn(fmt.Sprintf(format, v...))
	}
}

func (l *LeveledLogger) Errorf(format string, v ...any) {
	if l.level >= logrus.ErrorLevel {
		Error(fmt.Sprintf(format, v...))
	}
}

func NewLeveledLogger(level logrus.Level) *LeveledLogger {
	return &LeveledLogger{level: level}
}
