package logger_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/Concatena01/logger"
)

func TestErrorFStdout(t *testing.T) {
	myLogger := logger.NewLogger(&logger.Config{Name: "TestLogger", Output: "stdout"})
	message := myLogger.ErrorF("Test error message")
	if !strings.Contains(message, "TestLogger ERROR: Test error message") {
		t.Errorf("ErrorF did not return expected message, got: %s", message)
	}
}

func TestErrorFFile(t *testing.T) {
	logFile := "test.log"
	myLogger := logger.NewLogger(&logger.Config{Name: "TestLogger", Output: logFile})
	myLogger.ErrorF("Test error message")

	// Read the log file
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Check the log message
	if !strings.Contains(string(content), "TestLogger ERROR: Test error message") {
		t.Errorf("ErrorF did not log expected message, got: %s", string(content))
	}
}

func TestFatalStdout(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		myLogger := logger.NewLogger(&logger.Config{Name: "TestLogger", Output: "stdout"})
		myLogger.Fatal("Test fatal error")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatalStdout")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Process ran with err %v, want exit status 1", err)
}

func TestFatalFile(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		logFile := "test_fatal.log"
		myLogger := logger.NewLogger(&logger.Config{Name: "TestLogger", Output: logFile})
		myLogger.Fatal("Test fatal error")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatalFile")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Process ran with err %v, want exit status 1", err)

	// Read the log file
	content, err := os.ReadFile("test_fatal.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Check the log message
	if !strings.Contains(string(content), "TestLogger FATAL: Test fatal error") {
		t.Errorf("Fatal did not log expected message, got: %s", string(content))
	}
}

func TestInfoStdout(t *testing.T) {
	myLogger := logger.NewLogger(&logger.Config{Name: "TestLogger", Output: "stdout"})
	message := myLogger.Info("Test info message")
	if !strings.Contains(message, "TestLogger INFO: Test info message") {
		t.Errorf("Info did not return expected message, got: %s", message)
	}
}

func TestInfoFile(t *testing.T) {
	logFile := "test_info.log"
	myLogger := logger.NewLogger(&logger.Config{Name: "TestLogger", Output: logFile})
	myLogger.Info("Test info message")

	// Read the log file
	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Check the log message
	if !strings.Contains(string(content), "TestLogger INFO: Test info message") {
		t.Errorf("Info did not log expected message, got: %s", string(content))
	}
}
