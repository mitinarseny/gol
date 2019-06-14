package golog

import (
	"fmt"
	"log"
)

type UnLevelFunc func()

type Logger struct {
	*log.Logger
}

func New(l *log.Logger) *Logger {
	return &Logger{
		Logger: l,
	}
}

func (l *Logger) SetPrefix(prefix string) UnLevelFunc {
	oldPrefix := l.Prefix()
	l.Logger.SetPrefix(prefix)
	return func() {
		l.Logger.SetPrefix(oldPrefix)
	}
}

func (l *Logger) SetPrefixf(pattern string) UnLevelFunc {
	return l.SetPrefix(fmt.Sprintf(pattern, l.Prefix()))
}
