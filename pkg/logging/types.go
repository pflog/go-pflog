package logging

import "io"

type Logger interface {
	ContainerLogger
	FatalLogger
	ErrorLogger
	WarningLogger
	InfoLogger
}

type ContainerLogger interface {
	With(containers ...Container) Logger
}

type Silencer interface {
	V(int) InfoLogger
}

type InfoLogger interface {
	Info(string)
	Infof(string, ...interface{})
}

type WarningLogger interface {
	Warning(string)
	Warningf(string, ...interface{})
}

type ErrorLogger interface {
	Error(string)
	Errorf(string, ...interface{})
}

type FatalLogger interface {
	Fatal(string)
	Fatalf(string, ...interface{})
}

type Container interface {
	Enclosed() bool
	Kind() []byte

	TextConverter
}

type TextConverter interface {
	WriteTextTo(writer io.Writer) (int, error)
	ReadTextFrom(reader io.Reader) (int, error)
}
