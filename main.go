package main

import (
	"regexp"
	"strconv"
	"strings"
)

func toUpper(str string) string {
	return strings.ToUpper(str)
}

func toLower(str string) string {
	return strings.ToLower(str)
}

func toTitle(str string) string {
	return strings.Title(str)
}

func replaceHex(input string) string {
	// Regular expression to match a hexadecimal number
	hexRegex := regexp.MustCompile(`0x[0-9a-fA-F]+`)

	// Replace each hexadecimal number with its decimal equivalent
	output := hexRegex.ReplaceAllStringFunc(input, func(hexMatch string) string {
		hexValue, err := strconv.ParseInt(hexMatch[2:], 16, 64)
		if err != nil {
			// Ignore invalid hexadecimal numbers
			return hexMatch
		}
		decimalValue := strconv.FormatInt(hexValue, 10)
		return decimalValue
	})

	return output
}
