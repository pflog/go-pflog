package logger

import "github.com/mrcrgl/pflog/pkg/logging"

var _ logging.Logger = &noopLogger{}

type noopLogger struct{}

func (l *noopLogger) With(_ ...logging.Container) logging.Logger {
	return l
}

func (_ *noopLogger) Info(string)                    {}
func (_ noopLogger) Infof(string, ...interface{})    {}
func (_ noopLogger) Warning(string)                  {}
func (_ noopLogger) Warningf(string, ...interface{}) {}
func (_ noopLogger) Error(string)                    {}
func (_ noopLogger) Errorf(string, ...interface{})   {}
func (_ noopLogger) Fatal(string)                    {}
func (_ noopLogger) Fatalf(string, ...interface{})   {}
