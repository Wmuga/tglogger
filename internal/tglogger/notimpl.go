package tglogger

import (
	"errors"
	"fmt"
	"os"

	"github.com/wmuga/tglogger/internal/models"
)

type notImplementedLogger struct{}

var (
	ErrNotImplemented = errors.New("not implemented")
)

// AddChat implements tglogger.TgLogger.
func (*notImplementedLogger) AddChat(chatID int64) error {
	return ErrNotImplemented
}

// Error implements tglogger.TgLogger.
func (*notImplementedLogger) Error(message string, fields ...models.Field) error {
	return ErrNotImplemented
}

// Fatal implements tglogger.TgLogger.
func (*notImplementedLogger) Fatal(message string, fields ...models.Field) {
	fmt.Println(ErrNotImplemented.Error())
	os.Exit(1)
}

// Info implements tglogger.TgLogger.
func (*notImplementedLogger) Info(message string, fields ...models.Field) error {
	return ErrNotImplemented
}

// Print implements tglogger.TgLogger.
func (*notImplementedLogger) Print(level string, message string, fields ...models.Field) error {
	return ErrNotImplemented
}

// RemoveChat implements tglogger.TgLogger.
func (*notImplementedLogger) RemoveChat(chatID int64) error {
	return ErrNotImplemented
}
