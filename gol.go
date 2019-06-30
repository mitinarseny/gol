// gol is a small wrapper for standard log package with enhanced prefix system.
package gol

import (
	"fmt"
	"log"
	"os"
)

// RestoreFunc cancels change of prefix
type RestoreFunc func()

// Logger is used for logging
type Logger struct {
	*log.Logger
	persistentPrefix string
	prefix           string
}

// New creates and returns Logger from log.Logger ans sets persistent prefix to the prefix of log.Logger
func New(l *log.Logger) *Logger {
	return &Logger{
		Logger:           l,
		prefix:           "",
		persistentPrefix: l.Prefix(),
	}
}

var std = New(log.New(os.Stdout, "", log.LstdFlags))

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

// PersistentPrefix returns current persistent prefix of the standard logger
func PersistentPrefix() string {
	return std.persistentPrefix
}

// Prefix returns current prefix of the standard logger
func Prefix() string {
	return std.Prefix()
}

// FullPrefix returns current full prefix of the standard logger
func FullPrefix() string {
	return std.FullPrefix()
}

// SetPrefix sets prefix fo the standard logger
func SetPrefix(prefix string) RestoreFunc {
	return std.SetPrefix(prefix)
}

// SetPrefixf sets prefix of the logger to pattern formatted with current prefix of the standard logger by fmt.Sprintf
func SetPrefixf(pattern string) RestoreFunc {
	return std.SetPrefixf(pattern)
}

// SetPersistentPrefix sets persistent prefix fo the standard logger
func SetPersistentPrefix(prefix string) RestoreFunc {
	return std.SetPersistentPrefix(prefix)
}

// SetPersistentPrefixf sets persistent prefix of the standard logger to pattern
// formatted with current persistent prefix of the Logger by fmt.Sprintf
func SetPersistentPrefixf(prefix string) RestoreFunc {
	return std.SetPersistentPrefixf(prefix)
}
