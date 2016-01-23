package log

import (
	"fmt"
	"log"
)

const (
	debug   = "DEBUG"
	info    = "INFO"
	warning = "WARNING"
)

// Meta (data) for a potential series of log messages.
type Meta struct {
	User   string
	Module string
	Level  string
}

// Printf formats the input and writes a log entry.
func (m *Meta) Printf(format string, a ...interface{}) {
	log.Printf("[%s] %s: [%s] %s", m.Level, m.Module, m.User, fmt.Sprintf(format, a...))
}

// Print a log entry.
func (m *Meta) Print(a ...interface{}) {
	log.Printf("[%s] %s: [%s] %s", m.Level, m.Module, m.User, fmt.Sprint(a...))
}

// Debug level message.
func (m *Meta) Debug() *Meta {
	m.Level = debug
	return m
}

// Info (rmational) level message.
func (m *Meta) Info() *Meta {
	m.Level = info
	return m
}

// Warning level message.
func (m *Meta) Warning() *Meta {
	m.Level = warning
	return m
}

// For assigns meta data for log entries.
func For(module, user string) *Meta {
	return &Meta{
		User:   user,
		Module: module,
		Level:  debug,
	}
}
