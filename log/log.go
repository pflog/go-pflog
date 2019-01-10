package log

import (
	"os"

	"github.com/mrcrgl/pflog/pkg/logging"

	"github.com/mrcrgl/pflog/container"

	"github.com/mrcrgl/pflog"
	"github.com/mrcrgl/pflog/pkg/encoding/text"
	"github.com/mrcrgl/pflog/pkg/logger"
)

func init() {
	pflog.Register.Register(
		new(container.Ctx),
		new(container.Error),
		new(container.Message),
		new(container.Text),
	)
}

var defaultLogger = logger.New(text.NewEncoder(), os.Stderr)

func With(containers ...logging.Container) logging.Logger {
	return defaultLogger.With(containers...)
}

func V(level int) logging.InfoLogger {
	return defaultLogger.V(level)
}

func Info(s string) {
	defaultLogger.Info(s)
}

func Infof(s string, args ...interface{}) {
	defaultLogger.Infof(s, args...)
}

func Warning(s string) {
	defaultLogger.Warning(s)
}

func Warningf(s string, args ...interface{}) {
	defaultLogger.Warningf(s, args...)
}

func Error(s string) {
	defaultLogger.Error(s)
}

func Errorf(s string, args ...interface{}) {
	defaultLogger.Errorf(s, args...)
}

func Fatal(s string) {
	defaultLogger.Fatal(s)
}

func Fatalf(s string, args ...interface{}) {
	defaultLogger.Fatalf(s, args...)
}
