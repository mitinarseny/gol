package golog

import (
	"fmt"
	"log"
)

type RestoreFunc func()

type Logger struct {
	persistentPrefix string
	prefix           string
	*log.Logger
}

func New(l *log.Logger) *Logger {
	return &Logger{
		Logger:           l,
		persistentPrefix: l.Prefix(),
	}
}

func (l *Logger) PersistentPrefix() string {
	return l.persistentPrefix
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) FullPrefix() string {
	return l.PersistentPrefix() + l.Prefix()
}

func (l *Logger) SetPrefix(prefix string) RestoreFunc {
	oldPrefix := l.prefix
	l.prefix = prefix
	l.Logger.SetPrefix(l.FullPrefix())
	return func() {
		l.prefix = oldPrefix
		l.Logger.SetPrefix(l.FullPrefix())
	}
}

func (l *Logger) SetPrefixf(pattern string) RestoreFunc {
	return l.SetPrefix(fmt.Sprintf(pattern, l.prefix))
}

func (l *Logger) SetPersistentPrefix(prefix string) RestoreFunc {
	oldPersistentPrefix := l.persistentPrefix
	l.persistentPrefix = prefix
	l.Logger.SetPrefix(l.FullPrefix())
	return func() {
		l.persistentPrefix = oldPersistentPrefix
		l.Logger.SetPrefix(l.FullPrefix())
	}
}

func (l *Logger) SetPersistentPrefixf(pattern string) RestoreFunc {
	return l.SetPersistentPrefix(fmt.Sprintf(pattern, l.persistentPrefix))
}
