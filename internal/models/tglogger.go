package models

type TgLogger interface {
	// AddChat adds chat with given id to log info
	AddChat(chatID int64) error
	// RemoveChat removes chat from logger
	RemoveChat(chatID int64) error
	// Print sends message to connected chats
	Print(level string, message string, fields ...Field) error
	// Info sends message to connected chats with level "info"
	Info(message string, fields ...Field) error
	// Error sends message to connected chats with level "error"
	Error(message string, fields ...Field) error
	// Print sends message to connected chats with level "error" and exits with code 1
	Fatal(message string, fields ...Field)
}
