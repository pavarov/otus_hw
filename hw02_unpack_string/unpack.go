package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func validator(s string) (bool, error) {
	if len(s) == 0 {
		return false, nil
	}
	if unicode.IsDigit([]rune(s)[0]) {
		return false, ErrInvalidString
	}
	return true, nil
}

func writeSeq(ch string, multStr string, builder *strings.Builder) {
	mult, _ := strconv.Atoi(multStr)
	builder.WriteString(strings.Repeat(ch, mult))
}

func isRuneSlash(str rune) bool {
	return string(str) == "\\"
}

func inRange(idx int, runes []rune) bool {
	return idx+1 < len(runes)
}

func Unpack(s string) (string, error) {
	isValid, err := validator(s)
	if !isValid {
		return "", err
	}

	runes := []rune(s)
	var result strings.Builder
	var prevRune rune
	var isSlashPrev bool

	for idx, curr := range runes {
		switch {
		case unicode.IsLetter(curr):
			if inRange(idx, runes) && unicode.IsDigit(runes[idx+1]) {
				prevRune = curr
				break
			}
			result.WriteRune(curr)
		case unicode.IsDigit(curr):
			if inRange(idx, runes) && unicode.IsDigit(runes[idx+1]) && !isSlashPrev {
				return "", ErrInvalidString
			}
			if isSlashPrev {
				result.WriteRune(curr)
				prevRune = curr
				break
			}

			writeSeq(string(prevRune), string(curr), &result)
		case isRuneSlash(curr):
			isSlashPrev = true
		}
	}

	return result.String(), nil
}
