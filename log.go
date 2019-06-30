
// golog is a small wrapper for standard log package with enhanced prefix system.
package golog

import (
	"fmt"
	"log"
)

// RestoreFunc cancels change of prefix
type RestoreFunc func()

// Logger is used for logging
type Logger struct {
	persistentPrefix string
	prefix           string
	*log.Logger
}

// New creates and returns Logger from log.Logger ans sets persistent prefix to the prefix of log.Logger
func New(l *log.Logger) *Logger {
	return &Logger{
		Logger:           l,
		persistentPrefix: l.Prefix(),
	}
}

// PersistentPrefix returns current persistent prefix of the Logger
func (l *Logger) PersistentPrefix() string {
	return l.persistentPrefix
}

// Prefix returns current prefix of the Logger
func (l *Logger) Prefix() string {
	return l.prefix
}

// FullPrefix returns current full prefix of the Logger
func (l *Logger) FullPrefix() string {
	return l.PersistentPrefix() + l.Prefix()
}

// SetPrefix sets prefix fo the Logger
func (l *Logger) SetPrefix(prefix string) RestoreFunc {
	oldPrefix := l.prefix
	l.prefix = prefix
	l.Logger.SetPrefix(l.FullPrefix())
	return func() {
		l.prefix = oldPrefix
		l.Logger.SetPrefix(l.FullPrefix())
	}
}

// SetPrefixf sets prefix of the logger to pattern formatted with current prefix of the Logger by fmt.Sprintf
func (l *Logger) SetPrefixf(pattern string) RestoreFunc {
	return l.SetPrefix(fmt.Sprintf(pattern, l.prefix))
}

// SetPersistentPrefix sets persistent prefix fo the Logger
func (l *Logger) SetPersistentPrefix(prefix string) RestoreFunc {
	oldPersistentPrefix := l.persistentPrefix
	l.persistentPrefix = prefix
	l.Logger.SetPrefix(l.FullPrefix())
	return func() {
		l.persistentPrefix = oldPersistentPrefix
		l.Logger.SetPrefix(l.FullPrefix())
	}
}

// SetPersistentPrefixf sets persistent prefix of the logger to pattern
// formatted with current persistent prefix of the Logger by fmt.Sprintf
func (l *Logger) SetPersistentPrefixf(pattern string) RestoreFunc {
	return l.SetPersistentPrefix(fmt.Sprintf(pattern, l.persistentPrefix))
}
