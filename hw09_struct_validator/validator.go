package hw09structvalidator

import (
	"errors"
	"reflect"
	"strings"
)

const (
	orSep      = "|"
	funcValSep = ":"
)

func valueProcessing(value reflect.Value) error {
	ve := make(ValidationErrors, 0)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if !field.IsExported() {
			continue
		}
		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		rules := strings.Split(validateTag, orSep)
		if len(rules) == 0 {
			continue
		}
		for _, rule := range rules {
			var validationErr, err error

			ruleData := strings.Split(rule, funcValSep)
			if len(ruleData) != 2 || ruleData[0] == "" || ruleData[1] == "" {
				return ErrInvalidRule{Field: field.Name}
			}

			f, ok := validatorRules[ruleData[0]]
			if !ok {
				return ErrUnknownRule{Field: field.Name, Rule: ruleData[0]}
			}

			validationErr, err = f(value.Field(i), ruleData[1])

			if err != nil {
				return err
			}

			if validationErr != nil {
				ve = append(ve, ValidationError{Field: field.Name, Err: validationErr})
			}
		}
	}

	return ve
}

func Validate(v interface{}) error {
	value := reflect.ValueOf(v)

	if value.Kind() == reflect.Pointer && !value.IsNil() {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return ErrUnsupportedType
	}
	err := valueProcessing(value)

	var ve ValidationErrors
	if !errors.As(err, &ve) {
		return err
	}
	if errors.As(err, &ve) && len(ve) > 0 {
		return ve
	}

	return nil
}
