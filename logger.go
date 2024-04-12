package tglogger

import (
	"context"

	"github.com/wmuga/tglogger/internal/models"
	tgl "github.com/wmuga/tglogger/internal/tglogger"
)

// Int creates field with integer value
func Int(key string, value int) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldInt}
}

// Uint creates field with unsigned integer value
func Uint(key string, value uint) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldUint}
}

// Int64 creates field with integer value 64bit
func Int64(key string, value int64) models.Field {
	return models.Field{Key: key, Int: value, Type: models.FieldInt}
}

// Uint64 creates field with unsigned integer value 64bit
func Uint64(key string, value uint64) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldUint}
}

// Int32 creates field with integer value 32bit
func Int32(key string, value int32) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldInt}
}

// Uint32 creates field with unsigned integer value 32bit
func Uint32(key string, value uint32) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldUint}
}

// Int16 creates field with integer value 16bit
func Int16(key string, value int16) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldInt}
}

// Uint16 creates field with unsigned integer value 16bit
func Uint16(key string, value uint16) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldUint}
}

// Int8 creates field with integer value 8bit
func Int8(key string, value int8) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldInt}
}

// Uint8 creates field with unsigned integer value 8bit
func Uint8(key string, value uint8) models.Field {
	return models.Field{Key: key, Int: int64(value), Type: models.FieldUint}
}

// String creates field with string value
func String(key string, value string) models.Field {
	return models.Field{Key: key, Str: value, Type: models.FieldString}
}

// Float32 creates field with floating point value 32bit
func Float32(key string, value float32) models.Field {
	return models.Field{Key: key, Float: float64(value), Type: models.FieldFloat}
}

// Float64 creates field with floating point value 64bit
func Float64(key string, value float64) models.Field {
	return models.Field{Key: key, Float: value, Type: models.FieldFloat}
}

// Error creates field with error point value
func Error(value error) models.Field {
	return models.Field{Key: "error", Str: value.Error(), Type: models.FieldError}
}

// Any creates field with interface{} value
func Any(key string, value interface{}) models.Field {
	return models.Field{Key: key, Interface: value, Type: models.FieldInterface}
}

func New(ctx context.Context, topic, token string, chatIDs ...int64) (models.TgLogger, error) {
	return tgl.NewTgLogger(ctx, topic, token, chatIDs...)
}
