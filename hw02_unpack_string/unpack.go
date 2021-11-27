package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const escapeChar = '\\'

var (
	ErrInvalidString    = errors.New("invalid string")
	ErrUnexpectBehavior = errors.New("unexpected behavior")
)

func isSpecial(char rune) bool {
	return unicode.IsDigit(char) || char == escapeChar
}

func repeatRune(char rune, count rune) (string, error) {
	countInt, err := strconv.Atoi(string(count))

	return strings.Repeat(string(char), countInt), err
}

func Unpack(input string) (string, error) {
	var result strings.Builder

	casted := []rune(input)

	length := len(casted)
	for i := 0; i < length; {
		if unicode.IsDigit(casted[i]) {
			return "", ErrInvalidString
		}

		if casted[i] == escapeChar { // nolint:nestif
			if (i+1 >= length) || !isSpecial(casted[i+1]) {
				return "", ErrInvalidString
			}

			if i+2 < length && unicode.IsDigit(casted[i+2]) {
				str, err := repeatRune(casted[i+1], casted[i+2])
				if err != nil {
					panic(ErrUnexpectBehavior)
				}

				result.WriteString(str)
				i += 3
			} else {
				result.WriteRune(casted[i+1])
				i += 2
			}
		} else {
			if i+1 < length && unicode.IsDigit(casted[i+1]) {
				str, err := repeatRune(casted[i], casted[i+1])
				if err != nil {
					panic(ErrUnexpectBehavior)
				}

				result.WriteString(str)
				i += 2
			} else {
				result.WriteRune(casted[i])
				i++
			}
		}
	}

	return result.String(), nil
}
