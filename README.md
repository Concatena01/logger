# Logger Module

This module provides an interface for logging messages with different severity levels: INFO, ERROR, and FATAL. Messages can be logged to standard output or to a file.

## Usage

First, create a new instance of Logger:

```go
config := &logger.Config{Name: "MyLogger", Output: "mylog.log"}
myLogger := logger.NewLogger(config)
```

You can log messages with different severity levels:
```
myLogger.Info("This is an info message")
myLogger.ErrorF("This is an error message")
myLogger.Fatal("This is a fatal error message")
```

## Configuration
Logger configuration is done through the Config struct:
```
type Config struct {
    Name   string // The name of the logger
    Output string // "stdout" or the path to a log file
}
```