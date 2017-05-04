package log3

import (
	"io"

	"github.com/Sirupsen/logrus"
)

// Level ...
type Level int

const (
	// D ... Debug
	D Level = iota + 1
	// I ... Info
	I
	// W ... Warn
	W
	// E ... Error
	E
	// F ... Fatal
	F
	// P ... Panic
	P
)

type Logger struct {
	isSilent bool
	appID    string
}

// Option ...
type Option func(*Logger) error

// WithSilent ...
func WithSilent(isSilent bool) Option {
	return func(log *Logger) error {
		log.isSilent = isSilent
		return nil
	}
}

// WithLevel
func WithLevel(level string) Option {
	return func(log *Logger) error {
		lvl, err := logrus.ParseLevel(level)
		if err != nil {
			return err
		}
		logrus.SetLevel(lvl)
		return nil
	}
}

// WithOutput
func WithOutput(out io.Writer) Option {
	return func(log *Logger) error {
		logrus.SetOutput(out)
		return nil
	}
}
