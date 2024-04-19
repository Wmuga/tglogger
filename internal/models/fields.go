package models

import (
	"fmt"
	"strconv"
)

type FieldType int

type Field struct {
	Key  string
	Type FieldType

	Int       int64
	Str       string
	Float     float64
	Interface interface{}
}

const (
	Unknown FieldType = iota

	FieldInt
	FieldUint
	FieldFloat
	FieldString
	FieldError
	FieldDateTime
	FieldInterface
)

func (f Field) String() string {
	return fmt.Sprintf("%s:%s", f.Key, f.string())
}

func (f Field) string() string {
	switch f.Type {
	case FieldInt:
		return strconv.FormatInt(f.Int, 10)
	case FieldUint:
		return strconv.FormatUint(uint64(f.Int), 10)
	case FieldFloat:
		return strconv.FormatFloat(f.Float, 'g', -1, 64)
	case FieldString:
		return f.Str
	case FieldError:
		return f.Str
	case FieldInterface:
		return interfaceToStr(f.Interface)
	case FieldDateTime:
		return f.Str
	default:
		return ""
	}
}

func interfaceToStr(v interface{}) string {
	return fmt.Sprintf("%+v", v)
}
