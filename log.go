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
	return &Logger{l}
}

func (l *Logger) Append(prefix string) UnLevelFunc {
	oldPrefix := l.Prefix()
	l.SetPrefix(oldPrefix + prefix)
	return func() {
		l.SetPrefix(oldPrefix)
	}
}

func (l *Logger) Replace(prefix string) UnLevelFunc {
	oldPrefix := l.Prefix()
	l.SetPrefix(prefix)
	return func() {
		l.SetPrefix(oldPrefix)
	}
}

func (l *Logger) Format(pattern string) UnLevelFunc {
	oldPrefix := l.Prefix()
	l.SetPrefix(fmt.Sprintf(pattern, oldPrefix))
	return func() {
		l.SetPrefix(oldPrefix)
	}
}
