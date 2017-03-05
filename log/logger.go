package log

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

// Logger ...
type Logger struct {
	silent        bool
	appName       string
	entry         *logrus.Entry
	levelFuncMap  map[Level]func(...interface{})
	levelFuncfMap map[Level]func(string, ...interface{})
}

// Option ...
type Option func(*Logger) error

// WithSilent ...
func WithSilent(silent bool) Option {
	return func(log *Logger) error {
		log.silent = silent
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

// NewLogger ...
func NewLogger(appName string, options ...Option) (*Logger, error) {
	log := &Logger{
		silent:  false,
		appName: appName,
	}
	log.levelFuncMap = map[Level]func(...interface{}){
		D: func(args ...interface{}) { log.entry.Debug(args...) },
		I: func(args ...interface{}) { log.entry.Info(args...) },
		W: func(args ...interface{}) { log.entry.Warn(args...) },
		E: func(args ...interface{}) { log.entry.Error(args...) },
		F: func(args ...interface{}) { log.entry.Fatal(args...) },
		P: func(args ...interface{}) { log.entry.Panic(args...) },
	}
	log.levelFuncfMap = map[Level]func(string, ...interface{}){
		D: func(format string, args ...interface{}) { log.entry.Debugf(format, args...) },
		I: func(format string, args ...interface{}) { log.entry.Infof(format, args...) },
		W: func(format string, args ...interface{}) { log.entry.Warnf(format, args...) },
		E: func(format string, args ...interface{}) { log.entry.Errorf(format, args...) },
		F: func(format string, args ...interface{}) { log.entry.Fatalf(format, args...) },
		P: func(format string, args ...interface{}) { log.entry.Panicf(format, args...) },
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})

	for _, option := range options {
		err := option(log)
		if err != nil {
			return nil, err
		}
	}

	return log, nil
}

// WithField ...
func (l *Logger) WithField(key string, field interface{}) *Logger {
	if l == nil {
		return nil
	}
	l.initEntry()
	l.entry = l.entry.WithField(key, field)
	return l
}

// WithFields ...
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	if l == nil {
		return nil
	}
	l.initEntry()
	l.entry = l.entry.WithFields(fields)
	return l
}

// Log ...
func (l *Logger) Log(level Level, args ...interface{}) {
	if l == nil {
		return
	}
	l.initEntry()
	l.levelFuncMap[level](args...)
	l.entry = nil
}

// Logf ...
func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	if l == nil {
		return
	}
	l.initEntry()
	l.levelFuncfMap[level](format, args...)
	l.entry = nil
}

func (l *Logger) initEntry() {
	if l.entry == nil {
		l.entry = logrus.WithField("appName", l.appName)
	}
}
