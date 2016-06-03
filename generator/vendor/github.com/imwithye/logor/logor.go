package logor

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logor's log level. Higher number means higher level. Only if the current level
// is greater than the function level, logs will be printed. For example, if log
// level is set to TraceLevel, all logs will be printed; but if log level is set
// to ErrorLevel, only Fatal and Error logs will be printed.
const (
	Off        = 0
	FatalLevel = 100
	ErrorLevel = 200
	WarnLevel  = 300
	InfoLevel  = 400
	DebugLevel = 500
	TraceLevel = 600
)

var logor *Logor

// Logor struct defines the layout of Logor instance. Logor is actual based on
// the default logger of golang. So you are able to change any single logger inside
// Logor
type Logor struct {
	Level       int
	FatalLogger *log.Logger
	ErrorLogger *log.Logger
	WarnLogger  *log.Logger
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	TraceLogger *log.Logger
}

// New creates a logger using Stdout and Stderr
func New() *Logor {
	return NewCustomIO(os.Stdout, os.Stdout)
}

// NewCustomIO creates a logger with custom io
func NewCustomIO(out io.Writer, err io.Writer) *Logor {
	l := new(Logor)
	debugFlag := log.Lshortfile | log.LstdFlags

	l.Level = InfoLevel
	l.FatalLogger = log.New(err, "[F] ", log.LstdFlags)
	l.ErrorLogger = log.New(err, "[E] ", log.LstdFlags)
	l.WarnLogger = log.New(out, "[W] ", log.LstdFlags)
	l.InfoLogger = log.New(out, "[I] ", log.LstdFlags)
	l.DebugLogger = log.New(out, "[D] ", debugFlag)
	l.TraceLogger = log.New(out, "[T] ", debugFlag)

	return l
}

// GetLogor gets a shared logger using Stdout and Stderr
func GetLogor() *Logor {
	if logor == nil {
		logor = New()
	}
	return logor
}

// GetLogorCustomIO get a shared logger using custom io
func GetLogorCustomIO(out io.Writer, err io.Writer) *Logor {
	if logor == nil {
		logor = NewCustomIO(out, err)
	}
	return logor
}

// SetFlags sets the output flags for all loggers. Flags supported are list in
// http://golang.org/pkg/log/#pkg-constants
func (l *Logor) SetFlags(flag int) {
	l.FatalLogger.SetFlags(flag)
	l.ErrorLogger.SetFlags(flag)
	l.WarnLogger.SetFlags(flag)
	l.InfoLogger.SetFlags(flag)
	l.DebugLogger.SetFlags(flag)
	l.TraceLogger.SetFlags(flag)
}

// Fatal logs the error message to Stderr and exit program with code 1.
func (l *Logor) Fatal(v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintln(v...))
		os.Exit(1)
	}
}

// Fatalf logs the error message to Stderr and exit program with code 1.
// Message is formated by the format string.
func (l *Logor) Fatalf(format string, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintf(format, v...))
		os.Exit(1)
	}
}

// FatalCode logs the error message to Stderr and exit program with a custom code.
func (l *Logor) FatalCode(code int, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintln(v...))
		os.Exit(code)
	}
}

// FatalfCode logs the error message to Stderr and exit program with a custom code.
// Message is formated by the format string.
func (l *Logor) FatalfCode(code int, format string, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintf(format, v...))
		os.Exit(code)
	}
}

// Error logs the error message to Stderr.
func (l *Logor) Error(v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.ErrorLogger.Output(2, fmt.Sprintln(v...))
	}
}

// Errorf logs the error message to Stderr.
// Message is formated by the format string.
func (l *Logor) Errorf(format string, v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.ErrorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Warn logs the warn message to Stdout.
func (l *Logor) Warn(v ...interface{}) {
	if l.Level >= WarnLevel {
		l.WarnLogger.Output(2, fmt.Sprintln(v...))
	}
}

// Warnf logs the warn message to Stdout.
// Message is formated by the format string.
func (l *Logor) Warnf(format string, v ...interface{}) {
	if l.Level >= WarnLevel {
		l.WarnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Info logs the info message to Stdout.
func (l *Logor) Info(v ...interface{}) {
	if l.Level >= InfoLevel {
		l.InfoLogger.Output(2, fmt.Sprintln(v...))
	}
}

// Infof logs the info message to Stdout.
// Message is formated by the format string.
func (l *Logor) Infof(format string, v ...interface{}) {
	if l.Level >= InfoLevel {
		l.InfoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Debug logs the debug message to Stdout.
func (l *Logor) Debug(v ...interface{}) {
	if l.Level >= DebugLevel {
		l.DebugLogger.Output(2, fmt.Sprintln(v...))
	}
}

// Debugf logs the debug message to Stdout.
// Message is formated by the format string.
func (l *Logor) Debugf(format string, v ...interface{}) {
	if l.Level >= DebugLevel {
		l.DebugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Trace logs the trace message to Stdout.
func (l *Logor) Trace(v ...interface{}) {
	if l.Level >= TraceLevel {
		l.TraceLogger.Output(2, fmt.Sprintln(v...))
	}
}

// Tracef logs the trace message to Stdout.
// Message is formated by the format string.
func (l *Logor) Tracef(format string, v ...interface{}) {
	if l.Level >= TraceLevel {
		l.TraceLogger.Output(2, fmt.Sprintf(format, v...))
	}
}
