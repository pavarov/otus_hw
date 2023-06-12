package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder

	runes := []rune(s)
	runesLength := len(runes)

	if runesLength == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	for i := 0; i < runesLength; i++ {
		curr := runes[i]
		esc := string(curr) == "\\"
		nextIdx := i + 1
		if nextIdx >= runesLength {
			if unicode.IsLetter(curr) {
				result.WriteRune(curr)
				continue
			}
			continue
		}
		next := runes[nextIdx]

		if unicode.IsDigit(next) && unicode.IsDigit(curr) && !esc {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(next) {
			if nextIdx+1 <= runesLength-1 && unicode.IsDigit(runes[nextIdx+1]) && !esc {
				return "", ErrInvalidString
			}

			currStr := string(curr)
			if currStr == "\\" {
				nnIdx := nextIdx + 1

				if nnIdx < runesLength && unicode.IsDigit(runes[nnIdx]) {
					mi, _ := strconv.ParseInt(string(runes[nnIdx]), 10, 0)
					mult := int(mi)
					result.WriteString(strings.Repeat(string(next), mult))
					i += 2
					continue
				}

				result.WriteRune(next)
				i++
				continue
			}
			nextInt, _ := strconv.ParseInt(string(next), 10, 0)
			mult := int(nextInt)

			result.WriteString(strings.Repeat(currStr, mult))
			i++
			continue
		}

		if string(curr) == "\\" && string(next) == "\\" {
			nnIdx := nextIdx + 1

			if nnIdx < runesLength {
				if unicode.IsDigit(runes[nnIdx]) {
					mi, _ := strconv.ParseInt(string(runes[nnIdx]), 10, 0)
					mult := int(mi)
					result.WriteString(strings.Repeat(string(next), mult))
					i += 3
					continue
				} else if string(runes[nnIdx]) == "\\" {
					result.WriteString("\\")
					i++
					continue
				}
			}
			result.WriteString("\\")
		}

		if unicode.IsLetter(curr) {
			result.WriteRune(curr)
			continue
		}
	}
	return result.String(), nil
}
