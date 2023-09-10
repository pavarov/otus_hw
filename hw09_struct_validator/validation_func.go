package hw09structvalidator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Func = func(v reflect.Value, exp string) (validationErr, unknownErr error)

func regex(v reflect.Value, reg string) (validationErr, unknownErr error) {
	matchString, err := regexp.MatchString(reg, v.String())
	if err != nil {
		return nil, err
	}
	if !matchString {
		return ErrRegexNotMatched, nil
	}
	return nil, nil
}

func lenOf(v reflect.Value, lv string) (validationErr, unknownErr error) {
	lenVal, err := strconv.Atoi(lv)
	if err != nil {
		return nil, ErrUnexpected{err}
	}

	switch {
	case reflect.String == v.Kind():
		if len(v.String()) != lenVal {
			return ErrInvalidLen, nil
		}
	case reflect.Slice == v.Kind() && v.Type().Elem().Kind() == reflect.String:
		for _, val := range v.Interface().([]string) {
			if len(val) != lenVal {
				return ErrInvalidLen, nil
			}
		}
	default:
		return nil, ErrInvalidValidatorArgument{v.Type()}
	}

	return nil, nil
}

func include(v reflect.Value, val string) (validationErr, unknownErr error) {
	values := strings.Split(val, ",")
	for _, value := range values {
		//nolint:exhaustive
		switch v.Kind() {
		case reflect.Int:
			heapVal, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			if heapVal == v.Interface().(int) {
				return nil, nil
			}
		default:
			if v.String() == value {
				return nil, nil
			}
		}
	}
	return ErrNotInclude, nil
}

func min(v reflect.Value, val string) (validationErr, unknownErr error) {
	minVal, err := strconv.Atoi(val)
	if err != nil {
		return nil, ErrUnexpected{err}
	}
	if v.Kind() != reflect.Int {
		return nil, ErrInvalidValidatorArgument{Type: v.Type()}
	}
	if v.Interface().(int) < minVal {
		return ErrLenLess, nil
	}
	return nil, nil
}

func max(v reflect.Value, val string) (validationErr, unknownErr error) {
	maxVal, err := strconv.Atoi(val)
	if err != nil {
		return nil, ErrUnexpected{err}
	}
	if v.Kind() != reflect.Int {
		return nil, ErrInvalidValidatorArgument{Type: v.Type()}
	}
	if v.Interface().(int) > maxVal {
		return ErrLenBigger, nil
	}
	return nil, nil
}

var validatorRules = map[string]Func{
	"len":    lenOf,
	"regexp": regex,
	"in":     include,
	"min":    min,
	"max":    max,
}
