Logor
===

[![Build Status](https://travis-ci.org/imwithye/logor.svg)](https://travis-ci.org/imwithye/logor)
[![GoDoc](https://godoc.org/github.com/imwithye/logor?status.svg)](http://godoc.org/github.com/imwithye/logor)

Logor is a simple logger for Golang.

## Usage

The default logger is set to `InfoLevel`, so only messages of `Fatal, Error, Warn, Info` will be printed. You can modify the logger's level using `logger.Level = logor.DebugLevel` to enable debug message.
```go
import "github.com/imwithye/logor"

func main() {
  logger := logor.New()
  logger.Error("error")
  logger.Warn("warn")
  logger.Info("info")
  
  logger.Level = logor.DebugLevel
  logger.Debug("debug")
  
  logger.Level = loggor.TraceLevel
  logger.Trace("trace")
}
```

### Shared Logger

If you want to share logger in different files or packages, use `GetLogor()` instead of `New()`
```go
import "github.com/imwithye/logor"

func main() {
  logger := logor.GetLogor()
  logger.Error("error")
  logger.Warn("warn")
  logger.Info("info")
  
  logger.Level = logor.DebugLevel
  logger.Debug("debug")
  
  logger = logor.GetLogor()
  logger.Debug("Message will be printed because the level of the shared logger is Debuglevel")
}
```
