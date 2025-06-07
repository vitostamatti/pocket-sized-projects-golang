package pocketlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns logger ready to log a the required threshold.
// The default output is Stdout
func New(threshold Level, opts ...Option) *Logger {
	l := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	l.logf(LevelDebug, format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	l.logf(LevelInfo, format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	l.logf(LevelError, format, args...)
}

// logf prints the message to the output.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	contents := fmt.Sprintf(format, args...)

	msg := message{
		Level:   lvl.String(),
		Message: contents,
	}

	// encode the message
	formattedMessage, err := json.Marshal(msg)
	if err != nil {
		_, _ = fmt.Fprintf(l.output, "unable to format message for %v\n", contents)
		return
	}

	_, _ = fmt.Fprintln(l.output, string(formattedMessage))
}

// message represents the JSON structure of the logged messages.
type message struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// Logf formats and prints a message if the log level is high enough
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args...)
}
