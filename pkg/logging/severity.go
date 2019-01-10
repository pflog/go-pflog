package logging

type Severity byte

const (
	SeverityUndefined Severity = iota
	SeverityInfo
	SeverityWarning
	SeverityError
	SeverityFatal
)
