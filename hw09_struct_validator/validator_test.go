package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	HasUnknownRule struct {
		F int `validate:"some_rule:v_a_l|in:123"`
	}

	HasInvalidRuleExpression struct {
		F string `validate:"sr"`
	}

	HasInvalidRuleExpression2 struct {
		F string `validate:":"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     fmt.Sprintf("%36s", "A"),
				Age:    18,
				Email:  "email@email.ru",
				Role:   "admin",
				Phones: []string{"12345678901", "12345678901"},
				meta:   nil,
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "123",
				Name:   "name",
				Age:    -1,
				Email:  "emailemail.ru",
				Role:   "some role",
				Phones: []string{"12345678901", "123456789012"},
				meta:   json.RawMessage{},
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: ErrInvalidLen},
				ValidationError{Field: "Age", Err: ErrLenLess},
				ValidationError{Field: "Email", Err: ErrRegexNotMatched},
				ValidationError{Field: "Role", Err: ErrNotInclude},
				ValidationError{Field: "Phones", Err: ErrInvalidLen},
			},
		},
		{
			in:          "string",
			expectedErr: ErrUnsupportedType,
		},
		{
			in: App{Version: "version"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Version", Err: ErrInvalidLen},
			},
		},
		{
			in: Token{
				Header:    []byte{'v', 'a'},
				Payload:   nil,
				Signature: []byte{1},
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				Body: "",
			},
			expectedErr: nil,
		},
		{
			in: HasUnknownRule{F: 123},
			expectedErr: ErrUnknownRule{
				Field: "F",
				Rule:  "some_rule",
			},
		},
		{
			in:          HasInvalidRuleExpression{F: "some_str"},
			expectedErr: ErrInvalidRule{Field: "F"},
		},
		{
			in:          HasInvalidRuleExpression2{F: "some_str"},
			expectedErr: ErrInvalidRule{Field: "F"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			if tt.expectedErr == nil {
				require.ErrorIs(t, tt.expectedErr, err)
			} else {
				require.EqualError(t, tt.expectedErr, err.Error())
			}
		})
	}
}
