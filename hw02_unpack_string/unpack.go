package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var escapeSymbol = '\\'

func repeatRune(char rune, count rune) (string, error) {
	countInt, err := strconv.Atoi(string(count))

	return strings.Repeat(string(char), countInt), err
}

func Unpack(input string) (string, error) {
	var result strings.Builder

	prev := rune(0)
	escaped := false

	for _, symbol := range input {
		if symbol == escapeSymbol {
			if escaped {
				escaped = false
				prev = escapeSymbol
				continue
			} else {
				escaped = true
			}
		}

		if unicode.IsDigit(symbol) {
			if escaped {
				prev = symbol
				escaped = false
				continue
			}

			if prev == rune(0) {
				return "", ErrInvalidString
			}

			str, err := repeatRune(prev, symbol)
			if err != nil {
				return "", ErrInvalidString
			}

			result.WriteString(str)
			prev = rune(0)

			continue
		}

		if prev != rune(0) {
			result.WriteRune(prev)
		}

		prev = symbol
	}

	if escaped {
		return "", ErrInvalidString
	}

	if prev != rune(0) {
		result.WriteRune(prev)
	}

	return result.String(), nil
}
