package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mrcrgl/pflog/pkg/logging"

	"github.com/mrcrgl/pflog/container"
)

func New(encoder logging.Encoder, output io.Writer) *logger {
	return &logger{
		encoder: encoder,
		output:  output,
	}
}

var _ logging.Logger = &logger{}
var noop = new(noopLogger)

type logger struct {
	level      int
	encoder    logging.Encoder
	output     io.Writer
	containers []logging.Container
}

func (l *logger) V(level int) logging.InfoLogger {
	if l.level >= level {
		return l
	}

	return noop
}

func (l *logger) With(containers ...logging.Container) logging.Logger {
	// TODO alloc
	return &logger{
		level:      l.level,
		encoder:    l.encoder,
		output:     l.output,
		containers: append(l.containers, containers...),
	}
}

func (l *logger) Info(s string) {
	l.logf(logging.SeverityInfo, l.containers, s)
}

func (l *logger) Infof(s string, args ...interface{}) {
	l.logf(logging.SeverityInfo, l.containers, s, args...)
}

func (l *logger) Warning(s string) {
	l.logf(logging.SeverityWarning, l.containers, s)
}

func (l *logger) Warningf(s string, args ...interface{}) {
	l.logf(logging.SeverityWarning, l.containers, s, args...)
}

func (l *logger) Error(s string) {
	l.logf(logging.SeverityError, l.containers, s)
}

func (l *logger) Errorf(s string, args ...interface{}) {
	l.logf(logging.SeverityError, l.containers, s, args...)
}

func (l *logger) Fatal(s string) {
	l.logf(logging.SeverityFatal, l.containers, s)
}

func (l *logger) Fatalf(s string, args ...interface{}) {
	l.logf(logging.SeverityFatal, l.containers, s, args...)
}

func (_ logger) convertToMessageContainer(format string, args ...interface{}) logging.Container {
	if len(args) == 0 {
		return container.NewMessage([]byte(format))
	}

	// TODO alloc
	b := new(bytes.Buffer)
	_, _ = fmt.Fprintf(b, format, args...)

	return container.NewMessage(b.Bytes())
}

func (l *logger) logf(severity logging.Severity, containers []logging.Container, format string, args ...interface{}) {
	c := l.convertToMessageContainer(format, args...)
	if c != nil {
		containers = append(containers, c)
	}

	l.log(severity, containers)
}

func (l *logger) log(severity logging.Severity, containers []logging.Container) {
	// TODO alloc
	e := new(logging.Entry)
	e.Severity = severity
	e.Timestamp = time.Now()
	e.Containers = containers

	l.write(e)
}

func (l *logger) write(entry *logging.Entry) {
	// TODO alloc
	b, err := l.encoder.Encode(entry)
	if err != nil {
		fmt.Printf("Log encodimg error: %s\n", err.Error())
		return
	}

	if _, err := l.output.Write(b); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[Logger Error] Write of log stream failed: %v\n", err)
	}
}
