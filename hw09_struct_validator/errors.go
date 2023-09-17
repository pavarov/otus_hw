package hw09structvalidator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ErrInvalidRule struct {
	Field string
}

func (e ErrInvalidRule) Error() string {
	return fmt.Sprintf("field `%s` has invalid rule expression", e.Field)
}

type ErrUnknownRule struct {
	Field string
	Rule  string
}

func (e ErrUnknownRule) Error() string {
	return fmt.Sprintf("unknown rule `%s` for field `%s`", e.Rule, e.Field)
}

type ErrInvalidValidatorArgument struct {
	Type reflect.Type
}

func (e ErrInvalidValidatorArgument) Error() string {
	return fmt.Sprintf("invalid argument of the validator: %s. Passed: %s", reflect.Struct.String(), e.Type.String())
}

type ErrUnexpected struct {
	Err error
}

func (e ErrUnexpected) Error() string {
	return fmt.Sprintf("unknown error: %s", e.Err)
}

var (
	ErrUnsupportedType = errors.New("unsupported type for validator")
	ErrInvalidLen      = errors.New("length of the values not equal to expected")
	ErrNotInclude      = errors.New("value should be in expected")
	ErrRegexNotMatched = errors.New("value not matched to expected regexp")
	ErrLenBigger       = errors.New("length of the value is bigger than expected")
	ErrLenLess         = errors.New("length of the value is less than expected")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	buff := bytes.NewBufferString("")

	for i := 0; i < len(v); i++ {
		ve := v[i]
		buff.WriteString(fmt.Sprintf("field: `%s` - error: `%s`", ve.Field, ve.Err.Error()))
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}
