package log

import (
	"fmt"
	"log"
	"os"
)

type Stdout struct {
	logger *log.Logger
}

func NewStdout() Stdout {
	return Stdout{log.New(os.Stdout, "", log.LstdFlags)}
}

func (std Stdout) Info(msg string, args ...interface{}) {
	if len(args) >= 1 {
		msg = fmt.Sprintf(msg, args)
	}
	std.logger.Printf("[INFO] %s", msg)
}

func (std Stdout) Warn(msg string, args ...interface{}) {
	if len(args) >= 1 {
		msg = fmt.Sprintf(msg, args)
	}
	std.logger.Printf("[WARN] %s", msg)
}

func (std Stdout) Error(msg string, args ...interface{}) {
	if len(args) >= 1 {
		msg = fmt.Sprintf(msg, args)
	}
	std.logger.Printf("[ERROR] %s", msg)
}

func (std Stdout) Critical(msg string, args ...interface{}) {
	if len(args) >= 1 {
		msg = fmt.Sprintf(msg, args)
	}
	std.logger.Printf("[CRITICAL] %s", msg)
}
