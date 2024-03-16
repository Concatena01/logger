package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Severity string

const (
	SeverityInfo  Severity = "INFO"
	SeverityError Severity = "ERROR"
	SeverityFatal Severity = "FATAL"
)

func defaultConfig() *Config {
	return &Config{Output: "stdout", Name: "LOGGER"}
}

type Logger struct {
	cnf *Config
}

type Config struct {
	Name   string
	Output string // stdout, <filepath>
}

func NewLogger(config *Config) *Logger {
	var logger = new(Logger)
	if config != nil {
		if config.Output == "" {
			config.Output = defaultConfig().Output
		}
		if config.Name == "" {
			config.Name = defaultConfig().Name
		}
	} else {
		config = defaultConfig()
	}

	logger.cnf = config

	if logger.cnf.Output != "stdout" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logger.cnf.Output,
			MaxSize:    5, // megabytes
			MaxBackups: 100,
			MaxAge:     30, // days
		})

		log.SetFlags(0)
	}

	return logger
}

func (l Logger) LogF(level Severity, args ...any) string {
	message := fmt.Sprintf(
		"%s %s %s: %s",
		time.Now().UTC().Format(time.RFC3339Nano),
		l.cnf.Name, level, fmt.Sprint(args...),
	)

	switch l.cnf.Output {
	case "stdout":
		fmt.Println(message)
	default:
		log.Println(message)
	}

	return message
}

func (l Logger) ErrorF(args ...any) string {
	return l.LogF(SeverityError, args...)
}

func (l Logger) Fatal(args ...any) {
	l.LogF(SeverityFatal, args...)
	os.Exit(1)
}

func (l Logger) Info(args ...any) string {
	return l.LogF(SeverityInfo, args...)
}
