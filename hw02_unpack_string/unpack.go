package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func unpackInitState(s string) ([]rune, int, strings.Builder) {
	runes := []rune(s)
	runesLength := len(runes)
	var result strings.Builder
	return runes, runesLength, result
}

func unpackValidator(runesLength int, runes []rune) (bool, error) {
	if runesLength == 0 {
		return true, nil
	}
	if unicode.IsDigit(runes[0]) {
		return true, ErrInvalidString
	}
	return false, nil
}

func initIterationState(runes []rune, idx int) (rune, int, int, bool) {
	currentRune := runes[idx]
	nextIdx := idx + 1
	idxAfterNext := nextIdx + 1
	currentRuneIsSlash := string(currentRune) == "\\"
	return currentRune, nextIdx, idxAfterNext, currentRuneIsSlash
}

func writeLetter(r rune, builder *strings.Builder) {
	if unicode.IsLetter(r) {
		builder.WriteRune(r)
	}
}

func writeSeq(ch string, multStr string, builder *strings.Builder) {
	mult, _ := strconv.Atoi(multStr)
	builder.WriteString(strings.Repeat(ch, mult))
}

func isSlash(str string) bool {
	return str == "\\"
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func getUnpackedString(builder *strings.Builder, err error) (string, error) {
	return builder.String(), err
}

func Unpack(s string) (string, error) {
	runes, runesLength, result := unpackInitState(s)
	invalidStr, err := unpackValidator(runesLength, runes)
	if invalidStr {
		return getUnpackedString(&result, err)
	}

	for idx := 0; idx < runesLength; idx++ {
		currentRune, nextIdx, idxAfterNext, currIsSlash := initIterationState(runes, idx)
		currStr := string(currentRune)
		if nextIdx == runesLength {
			writeLetter(currentRune, &result)
			return getUnpackedString(&result, nil)
		}
		nextRune := runes[nextIdx]

		switch {
		case isDigit(nextRune) && isDigit(currentRune) && !currIsSlash:
		case isDigit(nextRune) && idxAfterNext <= runesLength-1 && isDigit(runes[idxAfterNext]) && !currIsSlash:
			return "", ErrInvalidString
		}

		if isDigit(nextRune) {
			switch {
			case isSlash(currStr) && idxAfterNext < runesLength && isDigit(runes[idxAfterNext]): // \45
				writeSeq(string(nextRune), string(runes[idxAfterNext]), &result)
				idx += 2
			case isSlash(currStr): // \5
				result.WriteRune(nextRune)
				idx++
			default:
				writeSeq(currStr, string(nextRune), &result)
				idx++
			}
			continue
		}

		if isSlash(string(currentRune)) && isSlash(string(nextRune)) { // \\
			switch {
			case idxAfterNext < runesLength && isDigit(runes[idxAfterNext]): // \\3...
				writeSeq(string(nextRune), string(runes[idxAfterNext]), &result)
				idx += 3
			case idxAfterNext < runesLength && isSlash(string(runes[idxAfterNext])): // \\\...
				result.WriteString("\\")
				idx++
			default:
				result.WriteString("\\")
			}
			continue
		}
		writeLetter(currentRune, &result)
	}
	return getUnpackedString(&result, nil)
}
