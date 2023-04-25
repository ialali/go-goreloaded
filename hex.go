package goreloaded

import (
	"strconv"
	"strings"
)

func ReplaceHexWithDec(text string) string {
	words := strings.Fields(text)
	prevWord := ""
	for i, word := range words {
		if word == "(hex)" && i > 0 {
			// Replace previous word with its decimal equivalent
			prevHex := strings.TrimPrefix(prevWord, "0x")
			decNum, err := strconv.ParseInt(prevHex, 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(decNum, 10)
			}
		} else {
			// Store current word as previous word for next iteration
			prevWord = word
		}
	}
	// Remove any remaining instances of "(hex)"
	replacedText := strings.Join(words, " ")
	replacedText = strings.ReplaceAll(replacedText, "(hex)", "")
	return replacedText
}
