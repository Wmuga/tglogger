package tglogger

import "github.com/wmuga/tglogger/internal/models"

type notImplementedLogger struct{}

// AddChat implements tglogger.TgLogger.
func (*notImplementedLogger) AddChat(chatID int64) {
	panic("unimplemented")
}

// Error implements tglogger.TgLogger.
func (*notImplementedLogger) Error(message string, fields ...models.Field) {
	panic("unimplemented")
}

// Fatal implements tglogger.TgLogger.
func (*notImplementedLogger) Fatal(message string, fields ...models.Field) {
	panic("unimplemented")
}

// Info implements tglogger.TgLogger.
func (*notImplementedLogger) Info(message string, fields ...models.Field) {
	panic("unimplemented")
}

// Print implements tglogger.TgLogger.
func (*notImplementedLogger) Print(level string, message string, fields ...models.Field) {
	panic("unimplemented")
}

// RemoveChat implements tglogger.TgLogger.
func (*notImplementedLogger) RemoveChat(chatID int64) {
	panic("unimplemented")
}
