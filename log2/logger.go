package log2

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

type Logger struct {
	attr map[string]interface{}
	mux  *sync.RWMutex
}

func New() *Logger {
	return &Logger{
		attr: make(map[string]interface{}),
		mux:  new(sync.RWMutex),
	}
}
func WithFuncName(l *Logger, funcName string) *Logger {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.attr["func_name"] = funcName
	return l
}

func WithUserID(l *Logger, userID string) *Logger {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.attr["user_id"] = userID
	return l
}

func Info(l *Logger, args ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	logrus.WithFields(l.attr).Info(args...)
}
